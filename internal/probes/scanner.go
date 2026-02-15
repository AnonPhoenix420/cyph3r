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

func ScanPorts(target string) []string {
	var results []string
	ninjaPorts := []int{
		80, 443, 8080, 8443, 8888, 9090,
		2082, 2083, 2086, 2087, 10000,
		21, 22, 23, 25, 53, 110, 143,
		3306, 5432, 6379, 27017, 3389,
	}

	for _, port := range ninjaPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		time.Sleep(150 * time.Millisecond) // Tactical jitter

		conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
		if err != nil {
			continue
		}

		info := ""
		if isSSL(port) {
			info = getSSLInfo(target, port)
		} else {
			// First, try a raw banner grab (works for SSH, FTP, SMTP)
			info = grabBanner(conn)
			// If it's silent, try an HTTP probe
			if info == "" {
				info = getHTTPInfo(target, port)
			}
		}
		conn.Close()

		if info != "" {
			results = append(results, fmt.Sprintf("%d (%s)", port, info))
		} else {
			results = append(results, fmt.Sprintf("%d", port))
		}
	}
	return results
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(800 * time.Millisecond))
	reader := bufio.NewReader(conn)
	banner, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
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
	if err != nil { return "OPEN" }
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
