package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// Define Flags
	target := flag.String("target", "", "Domain or IP address to scan")
	phone := flag.String("phone", "", "Phone number to triangulate (Global Format)")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: cyph3r [options]\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Logic for Network Target
	if *target != "" {
		output.PulseNode(*target)
		data, err := intel.GetTargetIntel(*target)
		if err != nil {
			fmt.Printf("\n[!] Error during target scan: %v\n", err)
			return
		}
		output.DisplayHUD(data)
	}

	// Logic for Phone Target
	if *phone != "" {
		output.PulseNode(*phone)
		pData, err := intel.GetPhoneIntel(*phone)
		if err != nil {
			fmt.Printf("\n[!] Error during phone triangulation: %v\n", err)
			return
		}
		output.DisplayPhoneHUD(pData)
	}

	// Default help if no arguments
	if *target == "" && *phone == "" {
		flag.Usage()
	}
}
