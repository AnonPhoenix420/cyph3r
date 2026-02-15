package probes

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// High-risk signatures for the analyst to flag
var highRiskSignatures = []string{"OpenSSH_7.2", "Apache/2.2", "php/5.", "IIS/6.0", "Expired", "vulnerable"}

func ScanPorts(target string) []string {
	var results []string
	ninjaPorts := []int{
		80, 443, 8080, 8443, 8888, 2082, 2083, 2086, 2087, 
		21, 22, 23, 25, 3306, 5432, 6379, 6443,
	}

	for _, port := range ninjaPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		time.Sleep(150 * time.Millisecond) // Ninja timing

		conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
		if err != nil {
			continue
		}

		info := ""
		if isSSL(port) {
			info = getSSLInfo(target, port)
		} else {
			info = grabBanner(conn)
			if info == "" {
				info = getHTTPInfo(target, port)
			}
		}
		conn.Close()

		// Flag potential vulnerabilities
		riskTag := ""
		for _, sig := range highRiskSignatures {
			if strings.Contains(strings.ToLower(info), strings.ToLower(sig)) {
				riskTag = " [!] VULN_POTENTIAL"
			}
		}

		if info != "" {
			results = append(results, fmt.Sprintf("%d (%s)%s", port, info, riskTag))
		} else {
			results = append(results, fmt.Sprintf("%d", port))
		}
	}
	return results
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(800 * time.Millisecond))
	reader := bufio.NewReader(conn)
	banner, _ := reader.ReadString('\n')
	return strings.TrimSpace(banner)
}

func isSSL(port int) bool {
	sslPorts := map[int]bool{443: true, 8443: true, 2083: true, 2087: true, 9443: true}
	return sslPorts[port]
}

func getHTTPInfo(target string, port int) string {
	client := http.Client{Timeout: 1 * time.Second}
	url := fmt.Sprintf("http://%s:%d", target, port)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")
	
	resp, err := client.Do(req)
	if err != nil { return "" }
	defer resp.Body.Close()
	
	server := resp.Header.Get("Server")
	if server != "" {
		return fmt.Sprintf("HTTP %d | %s", resp.StatusCode, server)
	}
	return fmt.Sprintf("HTTP %d", resp.StatusCode)
}

func getSSLInfo(target string, port int) string {
	conf := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 1 * time.Second}, "tcp", fmt.Sprintf("%s:%d", target, port), conf)
	if err != nil { return "SSL_OPEN" }
	defer conn.Close()
	
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		return fmt.Sprintf("SSL: %s", certs[0].Subject.CommonName)
	}
	return "SSL_OPEN"
}
