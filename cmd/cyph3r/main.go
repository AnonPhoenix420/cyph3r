package main

import (
	"flag"
	"os"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// 1. Setup target flag
	targetPtr := flag.String("target", "", "Target domain or IP address")
	flag.Parse()

	// 2. Clear UI and render banner
	output.Banner()

	// 3. Validate input
	if *targetPtr == "" {
		output.Error("No target specified. Use -target <domain/ip>")
		os.Exit(1)
	}

	// 4. Start Pulse (This is the function you were missing)
	output.PulseNode(*targetPtr)

	// 5. Fetch Remote Intel (No local data collected)
	data, err := intel.GetTargetIntel(*targetPtr)
	if err != nil {
		output.Error("Failed to resolve target intelligence.")
		os.Exit(1)
	}

	// 6. Display Remote Target HUD
	output.DisplayHUD(data)

	// 7. Execute Tactical Port Scan
	probes.RunFullScan(*targetPtr)

	// 8. Exit
	output.Success("Operation Complete.")
}
