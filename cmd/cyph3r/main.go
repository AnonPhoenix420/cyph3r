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
	phonePtr := flag.String("p", "", "Phone number")
	verbosePtr := flag.Bool("v", false, "Enable verbose mode")
	flag.Parse()

	// Call the banner from the output package
	output.PrintBanner()
	fmt.Println("\033[0;37m[*] INFO: Initializing Satellite Uplink...\033[0m")

	if *targetPtr != "" {
		executeTarget(*targetPtr, *verbosePtr)
		os.Exit(0)
	}

	if *phonePtr != "" {
		executePhone(*phonePtr)
		os.Exit(0)
	}

	fmt.Println("\n\033[0;31m[!] ERROR: Search vector required (-t or -p)\033[0m")
	os.Exit(1)
}

func executeTarget(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)
	data, err := intel.GetTargetIntel(target)
	done <- true
	if err != nil {
		fmt.Printf("\n[!] UPLINK ERROR: %v\n", err)
		return
	}
	output.DisplayHUD(data, verbose)
}

func executePhone(number string) {
	done := make(chan bool)
	go output.LoadingAnimation(done, number)
	data, err := intel.GetPhoneIntel(number)
	done <- true
	if err != nil {
		fmt.Printf("\n[!] TRACE ERROR: %v\n", err)
		return
	}
	output.DisplayPhoneHUD(data)
}
