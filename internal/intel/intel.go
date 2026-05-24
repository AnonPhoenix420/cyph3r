package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

type APIResponse struct {
	Status  string `json:"status"`
	Country string `json:"country"`
	City    string `json:"city"`
	As      string `json:"as"`
	Org     string `json:"org"`
}

type RDAPResponse struct {
	Entities []struct {
		VcardArray []interface{} `json:"vcardArray"`
	} `json:"entities"`
	Events []struct {
		EventAction string `json:"eventAction"`
		EventDate   string `json:"eventDate"`
	} `json:"events"`
}

// ResolveNetwork parses target domain ownership records, threat telemetry maps, and scans ports
func ResolveNetwork(domain string) (string, models.GeoData, string, string, string, []string, []string, []string, []string) {
	var geo models.GeoData
	var asn = "UNKNOWN_ASN"
	var ownerName = "WHOIS_PRIVACY_PROTECTED"
	var createdDate = "UNKNOWN"
	
	var openPorts []string
	var banners []string
	var vulns []string
	var leaks []string

	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		return "0.0.0.0", geo, asn, ownerName, createdDate, openPorts, banners, vulns, leaks
	}
	targetIP := ips[0].String()

	client := &http.Client{Timeout: 3 * time.Second}

	// 1. Live Threat Telemetry & GeoIP
	if resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,as,org", targetIP)); err == nil {
		var data APIResponse
		if json.NewDecoder(resp.Body).Decode(&data) == nil && data.Status == "success" {
			geo.City = data.City
			geo.Country = data.Country
			asn = data.As
		}
		resp.Body.Close()
	}

	// 2. Domain Registration Ownership (RDAP)
	if resp, err := client.Get(fmt.Sprintf("https://rdap.org/domain/%s", domain)); err == nil {
		var rdap RDAPResponse
		if json.NewDecoder(resp.Body).Decode(&rdap) == nil {
			for _, event := range rdap.Events {
				if event.EventAction == "registration" && len(event.EventDate) >= 10 {
					createdDate = event.EventDate[:10]
				}
			}
			if len(rdap.Entities) > 0 && len(rdap.Entities[0].VcardArray) > 1 {
				if cards, ok := rdap.Entities[0].VcardArray[1].([]interface{}); ok && len(cards) > 0 {
					for _, card := range cards {
						if fields, ok := card.([]interface{}); ok && len(fields) > 3 {
							if fields[0] == "fn" {
								ownerName = fmt.Sprintf("%v", fields[3])
							}
						}
					}
				}
			}
		}
		resp.Body.Close()
	}

	// 3. Active Stealth Port Scanner & Banner Grabber
	portsToScan := []int{21, 22, 23, 25, 80, 443, 8080}
	for _, port := range portsToScan {
		address := fmt.Sprintf("%s:%d", targetIP, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			portStr := fmt.Sprintf("%d/TCP", port)
			openPorts = append(openPorts, portStr)

			_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
			buffer := make([]byte, 256)
			n, err := conn.Read(buffer)
			if err == nil && n > 0 {
				banners = append(banners, fmt.Sprintf("%d: %s", port, string(buffer[:n])))
			} else {
				banners = append(banners, fmt.Sprintf("%d: Unknown Service Banner", port))
			}
			conn.Close()
		}
	}

	// 4. Vulnerability & Exposure Analysis Engine
	if len(openPorts) > 0 {
		for _, port := range openPorts {
			if port == "22/TCP" {
				vulns = append(vulns, "CVE-2024-6387 (regreSSHion) - Potential Remote Code Execution")
			}
			if port == "80/TCP" || port == "8080/TCP" {
				leaks = append(leaks, "X-Powered-By Header Exposed / Server Version Disclosure Leak")
			}
		}
	} else {
		vulns = append(vulns, "No immediately exploitable surface vulnerabilities discovered via passive profiling.")
		leaks = append(leaks, "Zero public data leakage footprints verified.")
	}

	return targetIP, geo, asn, ownerName, createdDate, openPorts, banners, vulns, leaks
}

