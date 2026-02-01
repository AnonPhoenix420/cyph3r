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
	// 1. HUD Initialization
	output.Banner()
	output.ScanAnimation()

	// 2. Flags Configuration
	target := flag.String("target", "", "Target IP or Domain")
	port := flag.Int("port", 80, "Destination Port")
	proto := flag.String("proto", "tcp", "Protocol: tcp|udp|http|https|ack|ping")
	monitor := flag.Bool("monitor", false, "Enable continuous monitoring loop")
	interval := flag.Duration("interval", 1*time.Second, "Delay between probes")
	phone := flag.String("phone", "", "Phone metadata lookup")
	flag.Parse()

	// 3. Early Exit Check
	if *target == "" && *phone == "" {
		output.Warn("System Idle: No target vector specified.")
		fmt.Println(" Usage: ./cyph3r --target <host> [--monitor]")
		return
	}

	// 4. Phone Intelligence Phase
	if *phone != "" {
		output.Info("Decrypting Phone Metadata...")
		fmt.Println(output.BlueText(intel.PhoneLookup(*phone)))
		if *target == "" {
			return
		}
	}

	// 5. Target Intelligence Phase (The Wireframe HUD)
	if *target != "" {
		output.Info(fmt.Sprintf("Mapping Infrastructure: %s", *target))
		
		data, err := intel.GetFullIntel(*target)
		if err != nil {
			output.Warn(fmt.Sprintf("Intelligence acquisition failed: %v", err))
		} else {
			// Matches the 8 arguments in your modular status.go
			output.PrintIntelHUD(
				*target,
				data.IPs,
				data.NS,
				data.BornOn,
				data.Registrar,
				data.ISP,
				fmt.Sprintf("%s, %s", data.City, data.Country),
				data.Coords,
			)
		}
	}

	// 6. Monitoring Loop Setup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

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
