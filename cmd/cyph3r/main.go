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
	output.Banner()

	target := flag.String("target", "", "Target IP/Domain")
	port := flag.Int("port", 80, "Port")
	proto := flag.String("proto", "tcp", "tcp|udp|http|https|ack|ping")
	interval := flag.Duration("interval", 2*time.Second, "Delay between loops")
	monitor := flag.Bool("monitor", false, "Enable continuous loop mode")
	flag.Parse()

	if *target == "" {
		output.Warn("Target is required.")
		return
	}

	// Run Intelligence once at start
	data, _, _ := intel.GetIntel(*target)
	if data != nil {
		output.Info(fmt.Sprintf("Target: %s (%s) | ISP: %s", *target, data.IP, data.ISP))
	}

	// Setting up Signal Handling for a clean exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	output.Success(fmt.Sprintf("Monitoring %s via %s...", *target, *proto))

	for {
		select {
		case <-sigChan:
			fmt.Println("\n")
			output.Warn("Monitoring stopped by user.")
			return
		default:
			start := time.Now()
			success, latency := probes.ExecuteProbe(*proto, *target, *port)

			if success {
				output.Success(fmt.Sprintf("[%s] %s is UP | Latency: %v", 
					time.Now().Format("15:04:05"), *target, latency))
			} else {
				output.Down(fmt.Sprintf("[%s] %s is DOWN | Timeout", 
					time.Now().Format("15:04:05"), *target))
			}

			if !*monitor {
				return // Exit after one run if monitor is false
			}

			time.Sleep(*interval) // Wait for the next loop
		}
	}
}
