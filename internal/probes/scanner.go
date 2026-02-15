package probes

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// ScanPorts probes the target and grabs service banners
func ScanPorts(target string) []string {
	var results []string
	commonPorts := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389, 8080}

	for _, port := range commonPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			banner := grabBanner(conn)
			if banner != "" {
				results = append(results, fmt.Sprintf("%d (%s)", port, banner))
			} else {
				results = append(results, fmt.Sprintf("%d", port))
			}
			conn.Close()
		}
	}
	return results
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	// Try to trigger a response from web servers, otherwise just wait for service announcement (SSH/FTP)
	fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n") 
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
