package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	output.Banner()

	target := flag.String("target", "", "Target Domain/IP")
	port := flag.Int("port", 80, "Port")
	proto := flag.String("proto", "tcp", "tcp|udp|http|https|ack")
	phone := flag.String("phone", "", "Phone number info")
	flag.Parse()

	// 1. Phone Intelligence
	if *phone != "" {
		output.Info("Phone Metadata: " + intel.PhoneLookup(*phone))
	}

	if *target == "" {
		if *phone == "" { output.Warn("Use -target or -phone. See --help"); os.Exit(1) }
		return
	}

	// 2. Deep Network Intel (The "Work Proper" Request)
	data, err := intel.GetIntel(*target)
	if err == nil && data.Status == "success" {
		fmt.Printf("\n%s\n", output.BoldText("===== TARGET INTELLIGENCE ====="))
		fmt.Printf("IP Address:   %s\n", data.IP)
		fmt.Printf("Hostname:     %s\n", data.ReverseDNS)
		fmt.Printf("ISP Handler:  %s\n", data.ISP)
		fmt.Printf("Organization: %s\n", data.Org)
		fmt.Printf("Location:     %s, %s (Zip: %s)\n", data.City, data.Country, data.Zip)
		fmt.Printf("GPS Coords:   %f, %f\n", data.Lat, data.Lon)
		fmt.Printf("Google Maps:  https://www.google.com/maps?q=%f,%f\n", data.Lat, data.Lon)
		fmt.Printf("WHOIS:        %s\n", intel.Whois(*target))
		fmt.Println("===============================\n")
	}

	// 3. The Multi-Protocol Probe
	output.Info(fmt.Sprintf("Probing %s via %s on port %d...", *target, *proto, *port))
	
	start := time.Now()
	success := false
	address := fmt.Sprintf("%s:%d", *target, *port)

	switch *proto {
	case "tcp", "ack":
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err == nil {
			conn.Close()
			success = true
		}
	case "udp":
		conn, err := net.DialTimeout("udp", address, 3*time.Second)
		if err == nil {
			conn.Close()
			success = true
		}
	case "http", "https":
		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(fmt.Sprintf("%s://%s", *proto, *target))
		if err == nil {
			resp.Body.Close()
			success = (resp.StatusCode < 400)
		}
	}

	if success {
		output.Success(fmt.Sprintf("Target UP! Latency: %v", time.Since(start)))
	} else {
		output.Down("Target Unreachable or Port Closed.")
	}
}