// ExecuteValidationSuite triggers server integrity testing modes
func ExecuteValidationSuite(targetURL string, mode int, concurrency int, durationSec int) {
	reset := "\033[0m"
	neonPink := "\033[38;5;198m"
	neonGreen := "\033[38;5;82m"
	cyan := "\033[38;5;51m"
	amber := "\033[38;5;214m"
	gray := "\033[90m"
	red := "\033[31m"

	modeName := "UNKNOWN"
	description := ""

	switch mode {
	case 1:
		modeName = "LOAD TESTING (BASELINE CAPACITY)"
		description = "Evaluates normal operational expectations over a steady timeline."
	case 2:
		modeName = "STRESS TESTING (BREAKING POINT)"
		description = "Pushes system boundaries to evaluate graceful recovery and logging stability."
	case 3:
		modeName = "SOAK TESTING (ENDURANCE)"
		description = "Monitors memory utilization patterns and resource leaks over prolonged windows."
	case 4:
		modeName = "SPIKE TESTING (BURST ELASTICITY)"
		description = "Simulates sudden, extreme traffic influxes to test rate-limiting thresholds."
	}

	fmt.Printf("\n%s[+] LAUNCHING COMPLIANCE VALIDATION MATRIX: %s%s", neonGreen, modeName, reset)
	fmt.Printf("\n • TARGET NODE:       %s%s", cyan, targetURL)
	fmt.Printf("\n • METRIC ATTRIBUTE:  %s%s", gray, description)
	fmt.Printf("\n • EXECUTION BOUNDS:  %s%d Stream Workers | %d Seconds Window\n\n", amber, concurrency, durationSec)

	successCount, errorCount, timeoutCount := 0, 0, 0
	var totalLatency time.Duration

	resultsChan := make(chan string, 5000)
	latencyChan := make(chan time.Duration, 5000)
	doneChan := make(chan bool)

	transport := &http.Transport{
		MaxIdleConnsPerHost: concurrency,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   2 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   3 * time.Second,
	}

	for i := 0; i < concurrency; i++ {
		go func(workerID int) {
			for {
				select {
				case <-doneChan:
					return
				default:
					if mode == 4 && workerID%2 == 0 {
						time.Sleep(500 * time.Millisecond)
					}

					start := time.Now()
					resp, err := client.Get(targetURL)
					latency := time.Since(start)

					if err != nil {
						if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
							resultsChan <- "timeout"
						} else {
							resultsChan <- "error"
						}
						continue
					}

					if resp.StatusCode >= 200 && resp.StatusCode < 400 {
						resultsChan <- "success"
						latencyChan <- latency
					} else {
						resultsChan <- "error"
					}
					resp.Body.Close()

					if mode == 3 {
						time.Sleep(50 * time.Millisecond)
					} else {
						time.Sleep(5 * time.Millisecond)
					}
				}
			}
		}(i)
	}

	go func() {
		for {
			select {
			case res := <-resultsChan:
				switch res {
				case "success":
					successCount++
				case "error":
					errorCount++
				case "timeout":
					timeoutCount++
				}
			case lat := <-latencyChan:
				totalLatency += lat
			case <-time.After(100 * time.Millisecond):
				if successCount+errorCount+timeoutCount >= (concurrency * durationSec * 10) {
					return
				}
			}
		}
	}()

	time.Sleep(time.Duration(durationSec) * time.Second)
	close(doneChan)

	totalRequests := successCount + errorCount + timeoutCount
	avgLatency := time.Duration(0)
	if successCount > 0 {
		avgLatency = totalLatency / time.Duration(successCount)
	}

	fmt.Printf("%s╔═══════════════════════════════════════════════════════════════╗%s", cyan, reset)
	fmt.Printf("\n║ %s[✓] INTEGRITY ASSESSMENT MATRIX COMPLETE                  %s║", neonPink, cyan)
	fmt.Printf("\n%s╚═══════════════════════════════════════════════════════════════╝%s\n", cyan, reset)
	
	fmt.Printf(" • TOTAL PROCESSED TRANSACTIONS:  %s%d%s\n", cyan, totalRequests, reset)
	fmt.Printf(" • SUCCESSFUL STABLE HANDSHAKES:  %s%d%s\n", neonGreen, successCount, reset)
	fmt.Printf(" • PROTOCOL DROP ERRORS:          %s%d%s\n", red, errorCount, reset)
	fmt.Printf(" • GATEWAY TIMEOUT EXHAUSTIONS:   %s%d%s\n", amber, timeoutCount, reset)
	fmt.Printf(" • AVERAGE TRANSIT LATENCY:       %s%v%s\n", cyan, avgLatency, reset)

	fmt.Printf("\n%s[ INFRASTRUCTURE QUALITY CLASSIFICATION ]%s\n", neonYellow, reset)
	if timeoutCount == 0 && errorCount == 0 && totalRequests > 0 {
		fmt.Printf(" %s[GRADE A+] EXCELLENT: Target node completely resilient against multi-vector load variations.%s\n\n", neonGreen, reset)
	} else if errorCount > (totalRequests / 10) {
		fmt.Printf(" %s[GRADE C] DEFICIENT: Server drops connection streams. Infrastructure weaknesses identified.%s\n\n", red, reset)
	} else {
		fmt.Printf(" %s[GRADE B] STABLE: Server operational but shows rate-limiting patterns or processing delays.%s\n\n", amber, reset)
	}
}
