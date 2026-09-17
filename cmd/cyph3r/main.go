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
	phoneFlag := flag.String("phone", "", "Target international phone number for metadata and carrier lookup")
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

	osintFlag := flag.Bool("osint", false, "Extract comprehensive intelligence report, unmasked IPs, and digital footprints")
	scanFlag := flag.Bool("scan", false, "Engage accelerated TCP port scanner and service probes")
	
	// Control & Timing Flags (with intelligent defaults)
	concurrencyFlag := flag.Int("c", 500, "Concurrency pool size (Default: 500 sockets/workers)")
	durationFlag := flag.Int("d", 0, "Test duration in seconds (Default: 0 for infinite/continuous until stopped)")
	intervalFlag := flag.Duration("interval", 2*time.Second, "HUD monitor ping interval")

	flag.Parse()

	// 1. Handle Dedicated Phone Intelligence Mode First
	if *phoneFlag != "" {
		fmt.Printf("[+] LAUNCHING PHONE METADATA DECRYPTION: %s\n", *phoneFlag)
		metrics := intel.GetPhoneMetrics(*phoneFlag)

		fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║               CYPH3R PHONE INTELLIGENCE REPORT                ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		fmt.Printf("  ↳ Number:          %s\n", *phoneFlag)
		fmt.Printf("  ↳ Line Status:     %s\n", metrics.LineStatus)
		fmt.Printf("  ↳ Carrier:         %s\n", metrics.Carrier)
		fmt.Printf("  ↳ Locale / Region: %s\n", metrics.Locale)
		fmt.Printf("  ↳ Country Code:    +%d\n", metrics.CountryCode)
		fmt.Printf("  ↳ National Format: %d\n", metrics.NationalNumber)
		fmt.Printf("  ↳ Is Mobile:       %t\n", metrics.IsMobile)
		fmt.Printf("  ↳ Risk Score:      %d/100\n", metrics.Risk)
		return
	}

	if *targetFlag == "" {
		fmt.Println("[!] Error: --target or --phone parameter is required (e.g., --target 192.168.1.100 or --phone +14155552671).")
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

	// 2. Deep OSINT & Comprehensive Intelligence Report Mode
	if *osintFlag {
		fmt.Printf("[+] LAUNCHING FULL-STACK COMPREHENSIVE INTEL SCAN: %s\n", targetHost)
		report := intel.ExecuteComprehensiveReport(targetHost)

		fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║         CYPH3R COMPREHENSIVE INTELLIGENCE FIELD REPORT        ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		
		fmt.Printf("\n[ TARGET METADATA ]\n")
		fmt.Printf("  ↳ Target:     %s\n", report.Target)
		fmt.Printf("  ↳ Type:       %s\n", report.TargetType)
		fmt.Printf("  ↳ Primary IP: %s\n", report.ReverseDNS)
		fmt.Printf("  ↳ Risk Score: %d/100\n", report.RiskScore)
		fmt.Printf("  ↳ Timestamp:  %s\n", report.Timestamp.Format(time.RFC3339))

		fmt.Printf("\n[ GEOLOCATION TELEMETRY ]\n")
		fmt.Printf("  ↳ Location:    %s, %s (%s)\n", report.Location.City, report.Location.Country, report.Location.CountryCode)
		fmt.Printf("  ↳ Coordinates: %s\n", report.Location.Coordinates)

		fmt.Printf("\n[ UNMASKED NODES & ASSOCIATED ASSETS ]\n")
		if len(report.Associated) > 0 {
			for _, asset := range report.Associated {
				fmt.Printf("  ↳ %s\n", asset)
			}
		} else {
			fmt.Println("  ↳ No auxiliary nodes mapped.")
		}

		fmt.Printf("\n[ HARVESTED EMAILS ]\n")
		if len(report.Emails) > 0 {
			for _, email := range report.Emails {
				fmt.Printf("  ↳ %s\n", email)
			}
		} else {
			fmt.Println("  ↳ None exposed.")
		}

		fmt.Printf("\n[ EXTRACTED PHONE VECTORS ]\n")
		if len(report.Phones) > 0 {
			for _, phone := range report.Phones {
				fmt.Printf("  ↳ %s\n", phone)
			}
		} else {
			fmt.Println("  ↳ None detected.")
		}

		fmt.Printf("\n[ DATABASE EXPOSURE METRICS ]\n")
		fmt.Printf("  ↳ SQL Exposed: %t (Risk: %s)\n", report.SQLCheck.Exposed, report.SQLCheck.RiskLevel)

		fmt.Printf("\n[ MAPPED DIGITAL FOOTPRINTS ]\n")
		if len(report.SocialProfiles) > 0 {
			for _, profile := range report.SocialProfiles {
				fmt.Printf("  ↳ [%s] %s (Confidence: %d%%)\n", profile.Platform, profile.ProfileURL, profile.Confidence)
			}
		} else {
			fmt.Println("  ↳ None mapped.")
		}
		return
	}

	// 3. Accelerated Tactical Port Scan Mode
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

	// 4. Stress & Benchmarking Engines
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

	// 5. Live HUD Monitor Mode
	if *monitorFlag {
		fmt.Printf("[+] LAUNCHING PERSISTENT HUD MONITOR METRICS FEED\n • ROUTE TARGET: %s\n", targetAddr)
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), *intervalFlag)
		return
	}

	fmt.Printf("[+] Target resolved: %s (Port: %d). Use --osint, --scan, --monitor, --phone, or stress engines to engage modules.\n", targetHost, targetPort)
}
