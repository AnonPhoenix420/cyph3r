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
	
	target := flag.String("target", "", "Target Host")
	port := flag.Int("port", 80, "Default Port")
	proto := flag.String("proto", "tcp", "Protocol (tcp/udp/http/https)")
	monitor := flag.Bool("monitor", false, "Continuous monitoring")
	scan := flag.Bool("scan", false, "Run Port Scanner")
	phone := flag.String("phone", "", "Phone metadata lookup")
	flag.Parse()

	if *target == "" && *phone == "" {
		output.Warn("Vector required. Use --target or --phone")
		return
	}

	if *phone != "" {
		output.Info("Decrypting Phone Metadata...")
		fmt.Println(output.BlueText(intel.PhoneLookup(*phone)))
	}

	if *target != "" {
		output.ScanAnimation()
		data, _ := intel.GetFullIntel(*target)
		output.PrintIntelHUD(*target, data.IPs, data.NS, data.BornOn, data.Registrar, data.ISP, fmt.Sprintf("%s, %s", data.City, data.Country), data.Coords)

		if *scan {
			output.Info("Initiating Port Reconnaissance...")
			results := probes.PortScanner(*target)
			output.PrintPortScan(results)
		}

		// Live Feed
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		fmt.Printf("\n%s\n", output.BlueText(fmt.Sprintf("──[ LIVE FEED: %s // %s ]──", *target, *proto)))

		for {
			select {
			case <-sigChan:
				output.Warn("HUD Signal Lost.")
				return
			default:
				success, latency := probes.ExecuteProbe(*proto, *target, *port)
				if success {
					output.Success(fmt.Sprintf("[%s] UP | Latency: %v", time.Now().Format("15:04:05"), latency))
				} else {
					output.Down("Node Unreachable")
				}
				if !*monitor { return }
				time.Sleep(1 * time.Second)
			}
		}
	}
}
