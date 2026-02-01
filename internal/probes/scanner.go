package probes

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ScanResult struct {
	Port   int
	Status string
}

// PortScanner executes a high-speed concurrent check
func PortScanner(target string) []ScanResult {
	var results []ScanResult
	var wg sync.WaitGroup
	
	// Top 15 essential ports for rapid recon
	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 445, 3306, 3389, 5432, 8080, 8443}
	resultChan := make(chan ScanResult, len(commonPorts))

	for _, port := range commonPorts {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, p)
			// Strict timeout for "Zero-Key" speed
			conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
			
			if err == nil {
				conn.Close()
				resultChan <- ScanResult{Port: p, Status: "OPEN"}
			}
		}(port)
	}

	// Wait for workers in the background
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		results = append(results, res)
	}
	return results
}
