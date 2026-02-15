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
	// Tactical List: Web, Admin Panels, DBs, and Cloud Management
	// 80, 443 (Web), 2082-2087 (cPanel), 8443 (Plesk), 10000 (Webmin), 6443 (K8s)
	webAdminPorts := []int{80, 443, 2082, 2083, 2086, 2087, 8080, 8443, 8888, 9443, 10000, 6443}

	for _, port := range webAdminPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		
		// 1. Check if port is open
		conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
		if err != nil {
			continue
		}
		conn.Close()

		// 2. Identify the service details
		info := ""
		if port == 443 || port == 8443 || port == 2083 || port == 2087 || port == 9443 || port == 6443 {
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
	if err != nil {
		return "OPEN"
	}
	defer resp.Body.Close()
	return fmt.Sprintf("HTTP %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

func getSSLInfo(target string, port int) string {
	conf := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 1 * time.Second}, "tcp", fmt.Sprintf("%s:%d", target, port), conf)
	if err != nil {
		return "SSL_ERR"
	}
	defer conn.Close()
	
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		return fmt.Sprintf("SSL: %s | EXP: %s", certs[0].Subject.CommonName, certs[0].NotAfter.Format("2006-01-02"))
	}
	return "SSL_OPEN"
}
