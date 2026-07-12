package stress

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// ExecuteTCPFlood tests how many stateful connections your firewall/server can maintain
func ExecuteTCPFlood(target string, concurrency int, durationSec int) {
	runTest(target, "tcp", concurrency, durationSec, func(addr string) {
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err == nil {
			conn.Close()
		}
	})
}

// ExecuteUDPFlood tests how your server handles connectionless packet saturation
func ExecuteUDPFlood(target string, concurrency int, durationSec int) {
	runTest(target, "udp", concurrency, durationSec, func(addr string) {
		conn, err := net.Dial("udp", addr)
		if err == nil {
			// UDP is fire and forget, send a small payload
			conn.Write([]byte("STRESS_TEST_PACKET"))
			conn.Close()
		}
	})
}

// ExecuteHTTPCapacityTest tests how your application layer handles high request volumes
func ExecuteHTTPCapacityTest(url string, method string, concurrency int, durationSec int) {
	client := &http.Client{Timeout: 3 * time.Second}
	runTest(url, "http", concurrency, durationSec, func(target string) {
		req, _ := http.NewRequest(method, target, nil)
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
		}
	})
}

// Internal helper to manage the "distress" lifecycle and concurrency
func runTest(target, mode string, concurrency, durationSec int, action func(string)) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)
	endTime := time.Now().Add(time.Duration(durationSec) * time.Second)

	fmt.Printf("[+] Resilience Test Engaged: Mode=%s, Target=%s, Concurrency=%d\n", mode, target, concurrency)

	for time.Now().Before(endTime) {
		semaphore <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()
			action(target)
		}()
	}
	wg.Wait()
	fmt.Printf("[+] Resilience Test Finished: %s\n", mode)
}
