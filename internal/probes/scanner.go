package intel

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts checks common tactical ports
func ScanPorts(target string) []int {
	openPorts := []int{}
	commonPorts := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389, 8080}

	for _, port := range commonPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	return openPorts
}
