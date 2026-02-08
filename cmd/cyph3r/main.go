package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func main() {
	// 1. Initialize System Flags
	target := flag.String("target", "", "Target domain or IP address")
	scan := flag.Bool("scan", false, "Execute multi-probe wave reconnaissance")
	flag.Parse()

	// 2. Deployment - Display Identity Banner
	output.Banner()

	// 3. Validation Guard
	if *target == "" {
		fmt.Printf("\n%s[!] FATAL: Target Node required.%s\n", output.NeonPink, output.Reset)
		fmt.Println("Usage: cyph3r --target <host.com> [--scan]")
		os.Exit(1)
	}

	// 4. Initialization Sequence (Pulse Animation)
	output.ScanAnimation()
	output.PulseNode(*target)

	// 5. Intelligence Retrieval Phase
	// Aggregates DNS, WHOIS, and Geo-IP data into the central model
	data, err := intel.GetFullIntel(*target)
	if err != nil {
		fmt.Printf("\n%s[!] INTEL_ERROR: %v%s\n", output.NeonPink, err, output.Reset)
	}

	// 6. HUD Execution (Visual Data Rendering)
	// Displays Neon Blue Identity
	output.PrintNodeIntel(data, *target)
	
	// Displays Neon Yellow Registry Info (Registrar, Created Date)
	output.PrintRegistryHUD(data)
	
	// Displays Neon Yellow Geo-data and Neon Pink Map Link
	output.PrintGeoHUD(data)

	// 7. Concurrent Scanning Phase
	if *scan {
		// Orchestrates the high-speed TCP handshake probes
		probes.RunFullScan(*target)
		
		fmt.Printf("\n%s[*] RECONNAISSANCE SEQUENCE TERMINATED SUCCESSFULLY.%s\n", output.NeonGreen, output.Reset)
	}
}
