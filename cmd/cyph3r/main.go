package main

import (
	"flag"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	target := flag.String("target", "", "Target domain or IP")
	phone := flag.String("phone", "", "Target phone number")
	scan := flag.Bool("scan", false, "Enable port scanning")
	flag.Parse()

	// 1. Check for Phone Intel first
	if *phone != "" {
		intel.GetPhoneIntel(*phone)
		return
	}

	// 2. Check for Domain Intel
	if *target == "" {
		output.Error("Input required. Use -target <domain> or -phone <number>")
		return
	}

	output.PulseNode(*target)

	data, err := intel.GetFullIntel(*target)
	if err != nil {
		output.Error("Failed to fetch intel")
		return
	}

	output.DisplayHUD(data)

	if *scan {
		probes.RunFullScan(*target)
	}
}
