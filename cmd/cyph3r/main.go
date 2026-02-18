package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	targetPtr := flag.String("t", "", "Target domain")
	verbosePtr := flag.Bool("v", false, "Verbose output")
	flag.Parse()

	output.Banner()
	
	// --- GHOST SHIELD INITIALIZATION ---
	fmt.Printf("\033[37m[*] INFO: Checking Shield Status... \033[0m")
	shield := intel.CheckShield()
	
	if !shield.IsActive {
		fmt.Printf("\033[31m[UNSECURED]\033[0m\n")
		fmt.Println("[!] OPSEC VIOLATION: Connect to Proton VPN to proceed.")
		os.Exit(1)
	}
	
	fmt.Printf("\033[32m[SECURE]\033[0m\n")
	fmt.Printf("\033[90m    ↳ Node: %s | ISP: %s\033[0m\n\n", shield.Location, shield.ISP)

	if *targetPtr != "" {
		runTargetScan(*targetPtr, *verbosePtr)
	}
}
