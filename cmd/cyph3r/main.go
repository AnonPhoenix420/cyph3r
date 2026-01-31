package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	// Ensure these paths match your module name in go.mod
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// 1. HUD Initialization Phase
	output.Banner()        // From banners.go
	output.ScanAnimation() // The new HUD calibration effect

	// 2. Flags Configuration
	target := flag.String("target", "", "Target IP or Domain")
	port := flag.Int("port", 80, "Destination Port")
	proto := flag.String("proto", "tcp", "Protocol: tcp|udp|http|https|ack|ping")
	monitor := flag.Bool("monitor", false, "Enable continuous monitoring loop")
	interval := flag.Duration("interval", 1*time.Second, "Delay between probes (e.g., 2s)")
	phone := flag.String("phone", "", "Phone metadata lookup")
	flag.Parse()

	// 3. Early Exit for Empty Run
	if *target == "" && *phone == "" {
		output.Warn("No target or phone specified.")
		fmt.Println(" Usage: ./cyph3r --target <host> [--monitor]")
		fmt.Println(" Usage: ./cyph3r --phone <number>")
		return
	}

	// 4. Phone Intelligence Phase
	if *phone != "" {
		output.Info("Decrypting Phone Metadata...")
		fmt.Println(output.BlueText(intel.PhoneLookup(*phone)))
		if *target == "" { return } 
	}

	// 5. Target Intelligence Phase
	output.Info(fmt.Sprintf("Mapping Target Grid: %s", *target))
	data, _, err := intel.GetIntel(*target)
	if err == nil && data != nil {
		fmt.Printf("\n%s\n", output.BlueText("──[ NODE INTELLIGENCE ]──"))
		fmt.Printf(" ISP/ORG:     %s / %s\n", data.ISP, data.Org)
		fmt.Printf(" LOCATION:    %s, %s (%s)\n", data.City, data.Country, data.Zip)
		fmt.Printf(" COORDINATES: %.4f, %.4f\n", data.Lat, data.Lon)
		fmt.Printf(" MAPS:        https://www.google.com/maps?q=%.4f,%.4f\n", data.Lat, data.Lon)
	}

	// 6. Monitoring Loop Setup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	fmt.Printf("\n%s\n", output.BlueText(fmt.Sprintf("──[ LIVE FEED: %s:%d // %s ]──", *target, *port, *proto)))

	// 7. Execution Loop
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
				// Aligned output for cleaner HUD reading
				output.Success(fmt.Sprintf("[%s] UP   | Latency: %-10v", timestamp, latency))
			} else {
				output.Down(fmt.Sprintf("[%s] DOWN | Request Timed Out", timestamp))
			}

			// Single-run logic
			if !*monitor {
				fmt.Println(output.BlueText("────────────────────────────────────────────────────────"))
				return
			}

			time.Sleep(*interval)
		}
	}
}
