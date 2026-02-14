package main

import (
	"flag"
	"cyph3r/internal/intel"
	"cyph3r/internal/output"
	"cyph3r/internal/probes"
)

func main() {
	target := flag.String("target", "", "Target domain/IP")
	scan := flag.Bool("scan", false, "Enable port scanning")
	flag.Parse()

	if *target == "" {
		output.Error("No target specified. Use -target <host>")
		return
	}

	output.PrintStatus("TARGET", *target)
	data, _ := intel.GetFullIntel(*target)
	output.DisplayHUD(data)

	if *scan {
		probes.RunFullScan(*target)
	}
}
