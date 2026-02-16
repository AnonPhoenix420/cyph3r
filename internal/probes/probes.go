package probes

import (
	"net"
	"sync"
	"time"
)

// ScanPorts performs a tactical scan on common administrative ports
func ScanPorts(target string) []string {
	var openPorts []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	
	// Tactical port list: SSH, HTTP, HTTPS, HTTP-ALT, MySQL
	commonPorts := []string{"22", "80", "443", "8080", "3306"}

	for _, port := range commonPorts {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			address := net.JoinHostPort(target, p)
			
			// DialTimeout keeps the scan fast without the context package
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err == nil {
				conn.Close()
				mu.Lock()
				openPorts = append(openPorts, p)
				mu.Unlock()
			}
		}(port)
	}
	wg.Wait()
	return openPorts
}
