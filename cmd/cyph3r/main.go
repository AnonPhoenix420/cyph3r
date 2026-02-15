package main

import (
	"flag"
	"strings"
	"cyph3r/internal/intel"
	"cyph3r/internal/output"
)

func main() {
	// Define optional flags for backwards compatibility
	targetFlag := flag.String("target", "", "Domain or IP")
	phoneFlag := flag.String("phone", "", "Phone number")
	_ = flag.Bool("scan", false, "Legacy scan mode")
	flag.Parse()

	// 1. CALL YOUR BANNER FROM banner.go
	output.DisplayBanner()

	// 2. CAPTURE INPUT (Flag OR Naked Argument)
	input := ""
	if *targetFlag != "" {
		input = *targetFlag
	} else if *phoneFlag != "" {
		input = *phoneFlag
	} else if flag.NArg() > 0 {
		input = flag.Arg(0) // Handles: ./cyph3r google.com
	}

	// Exit if no input at all
	if input == "" {
		return
	}

	// 3. SMART ROUTING
	// Treat as phone if it starts with '+' or is a long string of numbers
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

// Helper for numeric detection
func isNumeric(s string) bool {
	clean := strings.TrimPrefix(s, "+")
	for _, char := range clean {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
