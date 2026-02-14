package main

import (
	"flag"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// Define tactical flags
	target := flag.String("target", "", "Target domain or IP")
	phone := flag.String("phone", "", "Target phone number")
	scan := flag.Bool("scan", false, "Enable port scanning")
	flag.Parse()

	// 1. Initialize System Banner
	output.Banner()

	// 2. Branch Logic: Phone Intel
	if *phone != "" {
		intel.GetPhoneIntel(*phone)
		return
	}

	// 3. Branch Logic: Domain/IP Intel
	if *target == "" {
		output.Error("Input required. Use -target <domain> or -phone <number>")
		return
	}

	// Visual acquisition signal
	output.PulseNode(*target)

	// Fetch Full Intelligence Data
	data, err := intel.GetFullIntel(*target)
	if err != nil {
		output.Error("Failed to fetch node intelligence")
		return
	}

	// Render the HUD
	output.DisplayHUD(data)

	// Execute Tactical Probes if enabled
	if *scan {
		probes.RunFullScan(*target)
	}
}
