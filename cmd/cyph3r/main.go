package main

import (
	"flag"
	"fmt"
	"os"

	// Ensure these paths match your go.mod module name
	"github.com/AnonPhoenix420/cyph3r/internal/banner"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Define Command Line Flags
	targetPtr := flag.String("t", "", "Target domain (e.g., google.com)")
	phonePtr := flag.String("p", "", "Phone number (e.g., +13302354552)")
	verbosePtr := flag.Bool("v", false, "Enable verbose output (Raw JSON)")
	flag.Parse()

	// 2. Execute Banner from your banner.go
	banner.PrintBanner() 
	fmt.Println("\033[0;37m[*] INFO: Initializing Satellite Uplink...\033[0m")

	// 3. Vector Selection Logic
	if *targetPtr != "" {
		executeTargetVector(*targetPtr, *verbosePtr)
		os.Exit(0)
	}

	if *phonePtr != "" {
		executePhoneVector(*phonePtr)
		os.Exit(0)
	}

	// 4. Default Help Output
	fmt.Println("\n\033[0;31m[!] ERROR: No search vector provided.\033[0m")
	fmt.Println("Usage:")
	fmt.Println("  ./cyph3r -t <domain>   (Network Intelligence)")
	fmt.Println("  ./cyph3r -p <phone>    (Mobile Intelligence)")
	fmt.Println("  Options: -v            (Verbose/Raw Data)")
	os.Exit(1)
}

func executeTargetVector(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)

	data, err := intel.GetTargetIntel(target)
	done <- true // Kill animation

	if err != nil {
		fmt.Printf("\n\033[0;31m[!] UPLINK FAILURE: %v\033[0m\n", err)
		return
	}

	output.DisplayHUD(data, verbose)
}

func executePhoneVector(number string) {
	done := make(chan bool)
	go output.LoadingAnimation(done, number)

	data, err := intel.GetPhoneIntel(number)
	done <- true // Kill animation

	if err != nil {
		fmt.Printf("\n\033[0;31m[!] TRACE FAILURE: %v\033[0m\n", err)
		return
	}

	output.DisplayPhoneHUD(data)
}
