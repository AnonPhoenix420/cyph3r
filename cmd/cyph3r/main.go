package main

import (
	"flag"
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	output.Banner()

	target := flag.String("target", "", "Target Domain or IP")
	scan := flag.Bool("scan", false, "Enable Multi-Probe Wave Recon")
	flag.Parse()

	if *target == "" {
		fmt.Println("\033[31m[!] Error: No target specified. Use --target\033[0m")
		return
	}

	// 1. Calibration Sequence
	output.ScanAnimation()

	// 2. Gather Intelligence
	data, _ := intel.GetFullIntel(*target)

	// 3. 🛰️ DISPLAY IDENTITY (NS & IP) FIRST
	output.PrintNodeIntel(data, *target)
	
	// 4. DISPLAY LOCATION (Map Link)
	output.PrintGeoHUD(data.City, data.Country, data.Lat, data.Lon)

	// 5. 🌊 INITIATE PROBE WAVE (Last)
	if *scan {
		fmt.Println("\n[*] Initiating Multi-Probe Wave Reconnaissance...")
		fmt.Println("──[ PROBE ANALYSIS ]──")
		
		ports := []int{21, 22, 23, 25, 53, 80, 443, 8080}
		for _, p := range ports {
			method, status, convo := probes.ConductWave(*target, p)
			output.PrintWaveStatus(p, method, status, convo)
		}
	}
}
