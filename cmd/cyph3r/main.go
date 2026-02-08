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

	// 2. Clear terminal and Deploy Banner (Neon Blue)
	output.Banner()

	// 3. Validation Guard
	if *target == "" {
		fmt.Printf("\n%s[!] FATAL: Access Denied. Target Node required.%s\n", output.NeonPink, output.Reset)
		fmt.Println("Usage: cyph3r --target <host.com> [--scan]")
		os.Exit(1)
	}

	// 4. Initialization Pulse (Neon Blue Animation)
	output.ScanAnimation()
	output.PulseNode(*target)

	// 5. Intelligence Retrieval (DNS & Geo-IP)
	// This populates our models with all the tech intelligence
	data, err := intel.GetFullIntel(*target)
	if err != nil {
		fmt.Printf("\n%s[!] INTEL_ERROR: %v%s\n", output.NeonPink, err, output.Reset)
	}

	// 6. HUD Execution (Identity -> Geography)
	// Displays Neon Blue IPs, Neon Yellow NS, and Neon Pink Links
	output.PrintNodeIntel(data, *target)
	output.PrintGeoHUD(data)

	// 7. Concurrent Scanning Phase (Neon Green)
	if *scan {
		// RunFullScan uses the high-speed concurrency logic from scanner.go
		probes.RunFullScan(*target)
		
		fmt.Printf("\n%s[*] RECONNAISSANCE SEQUENCE TERMINATED SUCCESSFULLY.%s\n", output.NeonGreen, output.Reset)
	}
}
