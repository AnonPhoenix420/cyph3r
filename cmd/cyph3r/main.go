package main

import (
	"flag"
	"os"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// 1. Define Flags
	targetPtr := flag.String("target", "", "Target domain or IP address")
	flag.Parse()

	// 2. Initial UI
	output.Banner()

	if *targetPtr == "" {
		output.Error("No target specified. Use -target <domain/ip>")
		os.Exit(1)
	}

	// 3. Identification Phase
	output.PulseNode(*targetPtr)

	// 4. Intelligence Gathering (Renamed to match intel.go)
	data, err := intel.GetTargetIntel(*targetPtr)
	if err != nil {
		output.Error("Failed to resolve target intelligence.")
		os.Exit(1)
	}

	// 5. Display the Full Recon HUD
	output.DisplayHUD(data)

	// 6. Execution Phase (Port/Signal Scan)
	probes.RunFullScan(*targetPtr)

	output.Success("Operation Complete.")
}
