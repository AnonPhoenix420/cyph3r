package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	target := flag.String("t", "", "Target domain")
	phone := flag.String("p", "", "Phone number")
	flag.Parse()

	output.Banner()

	if *target != "" {
		executeTargetIntel(*target)
	} else if *phone != "" {
		executePhoneIntel(*phone)
	} else {
		output.Error("No target specified. Use -t <domain> or -p <phone>")
		os.Exit(1)
	}
}

func executeTargetIntel(target string) {
	output.PulseNode(target)
	done := make(chan bool)
	go output.LoadingAnimation(done, "Remote Node Intelligence")

	data, err := intel.GetTargetIntel(target)
	done <- true

	if err != nil {
		output.Error(fmt.Sprintf("Intel Extraction Failed: %v", err))
		return
	}
	output.DisplayHUD(data)
}

func executePhoneIntel(number string) {
	output.Info("Initializing Satellite Uplink...")
	done := make(chan bool)
	go output.LoadingAnimation(done, "Digital Footprint Analysis")

	data, err := intel.GetPhoneIntel(number)
	done <- true

	if err != nil {
		output.Error(fmt.Sprintf("Phone OSINT Failed: %v", err))
		return
	}
	output.DisplayPhoneHUD(data)
}
