package main

import (
	"flag"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// Define flags
	targetFlag := flag.String("target", "", "Domain or IP")
	phoneFlag := flag.String("phone", "", "Phone number")
	_ = flag.Bool("scan", false, "Legacy scan mode")
	flag.Parse()

	// 1. MATCHED TO YOUR LOGO: Calls func Banner() in banner.go
	output.Banner()

	// 2. CAPTURE INPUT (Flag OR Naked Argument)
	var input string
	if *targetFlag != "" {
		input = *targetFlag
	} else if *phoneFlag != "" {
		input = *phoneFlag
	} else if flag.NArg() > 0 {
		input = flag.Arg(0) // Handles: ./cyph3r google.com
	}

	// Exit silently if no input
	if input == "" {
		return
	}

	// 3. SMART ROUTING
	// Treat as phone if it starts with '+' or is a long numeric string
	if strings.HasPrefix(input, "+") || (len(input) > 7 && isNumeric(input)) {
		output.PulseNode(input)
		pData, _ := intel.GetPhoneIntel(input)
		output.DisplayPhoneHUD(pData)
	} else {
		// Treat as domain/IP
		output.PulseNode(input)
		data, _ := intel.GetTargetIntel(input)
		output.DisplayHUD(data)
	}
}

// isNumeric checks if the string is purely digits (ignoring the + prefix)
func isNumeric(s string) bool {
	clean := strings.TrimPrefix(s, "+")
	for _, char := range clean {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
