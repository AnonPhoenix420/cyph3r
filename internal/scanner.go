package probes
import (
	"fmt"
	"net"
	"sync"
	"time"
)
type ScanResult struct { Port int; Status string }
func PortScanner(t string) []ScanResult {
	var results []ScanResult
	var wg sync.WaitGroup
	ports := []int{21, 22, 80, 443, 3306, 8080}
	resChan := make(chan ScanResult, len(ports))
	for _, p := range ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, _ := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t, port), 1*time.Second)
			if conn != nil {
				conn.Close()
				resChan <- ScanResult{Port: port, Status: "OPEN"}
			}
		}(p)
	}
	wg.Wait()
	close(resChan)
	for r := range resChan { results = append(results, r) }
	return results
}
