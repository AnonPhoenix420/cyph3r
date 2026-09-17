package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
	"github.com/AnonPhoenix420/cyph3r/internal/stress"
)

func main() {
	targetFlag := flag.String("target", "", "Target domain, IPv4 address, or full URL (e.g., 192.168.1.50, example.com, https://target.com:8443)")
	portFlag := flag.Int("p", 0, "Target port (Optional: auto-detected from URL or defaults to 80/443)")
	protoFlag := flag.String("proto", "tcp", "Wire protocol (tcp/udp)")
	monitorFlag := flag.Bool("monitor", false, "Engage live HUD connection monitor loop")
	
	// Stress & Benchmarking Flags
	hulkFlag := flag.Bool("hulk", false, "Engage continuous resilience HULK stress engine")
	slowlorisFlag := flag.Bool("slowloris", false, "Engage Slowloris slow-rate exhaustion engine")
	synFloodFlag := flag.Bool("synflood", false, "Engage Layer 4 TCP state/SYN flood engine")
	wrkFlag := flag.Bool("wrk", false, "Engage Wrk-style high-performance benchmarking engine")
	rudyFlag := flag.Bool("rudy", false, "Engage RUDY slow-POST exhaustion engine")
	h2Flag := flag.Bool("h2", false, "Engage HTTP/2 rapid reset stream engine")
	wsFlag := flag.Bool("ws", false, "Engage WebSocket connection & frame exhaustion engine")

	osintFlag := flag.Bool("osint", false, "Extract unmasked origin IPs, emails, phones, and social footprints")
	scanFlag := flag.Bool("scan", false, "Engage accelerated TCP port scanner and service probes")
	
	// Control & Timing Flags (with intelligent defaults)
	concurrencyFlag := flag.Int("c", 500, "Concurrency pool size (Default: 500 sockets/workers)")
	durationFlag := flag.Int("d", 0, "Test duration in seconds (Default: 0 for infinite/continuous until stopped)")
	intervalFlag := flag.Duration("interval", 2*time.Second, "HUD monitor ping interval")

	flag.Parse()

	if *targetFlag == "" {
		fmt.Println("[!] Error: --target parameter is required (e.g., --target 192.168.1.100 or --target https://example.com:8080).")
		return
	}

	// ==========================================
	// SMART TARGET & URL NORMALIZATION PARSER
	// ==========================================
	rawTarget := *targetFlag
	var finalURL string
	var targetHost string
	var targetPort int
	var isTLS bool

	// Ensure scheme exists for URL parsing fallback
	if !strings.HasPrefix(rawTarget, "http://") && !strings.HasPrefix(rawTarget, "https://") && !strings.HasPrefix(rawTarget, "ws://") && !strings.HasPrefix(rawTarget, "wss://") {
		if *portFlag == 443 || strings.Contains(rawTarget, ":443") {
			rawTarget = "https://" + rawTarget
		} else {
			rawTarget = "http://" + rawTarget
		}
	}

	parsedURL, err := url.Parse(rawTarget)
	if err != nil {
		targetHost = rawTarget
		isTLS = strings.HasPrefix(rawTarget, "https") || strings.HasPrefix(rawTarget, "wss")
	} else {
		targetHost = parsedURL.Hostname()
		isTLS = parsedURL.Scheme == "https" || parsedURL.Scheme == "wss"
		
		if parsedURL.Port() != "" {
			fmt.Sscanf(parsedURL.Port(), "%d", &targetPort)
		}
	}

	// Apply Port Fallback Logic
	if *portFlag > 0 {
		targetPort = *portFlag
	} else if targetPort == 0 {
		if isTLS {
			targetPort = 443
		} else {
			targetPort = 80
		}
	}

	// Construct standardized targets
	if strings.HasPrefix(rawTarget, "ws://") || strings.HasPrefix(rawTarget, "wss://") {
		finalURL = rawTarget
	} else {
		scheme := "http"
		if isTLS {
			scheme = "https"
		}
		if (scheme == "http" && targetPort == 80) || (scheme == "https" && targetPort == 443) {
			finalURL = fmt.Sprintf("%s://%s", scheme, targetHost)
		} else {
			finalURL = fmt.Sprintf("%s://%s:%d", scheme, targetHost, targetPort)
		}
	}

	targetAddr := fmt.Sprintf("%s:%d", targetHost, targetPort)

	// 1. Deep OSINT & Origin Unmasking Mode
	if *osintFlag {
		fmt.Printf("[+] LAUNCHING DEEP OSINT & PROXY-BYPASS INTEL SCAN: %s\n", targetHost)
		results := intel.DiscoverOriginAndOSINT(targetHost)

		fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║               CYPH3R DEEP OSINT FIELD INTELLIGENCE          ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		
		fmt.Println("\n[ UNMASKED REAL ORIGIN NODES ]")
		if len(results.RealIPs) > 0 {
			for _, ip := range results.RealIPs {
				fmt.Printf("  ↳ %s\n", ip)
			}
		} else {
			fmt.Println("  ↳ No unmasked origin IPs detected (Strict CDN edge encapsulation active).")
		}

		fmt.Println("\n[ HARVESTED EMAILS ]")
		if len(results.Emails) > 0 {
			for _, email := range results.Emails {
				fmt.Printf("  ↳ %s\n", email)
			}
		} else {
			fmt.Println("  ↳ None exposed in public records.")
		}

		fmt.Println("\n[ EXTRACTED PHONE VECTORS ]")
		if len(results.PhoneNumbers) > 0 {
			for _, phone := range results.PhoneNumbers {
				fmt.Printf("  ↳ %s\n", phone)
			}
		} else {
			fmt.Println("  ↳ None detected.")
		}

		fmt.Println("\n[ LEAKED SOCIAL MEDIA REFERENCES ]")
		if len(results.SocialHandles) > 0 {
			for _, soc := range results.SocialHandles {
				fmt.Printf("  ↳ https://%s\n", soc)
			}
		} else {
			fmt.Println("  ↳ None mapped.")
		}
		return
	}

	// 2. Accelerated Tactical Port Scan Mode
	if *scanFlag {
		fmt.Printf("[+] LAUNCHING ACCELERATED PORT SCANNER & SERVICE PROBES: %s\n", targetHost)
		openPorts := probes.ExecutePortScan(targetHost)
		
		if len(openPorts) > 0 {
			fmt.Printf("\n[+] Verified Open Listeners:\n")
			for _, portInfo := range openPorts {
				fmt.Printf("  ↳ %s\n", portInfo)
			}
		} else {
			fmt.Printf("\n[-] No open listening ports detected on standard profiles.\n")
		}
		return
	}

	// 3. Stress & Benchmarking Engines
	if *hulkFlag {
		stress.ExecuteContinuousStress(finalURL, *concurrencyFlag, *durationFlag)
		return
	}
	if *slowlorisFlag {
		stress.ExecuteSlowRateStress(finalURL, *concurrencyFlag, *durationFlag)
		return
	}
	if *synFloodFlag {
		stress.ExecuteTransportSynFlood(targetAddr, *concurrencyFlag, *durationFlag)
		return
	}
	if *wrkFlag {
		stress.ExecuteWrkBenchmark(finalURL, *concurrencyFlag, *durationFlag)
		return
	}
	if *rudyFlag {
		stress.ExecuteRudyStress(finalURL, *concurrencyFlag, *durationFlag)
		return
	}
	if *h2Flag {
		stress.ExecuteH2RapidResetStress(finalURL, *concurrencyFlag, *durationFlag)
		return
	}
	if *wsFlag {
		stress.ExecuteWebSocketStress(finalURL, *concurrencyFlag, *durationFlag)
		return
	}

	// 4. Live HUD Monitor Mode
	if *monitorFlag {
		fmt.Printf("[+] LAUNCHING PERSISTENT HUD MONITOR METRICS FEED\n • ROUTE TARGET: %s\n", targetAddr)
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), *intervalFlag)
		return
	}

	fmt.Printf("[+] Target resolved: %s (Port: %d). Use --osint, --scan, --monitor, --hulk, --slowloris, --synflood, --wrk, --rudy, --h2, or --ws to engage modules.\n", targetHost, targetPort)
}
