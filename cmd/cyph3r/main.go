package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// 1. HUD Boot Sequence
	output.Banner()
	output.ScanAnimation()

	// 2. Flags Configuration
	target := flag.String("target", "", "Target IP or Domain")
	port := flag.Int("port", 80, "Destination Port")
	proto := flag.String("proto", "tcp", "Protocol: tcp|udp|http|https|ack|ping")
	monitor := flag.Bool("monitor", false, "Continuous monitoring loop")
	interval := flag.Duration("interval", 1*time.Second, "Delay between probes")
	phone := flag.String("phone", "", "Phone metadata lookup")
	flag.Parse()

	// 3. Early Exit Check
	if *target == "" && *phone == "" {
		output.Warn("System Idle: No target/phone vector specified.")
		fmt.Println(" Usage: ./cyph3r --target <host> [--monitor]")
		return
	}

	// 4. Phone OSINT Phase
	if *phone != "" {
		output.Info("Decrypting Phone Metadata...")
		fmt.Println(output.BlueText(intel.PhoneLookup(*phone)))
		if *target == "" { return }
	}

	// 5. Advanced Intelligence Phase (The Satisfaction Fix)
	if *target != "" {
		output.Info(fmt.Sprintf("Mapping Infrastructure: %s", *target))
		
		// Run the high-density engine (New function)
		nodeData, err := intel.GetFullIntel(*target)
		
		// Fetch Geo-IP (Existing function)
		geoData, _ := intel.GetGeoData(*target)

		if err == nil {
			// Display the professional-grade HUD
			output.PrintIntelDisplay(
				*target,
				nodeData.IPs,
				nodeData.NS,
				nodeData.BornOn,
				nodeData.Registrar,
				geoData.ISP,
				geoData.Location,
				geoData.Coords,
			)
		}
	}

	// 6. Signal Handling for Graceful Exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// 7. Live Feed Initialization
	fmt.Printf("\n%s\n", output.BlueText(fmt.Sprintf("──[ LIVE FEED: %s:%d // %s ]──", *target, *port, *proto)))

	for {
		select {
		case <-sigChan:
			fmt.Println()
			output.Warn("HUD Signal Lost. Connection Terminated.")
			return
		default:
			start := time.Now()
			success, latency := probes.ExecuteProbe(*proto, *target, *port)
			timestamp := start.Format("15:04:05")

			if success {
				output.Success(fmt.Sprintf("[%s] UP   | Latency: %-10v", timestamp, latency))
			} else {
				output.Down(fmt.Sprintf("[%s] DOWN | Request Timed Out", timestamp))
			}

			if !*monitor {
				fmt.Println(output.BlueText("────────────────────────────────────────────────────────"))
				return
			}
			time.Sleep(*interval)
		}
	}
}
