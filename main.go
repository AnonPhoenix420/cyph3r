package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	output.Banner()

	// 1. Setup Flags (Keeping all your original options)
	target := flag.String("target", "localhost", "Target host/IP")
	port := flag.Int("port", 80, "Port")
	proto := flag.String("proto", "tcp", "tcp|udp|http|https|ack")
	rps := flag.Int("rps", 10, "Requests Per Second")
	duration := flag.Duration("duration", 30*time.Second, "Test duration")
	doIntel := flag.Bool("intel", false, "Trigger full GeoIP/WHOIS/Maps/ISP lookup")
	phone := flag.String("phone", "", "Phone number lookup")
	flag.Parse()

	// 2. Intelligence Phase (The "Work Proper" fixes)
	if *phone != "" {
		pData := intel.PhoneIntel(*phone)
		output.Info(fmt.Sprintf("Phone Metadata: %v", pData))
	}

	if *doIntel {
		output.Info("Gathering deep intelligence...")
		data, err := intel.GlobalLookup(*target)
		if err == nil {
			// This pulls all the specific data you requested: ISP, Zip, Maps, Lat/Lon
			fmt.Printf("\n--- [ TARGET INTELLIGENCE ] ---\n")
			fmt.Printf("IP Address:  %s\n", data.IP)
			fmt.Printf("Organization: %s\n", data.Org)
			fmt.Printf("ISP Handler:  %s\n", data.ISP)
			fmt.Printf("Location:     %s, %s, %s (Zip: %s)\n", data.City, data.Region, data.Country, data.Zip)
			fmt.Printf("Coordinates:  %f, %f\n", data.Lat, data.Lon)
			fmt.Printf("Google Maps:  %s\n", data.MapsURL)
			fmt.Printf("-------------------------------\n\n")
		}
		
		fmt.Println("--- [ WHOIS RECORD ] ---")
		fmt.Println(intel.Whois(*target))
	}

	// 3. The "Stress Test" Phase (Using your Worker Pool logic)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	output.Success(fmt.Sprintf("Starting %s probe against %s:%d at %d RPS", *proto, *target, *port, *rps))
	
	ticker := time.NewTicker(time.Second / time.Duration(*rps))
	defer ticker.Stop()
	
	timer := time.NewTimer(*duration)
	
	for {
		select {
		case <-ctx.Done():
			output.Warn("Test interrupted by user.")
			return
		case <-timer.C:
			output.Success("Duration reached. Test complete.")
			return
		case <-ticker.C:
			go func() {
				// This calls the fixed probe logic that handles all protocols safely
				success, latency := probes.RunProbe(*proto, *target, *port)
				if success {
					output.Info(fmt.Sprintf("Response from %s: lat=%v", *target, latency))
				} else {
					output.Down(fmt.Sprintf("Failed to reach %s", *target))
				}
			}()
		}
	}
}
