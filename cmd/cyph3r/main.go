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
	port := flag.Int("port", 80, "Port number")
	proto := flag.String("proto", "tcp", "tcp|udp|http|https|ack|ping")
	interval := flag.Duration("interval", 1*time.Second, "Delay between loops")
	phone := flag.String("phone", "", "Phone number for metadata")
	monitor := flag.Bool("monitor", false, "Enable continuous loop mode")
	flag.Parse()

	if *phone != "" {
		fmt.Println(output.BlueText("--- PHONE INTELLIGENCE ---"))
		fmt.Println(intel.PhoneLookup(*phone))
	}

	if *target == "" { return }

	// Phase 1: Intelligence
	data, _, err := intel.GetIntel(*target)
	if err == nil {
		fmt.Printf("\n%s\n", output.BlueText("--- TARGET INTELLIGENCE ---"))
		fmt.Printf("ISP/Handler:  %s\nOrganization: %s\n", data.ISP, data.Org)
		fmt.Printf("Location:     %s, %s (%s)\n", data.City, data.Country, data.Zip)
		fmt.Printf("Maps Link:    https://www.google.com/maps?q=%f,%f\n\n", data.Lat, data.Lon)
	}

	// Phase 2: Monitoring Loop
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	output.Info(fmt.Sprintf("Monitoring %s via %s (Port %d)...", *target, *proto, *port))

	for {
		select {
		case <-sigChan:
			fmt.Println("\n" + output.RedText("[!] Monitoring Stopped."))
			return
		default:
			ok, lat := probes.ExecuteProbe(*proto, *target, *port)
			ts := time.Now().Format("15:04:05")
			if ok {
				output.Success(fmt.Sprintf("[%s] UP | Latency: %v", ts, lat))
			} else {
				output.Down(fmt.Sprintf("[%s] DOWN | Request Timeout", ts))
			}

			if !*monitor { return }
			time.Sleep(*interval)
		}
	}
}
