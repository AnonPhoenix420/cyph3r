package main

import (
	"flag"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	targetFlag := flag.String("target", "", "Domain")
	phoneFlag := flag.String("phone", "", "Phone")
	_ = flag.Bool("scan", false, "Legacy")
	flag.Parse()

	output.DisplayBanner()

	input := ""
	if *targetFlag != "" { input = *targetFlag } else if *phoneFlag != "" { input = *phoneFlag } else if flag.NArg() > 0 { input = flag.Arg(0) }

	if input == "" { return }

	if strings.HasPrefix(input, "+") || (len(input) > 7 && isNumeric(input)) {
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
	for _, c := range clean { if c < '0' || c > '9' { return false } }
	return true
}
