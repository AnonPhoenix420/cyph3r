package probes

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(target string) []string {
	var results []string
	ports := []int{80, 443, 8080, 8443, 2082, 2083}
	for _, p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			conn.Close()
			results = append(results, fmt.Sprintf("PORT %d: OPEN [ACK/SYN]", p))
		}
	}
	return results
}
