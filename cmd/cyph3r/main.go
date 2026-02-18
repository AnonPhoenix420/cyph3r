package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Tactical Argument Mapping
	targetPtr := flag.String("t", "", "Target domain (e.g., google.com)")
	phonePtr := flag.String("p", "", "Phone number (e.g., +13302354552)")
	verbosePtr := flag.Bool("v", false, "Enable verbose output (Raw JSON)")
	flag.Parse()

	// 2. Initialize Visual Uplink & Banner
	output.Banner()
	
	// 3. Pre-Flight Shield Check (Synchronized with status.go)
	shield := intel.CheckShield()
	output.PrintShieldStatus(shield.IsActive, shield.Location, shield.ISP)
	
	if !shield.IsActive {
		fmt.Printf("\n\033[31m[!] CRITICAL: VPN required for OPSEC. Connection Terminated.\033[0m\n")
		os.Exit(1)
	}

	// 4. Vector Execution Logic
	if *targetPtr != "" {
		runTargetScan(*targetPtr, *verbosePtr)
		os.Exit(0)
	}

	if *phonePtr != "" {
		runPhoneTrace(*phonePtr)
		os.Exit(0)
	}

	// 5. Default Failover
	fmt.Printf("\n\033[31m[!] ERROR: No search vector provided.\033[0m\n")
	fmt.Println("Usage:")
	fmt.Println(" ./cyph3r -t <domain> (Network Intelligence)")
	fmt.Println(" ./cyph3r -p <phone> (Mobile Intelligence)")
	os.Exit(1)
}

// runTargetScan bridges the engine to the HUD
func runTargetScan(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)
	
	data, err := intel.GetTargetIntel(target)
	done <- true 

	if err != nil {
		fmt.Printf("\n\033[31m[!] UPLINK FAILURE: %v\033[0m\n", err)
		return
	}

	output.DisplayHUD(data, verbose)
}

// runPhoneTrace bridges the engine to the Phone HUD
func runPhoneTrace(number string) {
	done := make(chan bool)
	go output.LoadingAnimation(done, number)
	
	data, err := intel.GetPhoneIntel(number)
	done <- true 

	if err != nil {
		fmt.Printf("\n\033[31m[!] TRACE FAILURE: %v\033[0m\n", err)
		return
	}
	
	output.DisplayPhoneHUD(data)
}
