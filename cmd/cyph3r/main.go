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
	// Define optional flags
	targetFlag := flag.String("target", "", "Domain to scan")
	phoneFlag := flag.String("phone", "", "Phone to triangulate")
	_ = flag.Bool("scan", false, "Legacy scan mode")

	flag.Parse()

	// 1. REINSTATE TACTICAL DISPLAY
	fmt.Printf("%s______      ____  __  __ _____ ____\n", output.NeonPink)
	fmt.Printf(" / ____/_  __/ __ \\/ / / /|__  // __ \\\n")
	fmt.Printf("/ /   / / / / /_/ / /_/ /  /_ </ /_/ /\n")
	fmt.Printf("/ /___/ /_/ / ____/ __  / ___/ / _, _/\n")
	fmt.Printf("\\____/\\__, /_/   /_/ /_/ /____/_/ |_|\n")
	fmt.Printf("     /____/         NETWORK_INTEL_SYSTEM%s\n", output.Reset)

	// 2. GET INPUT (Flag or Naked Positional)
	var input string
	if *targetFlag != "" {
		input = *targetFlag
	} else if *phoneFlag != "" {
		input = *phoneFlag
	} else if flag.NArg() > 0 {
		input = flag.Arg(0) 
	}

	if input == "" {
		fmt.Printf("\n%s[!] Error: No target provided.%s\n", output.NeonPink, output.Reset)
		return
	}

	// 3. AUTOMATIC ROUTING
	if strings.HasPrefix(input, "+") || isNumeric(input) {
		output.PulseNode(input)
		pData, err := intel.GetPhoneIntel(input)
		if err == nil {
			output.DisplayPhoneHUD(pData)
		}
	} else {
		output.PulseNode(input)
		data, err := intel.GetTargetIntel(input)
		if err == nil {
			output.DisplayHUD(data)
		}
	}
}

func isNumeric(s string) bool {
	clean := strings.TrimPrefix(s, "+")
	for _, char := range clean {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(clean) > 7
}
