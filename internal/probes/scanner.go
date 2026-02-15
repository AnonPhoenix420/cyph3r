package probes

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

func ScanPorts(target string) []string {
	var results []string
	
	// Tactical "Shadow" Port List
	// Includes: Web, Admin, DBs, DevOps, and Obscure Backdoors
	ninjaPorts := []int{
		80, 443, 8080, 8443, 8888, 9090,        // Web & Proxies
		2082, 2083, 2086, 2087, 10000,          // Admin Panels
		21, 22, 23, 25, 53, 110, 143,           // Infrastructure
		3306, 5432, 6379, 27017, 1521,          // Databases
		2375, 6443, 9000, 3000, 5000,           // DevOps/Cloud
		8000, 8081, 9443, 3389, 5900,           // Shadow/Remote
	}

	for _, port := range ninjaPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		
		// 1. Tactical Delay (Ninja Timing)
		// Sleeping 150ms prevents triggering most basic "Port Scan" alerts
		time.Sleep(150 * time.Millisecond)

		conn, err := net.DialTimeout("tcp", address, 1200*time.Millisecond)
		if err != nil {
			continue // Port is likely closed or filtered
		}
		conn.Close()

		info := ""
		if isSSL(port) {
			info = getSSLInfo(target, port)
		} else {
			info = getHTTPInfo(target, port)
		}

		if info != "" {
			results = append(results, fmt.Sprintf("%d (%s)", port, info))
		} else {
			results = append(results, fmt.Sprintf("%d", port))
		}
	}
	return results
}

func isSSL(port int) bool {
	sslPorts := map[int]bool{443: true, 8443: true, 2083: true, 2087: true, 9443: true, 6443: true}
	return sslPorts[port]
}

func getHTTPInfo(target string, port int) string {
	client := http.Client{Timeout: 1 * time.Second}
	url := fmt.Sprintf("http://%s:%d", target, port)
	
	req, _ := http.NewRequest("GET", url, nil)
	// Ninja Move: Spoof a real browser to avoid being flagged as a bot
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	
	resp, err := client.Do(req)
	if err != nil { return "OPEN" }
	defer resp.Body.Close()
	return fmt.Sprintf("HTTP %d", resp.StatusCode)
}

func getSSLInfo(target string, port int) string {
	conf := &tls.Config{InsecureSkipVerify: true}
	dialer := &net.Dialer{Timeout: 1 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", fmt.Sprintf("%s:%d", target, port), conf)
	if err != nil { return "SSL_OPEN" }
	defer conn.Close()
	
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		return fmt.Sprintf("SSL: %s", certs[0].Subject.CommonName)
	}
	return "SSL_OPEN"
}
