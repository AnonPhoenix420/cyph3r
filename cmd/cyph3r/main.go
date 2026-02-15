package main

import (
	"flag"
	"strings"
	"cyph3r/internal/intel"
	"cyph3r/internal/output"
)

func main() {
	// 1. Define flags so they don't error out
	targetFlag := flag.String("target", "", "Domain")
	phoneFlag := flag.String("phone", "", "Phone")
	_ = flag.Bool("scan", false, "Scan mode")
	flag.Parse()

	// 2. Call YOUR banner (Make sure this is Capitalized in banner.go)
	output.DisplayBanner()

	// 3. Logic to handle: ./cyph3r google.com OR ./cyph3r --target google.com
	input := ""
	if *targetFlag != "" {
		input = *targetFlag
	} else if *phoneFlag != "" {
		input = *phoneFlag
	} else if flag.NArg() > 0 {
		input = flag.Arg(0)
	}

	if input == "" {
		return
	}

	// 4. Router: Detect if it's a Phone or Domain
	if strings.HasPrefix(input, "+") || isNumeric(input) {
		output.PulseNode(input)
		p, _ := intel.GetPhoneIntel(input)
		output.DisplayPhoneHUD(p)
	} else {
		output.PulseNode(input)
		d, _ := intel.GetTargetIntel(input)
		output.DisplayHUD(d)
	}
}

func isNumeric(s string) bool {
	clean := strings.TrimPrefix(s, "+")
	if len(clean) < 7 { return false }
	for _, c := range clean {
		if c < '0' || c > '9' { return false }
	}
	return true
}
