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
	output.Banner()

	// Flags
	target := flag.String("target", "", "Target IP/Domain")
	port := flag.Int("port", 80, "Port")
	proto := flag.String("proto", "tcp", "tcp|udp|http|https|ack")
	rps := flag.Int("rps", 1, "Requests per second")
	phone := flag.String("phone", "", "Phone number info")
	flag.Parse()

	if *phone != "" {
		output.Info("Phone Intel: " + intel.PhoneLookup(*phone))
	}

	if *target == "" {
		output.Warn("No target specified. Exit.")
		os.Exit(1)
	}

	// 1. Deep Intel Phase
	data, whois, err := intel.GetDeepIntel(*target)
	if err == nil {
		fmt.Printf("\n%s\n", output.BoldText("--- CORE INTELLIGENCE ---"))
		fmt.Printf("ISP/Handler: %s\nOrganization: %s\nLocation: %s, %s (Zip: %s)\n", data.ISP, data.Org, data.City, data.Country, data.Zip)
		fmt.Printf("Coordinates: %f, %f\n", data.Lat, data.Lon)
		fmt.Printf("Google Maps: https://www.google.com/maps?q=%f,%f\n", data.Lat, data.Lon)
		fmt.Printf("Hostname: %s\n", data.Reverse)
		fmt.Println(output.BlueText("\n--- WHOIS DATA ---\n") + whois[:500] + "...") // Snippet
	}

	// 2. High-Speed Probe Phase
	output.Success(fmt.Sprintf("Starting %s stress at %d RPS...", *proto, *rps))
	ticker := time.NewTicker(time.Second / time.Duration(*rps))

	for range ticker.C {
		go func() {
			ok, lat := probes.ExecuteProbe(*proto, *target, *port)
			if ok {
				output.Success(fmt.Sprintf("[%s] Latency: %v", *proto, lat))
			} else {
				output.Down("Connection Failed")
			}
		}()
	}
}
