package probes

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

// ExecuteContinuousMonitor runs the persistent latency HUD tracking loop
func ExecuteContinuousMonitor(target string, proto string, interval time.Duration) {
	fmt.Print(output.ClearLine)
	output.Banner()

	fmt.Printf("%s[+] LAUNCHING PERSISTENT HUD MONITOR METRICS FEED%s\n", output.NeonGreen, output.Reset)
	fmt.Printf(" • ROUTE TARGET:  %s%s%s\n", output.Cyan, target, output.Reset)
	fmt.Printf(" • WIRE PROTOCOL: %s%s%s\n", output.NeonPink, proto, output.Reset)
	fmt.Printf(" • DELAY INTERVAL: %s%v%s\n\n", output.Reset, interval, output.Reset)

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
		default: // TCP fallback
			address := target
			if !hasPort(target) {
				if proto == "https" {
					address += ":443"
				} else {
					address += ":80"
				}
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
			fmt.Printf("[%s] TRACE INTERCEPT ──> %sERR:%s %v\n", 
				time.Now().Format("15:04:05"), output.Red, output.Reset, err)
		} else {
			fmt.Printf("[%s] TRACE INTERCEPT ──> %sSTATUS:%s %-22s %s| LATENCY: %s%v%s\n", 
				time.Now().Format("15:04:05"), output.NeonGreen, output.Reset, status, 
				output.Gray, output.Cyan, elapsed, output.Reset)
		}

		time.Sleep(interval)
	}
}

func hasPort(host string) bool {
	_, _, err := net.SplitHostPort(host)
	return err == nil
}

// ExecutePortScan performs a full port scan (used by --full dox)
func ExecutePortScan(target string) []string {
	// Placeholder - expand with your full port scanning logic
	fmt.Printf("%s[*] Starting port scan on %s...%s\n", output.Cyan, target, output.Reset)
	
	// Example common ports (expand as needed)
	commonPorts := []int{21, 22, 23, 25, 53, 80, 443, 445, 1433, 3306, 3389, 5432, 5900, 8080}
	openPorts := []string{}

	for _, port := range commonPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 800*time.Millisecond)
		if err == nil {
			openPorts = append(openPorts, fmt.Sprintf("%d (OPEN)", port))
			conn.Close()
		}
	}

	return openPorts
}
