package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// Initialize HUD
	output.Banner()

	// Define Flags
	target := flag.String("target", "", "Target Domain or IP")
	scan := flag.Bool("scan", false, "Perform accelerated port scan")
	phone := flag.String("phone", "", "Lookup international phone metadata")
	monitor := flag.Bool("monitor", false, "Enable live HUD latency feed")
	proto := flag.String("proto", "tcp", "Protocol for monitor (tcp/http/https)")
	port := flag.Int("port", 80, "Port for monitor/scan")
	interval := flag.Duration("interval", 2*time.Second, "Monitoring refresh interval")

	flag.Parse()

	// Logic 1: Phone Metadata
	if *phone != "" {
		output.Info("Decrypting Phone Vector...")
		fmt.Printf(" %s\n\n", output.MagText(intel.PhoneLookup(*phone)))
		return
	}

	// Logic 2: Target Recon & Intelligence
	if *target != "" {
		output.ScanAnimation()

		// OSINT Gathering
		data, _ := intel.GetFullIntel(*target)
		output.PrintIntelHUD(*target, data.IPs, data.ISP, fmt.Sprintf("%s, %s", data.City, data.Country))

		// Optional Accelerated Port Scan
		if *scan {
			output.Info("Initiating Accelerated Reconnaissance...")
			results := probes.PortScanner(*target)
			output.PrintPortScan(results)
		}

		// Optional Live HUD Monitor
		if *monitor {
			output.Info("Starting Live HUD Feed (Ctrl+C to exit)...")
			for {
				up, lat := probes.ExecuteProbe(*proto, *target, *port)
				status := output.RedText("DOWN")
				if up {
					status = output.GreenText("ACTIVE")
				}
				fmt.Printf("\r [%s] Protocol: %s | Latency: %s   ", status, output.YellowText(*proto), output.CyanText(lat))
				time.Sleep(*interval)
			}
		}
	} else {
		// No args provided
		fmt.Println(output.YellowText(" [!] No target specified. Use --help for usage guides."))
	}
}
