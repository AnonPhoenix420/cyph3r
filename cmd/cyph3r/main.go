package main

import (
	"flag"
	"os"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// REGISTER FLAGS (This fixes your "not defined" error)
	targetPtr := flag.String("target", "", "Target domain or IP address")
	phonePtr := flag.String("phone", "", "International phone number")
	scanPtr := flag.Bool("scan", false, "Enable tactical port scan")
	flag.Parse()

	output.Banner()

	// 1. Check for Phone Input
	if *phonePtr != "" {
		output.PulseNode(*phonePtr)
		pData, err := intel.GetPhoneIntel(*phonePtr)
		if err != nil {
			output.Error("Phone lookup failed.")
			os.Exit(1)
		}
		output.DisplayPhoneHUD(pData)
		os.Exit(0)
	}

	// 2. Check for Target Input
	if *targetPtr != "" {
		output.PulseNode(*targetPtr)
		data, err := intel.GetTargetIntel(*targetPtr)
		if err != nil {
			output.Error("Target resolution failed.")
			os.Exit(1)
		}

		output.DisplayHUD(data)

		// Run scan only if --scan flag is present
		if *scanPtr {
			probes.RunFullScan(*targetPtr)
		}
		
		output.Success("Operation Complete.")
		os.Exit(0)
	}

	// 3. Fallback if no flags
	output.Error("Missing input. Use --target <host> or --phone <number>")
	flag.Usage()
	os.Exit(1)
}
