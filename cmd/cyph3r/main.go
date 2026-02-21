package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Tactical Argument Mapping (RECON + GHOST TESTS)
	targetPtr := flag.String("t", "", "Target domain (Network Intelligence)")
	phonePtr  := flag.String("p", "", "Phone number (Mobile Intelligence)")
	testVec   := flag.String("test", "", "Tactical Vector: SYN, UDP, HULK")
	pps       := flag.Int("pps", 10, "Speed: Packets/Requests Per Second")
	duration  := flag.Int("time", 30, "Test duration in seconds")
	verbose   := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	// 2. Initialize Visual Uplink (Banner)
	output.Banner()

	// 3. Pre-Flight GHOST_MODE Shield Check
	shield := intel.CheckShield()
	output.PrintShieldStatus(shield.IsActive, shield.Location, shield.ISP)

	if !shield.IsActive {
		fmt.Printf("\n\033[31m[!] GHOST_MODE FAILURE: VPN Connection required. ABORTING.\033[0m\n")
		os.Exit(1)
	}

	// 4. Vector Execution Logic
	if *targetPtr != "" {
		// Mandatory Intel Phase
		data := runTargetScan(*targetPtr, *verbose)

		// Check if Tactical Test is requested
		if *testVec != "" {
			executeTacticalGhostTest(*targetPtr, *testVec, *pps, *duration)
		}
		os.Exit(0)
	}

	if *phonePtr != "" {
		runPhoneTrace(*phonePtr)
		os.Exit(0)
	}

	// 5. Default Failover
	flag.Usage()
	os.Exit(1)
}

func runTargetScan(target string, verbose bool) interface{} {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)

	data, err := intel.GetTargetIntel(target)
	done <- true
	if err != nil {
		fmt.Printf("\n\033[31m[!] UPLINK FAILURE: %v\033[0m\n", err)
		os.Exit(1)
	}
	output.DisplayHUD(data, verbose)
	return data
}

func executeTacticalGhostTest(target, vector string, speed, seconds int) {
	// Create a context that auto-terminates to prevent "Self-Suicide"
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	cfg := intel.TacticalConfig{
		Target: target,
		Vector: vector,
		PPS:    speed,
	}

	// Hand over to the Scrubbed Tactical Engine
	intel.RunTacticalTest(cfg, ctx)
}

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
