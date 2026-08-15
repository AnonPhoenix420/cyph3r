package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"cyph3r/internal/intel"
	"cyph3r/internal/probes"
	"cyph3r/internal/stress"
)

func main() {
	targetFlag := flag.String("target", "", "Target domain or IP address")
	portFlag := flag.Int("p", 443, "Target port")
	protoFlag := flag.String("proto", "tcp", "Wire protocol (tcp/udp)")
	monitorFlag := flag.Bool("monitor", false, "Engage live HUD connection monitor loop")
	hulkFlag := flag.Bool("hulk", false, "Engage continuous resilience stress engine")
	osintFlag := flag.Bool("osint", false, "Extract unmasked origin IPs, emails, phones, and social footprints")
	concurrencyFlag := flag.Int("c", 1000, "Concurrency pool size for stress tests")
	durationFlag := flag.Int("d", 0, "Test duration in seconds (0 for infinite loop)")
	intervalFlag := flag.Duration("interval", 2*time.Second, "HUD monitor ping interval")

	flag.Parse()

	if *targetFlag == "" {
		fmt.Println("[!] Error: --target parameter is required.")
		return
	}

	cleanHost := strings.TrimPrefix(strings.TrimPrefix(*targetFlag, "https://"), "http://")
	if idx := strings.Index(cleanHost, "/"); idx != -1 {
		cleanHost = cleanHost[:idx]
	}

	targetAddr := fmt.Sprintf("%s:%d", cleanHost, *portFlag)

	// 1. Deep OSINT & Origin Unmasking Mode
	if *osintFlag {
		fmt.Printf("[+] LAUNCHING DEEP OSINT & PROXY-BYPASS INTEL SCAN: %s\n", cleanHost)
		results := intel.DiscoverOriginAndOSINT(cleanHost)

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

	// 2. Continuous Stress / Hulk Mode
	if *hulkFlag {
		stress.ExecuteContinuousStress(targetAddr, *concurrencyFlag, *durationFlag)
		return
	}

	// 3. Live HUD Monitor Mode
	if *monitorFlag {
		fmt.Printf("[+] LAUNCHING PERSISTENT HUD MONITOR METRICS FEED\n • ROUTE TARGET: %s\n", targetAddr)
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), *intervalFlag)
		return
	}

	fmt.Printf("[+] Target specified: %s. Use --osint, --monitor, or --hulk to engage modules.\n", targetAddr)
}
