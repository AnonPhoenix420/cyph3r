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
	// Web, Admin Panels (cPanel, Plesk), and Server Management
	adminPorts := []int{80, 443, 2082, 2083, 2086, 2087, 8080, 8443, 10000, 6443}

	for _, port := range adminPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			continue
		}
		conn.Close()

		info := ""
		if port == 443 || port == 8443 || port == 2083 || port == 2087 || port == 6443 {
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

func getHTTPInfo(target string, port int) string {
	client := http.Client{Timeout: 1 * time.Second}
	url := fmt.Sprintf("http://%s:%d", target, port)
	resp, err := client.Get(url)
	if err != nil { return "OPEN" }
	defer resp.Body.Close()
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
