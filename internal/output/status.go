package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

// PrintIntelHUD handles the OSINT display
func PrintIntelHUD(target string, ips []string, isp, loc string) {
	fmt.Println(CyanText("\n──[ NODE INTELLIGENCE ]──"))
	fmt.Printf(" %-15s %s\n", WhiteText("TARGET:"), YellowText(target))
	fmt.Printf(" %-15s %s\n", WhiteText("IP ADDRESS:"), fmt.Sprintf("%v", ips))
	fmt.Printf(" %-15s %s\n", WhiteText("ISP/ORG:"), isp)
	fmt.Printf(" %-15s %s\n", WhiteText("LOCATION:"), loc)
}

// PrintPortScan handles the Scanner display
func PrintPortScan(results []probes.ScanResult) {
	fmt.Println(CyanText("──[ OPEN SERVICES ]──"))
	if len(results) == 0 {
		fmt.Printf(" %s\n", YellowText("Scanning complete: No open common ports found."))
		return
	}

	for _, res := range results {
		fmt.Printf(" %-15s %-10d %s\n", WhiteText("PORT:"), res.Port, GreenText("[OPEN]"))
	}
	fmt.Println()
}
