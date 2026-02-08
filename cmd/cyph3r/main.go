package main

import (
	"flag"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	target := flag.String("target", "", "Target domain/IP")
	scan := flag.Bool("scan", false, "Enable port scan")
	flag.Parse()

	if *target == "" {
		output.Error("Target required. Use -target <domain>")
		return
	}

	output.PrintBanner()
	output.PulseNode(*target)

	data, err := intel.GetFullIntel(*target)
	if err != nil {
		output.Error("Intel retrieval failed")
		return
	}

	output.DisplayHUD(data)

	if *scan {
		probes.RunFullScan(*target)
	}
}
