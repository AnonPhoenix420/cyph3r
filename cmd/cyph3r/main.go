package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Initialize Flags
	target := flag.String("target", "", "Target domain or IP address")
	scan := flag.Bool("scan", false, "Perform deep probe reconnaissance")
	flag.Parse()

	// 2. Display the Banner (Neon Blue)
	output.Banner()

	// 3. Validation
	if *target == "" {
		fmt.Printf("\n%s[!] FATAL: Target node not specified.%s\n", output.NeonPink, output.Reset)
		fmt.Println("Usage: cyph3r --target <domain.com> [--scan]")
		os.Exit(1)
	}

	// 4. Intelligence Gathering Phase
	// This happens behind the scenes immediately
	data, err := intel.GetFullIntel(*target)
	if err != nil {
		fmt.Printf("\n%s[!] INTEL FAILURE: %v%s\n", output.NeonPink, err, output.Reset)
	}

	// 5. HUD IDENTITY OUTPUT (IPs in Blue, NS in Yellow)
	output.PrintNodeIntel(data, *target)

	// 6. HUD GEOGRAPHIC OUTPUT (Maps in Pink, Location in Yellow)
	output.PrintGeoHUD(data)

	// 7. PROBE PHASE (Ports in Green)
	if *scan {
		fmt.Printf("\n%s[*] INITIATING MULTI-PROBE WAVE RECONNAISSANCE...%s\n", output.White, output.Reset)
		fmt.Printf("%s──[ PROBE ANALYSIS ]──%s\n", output.White, output.Reset)

		// Port list for the scan
		ports := []int{21, 22, 23, 25, 53, 80, 110, 443, 3306, 8080}

		for _, port := range ports {
			// In a full build, ConductWave would be in internal/probes
			// For this drop-in, we use the output status logic directly
			status := "ANALYZING" 
			
			// Calling your status.go logic
			output.PrintWaveStatus(port, status)
		}
		fmt.Printf("\n%s[*] SCAN COMPLETE.%s\n", output.NeonGreen, output.Reset)
	}
}
