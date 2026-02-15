package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"cyph3r/internal/intel"
	"cyph3r/internal/output"
)

func main() {
	// 1. Define Optional Flags
	targetFlag := flag.String("target", "", "Domain to scan")
	phoneFlag := flag.String("phone", "", "Phone to triangulate")
	_ = flag.Bool("scan", false, "Legacy scan mode") // Keeps -scan from crashing

	flag.Parse()

	// 2. REINSTATE TACTICAL DISPLAY (Your Branding)
	fmt.Printf("%s______      ____  __  __ _____ ____\n", output.NeonPink)
	fmt.Printf(" / ____/_  __/ __ \\/ / / /|__  // __ \\\n")
	fmt.Printf("/ /   / / / / /_/ / /_/ /  /_ </ /_/ /\n")
	fmt.Printf("/ /___/ /_/ / ____/ __  / ___/ / _, _/\n")
	fmt.Printf("\\____/\\__, /_/   /_/ /_/ /____/_/ |_|\n")
	fmt.Printf("     /____/         NETWORK_INTEL_SYSTEM%s\n", output.Reset)

	// 3. GET INPUT (Flag or Positional)
	var input string
	if *targetFlag != "" {
		input = *targetFlag
	} else if *phoneFlag != "" {
		input = *phoneFlag
	} else if flag.NArg() > 0 {
		input = flag.Arg(0) // Takes the "naked" input like: ./cyph3r google.com
	}

	if input == "" {
		fmt.Printf("\n%s[!] Error: No target provided.%s\n", output.NeonPink, output.Reset)
		fmt.Println("Usage: ./cyph3r <target> or ./cyph3r --target <target>")
		return
	}

	// 4. AUTOMATIC ROUTING
	// If it starts with '+' or is all numbers, treat as phone. Else, treat as domain.
	if strings.HasPrefix(input, "+") || isNumeric(input) {
		output.PulseNode(input)
		pData, _ := intel.GetPhoneIntel(input)
		output.DisplayPhoneHUD(pData)
	} else {
		output.PulseNode(input)
		data, _ := intel.GetTargetIntel(input)
		output.DisplayHUD(data)
	}
}

// Helper to detect if input is a phone number
func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(s) > 7 // Typical min length for a phone string
}
