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

	// 2. Initialize Visual Uplink
	output.Banner() 
	fmt.Printf("%s[*] INFO: Initializing Satellite Uplink...%s\n", output.White, output.Reset)

	// 3. Vector Execution Logic
	if *targetPtr != "" {
		runTargetScan(*targetPtr, *verbosePtr)
		os.Exit(0)
	}

	if *phonePtr != "" {
		runPhoneTrace(*phonePtr)
		os.Exit(0)
	}

	// 4. Default Failover
	fmt.Printf("\n%s[!] ERROR: No search vector provided.%s\n", output.Red, output.Reset)
	fmt.Println("Usage:")
	fmt.Println("  ./cyph3r -t <domain>   (Network Intelligence)")
	fmt.Println("  ./cyph3r -p <phone>    (Mobile Intelligence)")
	fmt.Println("  Options: -v            (Verbose/Raw Data)")
	os.Exit(1)
}

func runTargetScan(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)

	data, err := intel.GetTargetIntel(target)
	done <- true // Terminate animation

	if err != nil {
		fmt.Printf("\n%s[!] UPLINK FAILURE: %v%s\n", output.Red, err, output.Reset)
		return
	}

	output.DisplayHUD(data, verbose)
}

func runPhoneTrace(number string) {
	done := make(chan bool)
	go output.LoadingAnimation(done, number)

	data, err := intel.GetPhoneIntel(number)
	done <- true // Terminate animation

	if err != nil {
		fmt.Printf("\n%s[!] TRACE FAILURE: %v%s\n", output.Red, err, output.Reset)
		return
	}

	output.DisplayPhoneHUD(data)
}
