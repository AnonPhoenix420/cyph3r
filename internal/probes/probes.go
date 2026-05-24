package probes

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

// ExecuteContinuousMonitor runs the persistent latency HUD tracking loop
func ExecuteContinuousMonitor(target string, proto string, interval time.Duration) {
	cyan := "\033[38;5;51m"
	neonGreen := "\033[38;5;82m"
	neonPink := "\033[38;5;198m"
	reset := "\033[0m"
	red := "\033[31m"

	fmt.Printf("\n%s[+] LAUNCHING PERSISTENT HUD MONITOR METRICS FEED%s", neonGreen, reset)
	fmt.Printf("\n • ROUTE TARGET:  %s%s", cyan, target)
	fmt.Printf("\n • WIRE PROTOCOL: %s%s", neonPink, proto)
	fmt.Printf("\n • DELAY INTERVAL:%s %v\n\n", reset, interval)

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	for {
		startTime := time.Now()
		status := "OFFLINE"
		var err error

		switch proto {
		case "http", "https":
			urlStr := fmt.Sprintf("%s://%s", proto, target)
			resp, httpErr := client.Get(urlStr)
			if httpErr == nil {
				status = fmt.Sprintf("HTTP_%d_OK", resp.StatusCode)
				resp.Body.Close()
			} else {
				err = httpErr
			}
		default: // Fallback to raw TCP connection testing block
			address := target
			if !hasPort(target) {
				if proto == "https" { address += ":443" } else { address += ":80" }
			}
			conn, tcpErr := net.DialTimeout("tcp", address, 2*time.Second)
			if tcpErr == nil {
				status = "CONNECTED_ESTABLISHED"
				conn.Close()
			} else {
				err = tcpErr
			}
		}

		elapsed := time.Since(startTime)

		if err != nil {
			fmt.Printf("[%s] TRACE INTERCEPT ──> ERR: %s%v%s\n", time.Now().Format("15:04:05"), red, err, reset)
		} else {
			fmt.Printf("[%s] TRACE INTERCEPT ──> STATUS: %s%-22s%s | LATENCY: %s%v%s\n", 
				time.Now().Format("15:04:05"), neonGreen, status, reset, cyan, elapsed, reset)
		}

		time.Sleep(interval)
	}
}

func hasPort(host string) bool {
	_, _, err := net.SplitHostPort(host)
	return err == nil
}
