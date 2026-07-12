package stress

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ExecuteHighIntensityStress saturates a target with connections
func ExecuteHighIntensityStress(target string, concurrency int, durationSec int) {
	var wg sync.WaitGroup
	// Semaphore to control concurrency
	semaphore := make(chan struct{}, concurrency)
	endTime := time.Now().Add(time.Duration(durationSec) * time.Second)

	fmt.Printf("[+] Stress test engaged: Target=%s, Concurrency=%d, Duration=%ds\n", target, concurrency, durationSec)

	for time.Now().Before(endTime) {
		semaphore <- struct{}{} // Acquire
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }() // Release

			// Attempt connection
			conn, err := net.DialTimeout("tcp", target, 2*time.Second)
			if err == nil {
				// Artificial pressure duration
				time.Sleep(50 * time.Millisecond)
				conn.Close()
			}
		}()
	}
	wg.Wait()
	fmt.Println("[+] Stress test complete.")
}
