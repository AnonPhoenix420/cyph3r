package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Setup Flags
	target := flag.String("t", "", "Target domain (e.g., president.ir)")
	phone := flag.String("p", "", "Phone number (e.g., +98912...)")
	flag.Parse()

	// 2. Banner Logic
	fmt.Print(output.NeonPink)
	fmt.Println("______      ____  __  __ _____ ____ ")
	fmt.Println("  / ____/_  __/ __ \\/ / / /|__  // __ \\")
	fmt.Println(" / /   / / / / /_/ / /_/ /  /_ </ /_/ /")
	fmt.Println("/ /___/ /_/ / ____/ __  / ___/ / _, _/")
	fmt.Println("\\____/\\__, /_/   /_/ /_/ /____/_/ |_| ")
	fmt.Println("     /____/         NETWORK_INTEL_SYSTEM")
	fmt.Print(output.Reset)

	// 3. Routing Logic
	if *target != "" {
		executeTargetIntel(*target)
	} else if *phone != "" {
		executePhoneIntel(*phone)
	} else {
		output.Error("No target specified. Use -t <domain> or -p <phone>")
		os.Exit(1)
	}
}

func executeTargetIntel(target string) {
	// Pulse the node first
	output.PulseNode(target)

	// Start Fancy Animation
	done := make(chan bool)
	go output.LoadingAnimation(done, "Remote Node Intelligence")

	// Perform High-End Intel Gathering
	data, err := intel.GetTargetIntel(target)
	
	// Kill Animation
	done <- true

	if err != nil {
		output.Error(fmt.Sprintf("Intel Extraction Failed: %v", err))
		return
	}

	// Render the Chromatic HUD
	output.DisplayHUD(data)
}

func executePhoneIntel(number string) {
	output.Info("Initializing Satellite Uplink...")
	
	done := make(chan bool)
	go output.LoadingAnimation(done, "Digital Footprint Analysis")

	data, err := intel.GetPhoneIntel(number)
	
	done <- true

	if err != nil {
		output.Error(fmt.Sprintf("Phone OSINT Failed: %v", err))
		return
	}

	output.DisplayPhoneHUD(data)
}
