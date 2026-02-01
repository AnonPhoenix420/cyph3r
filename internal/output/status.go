package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func PrintIntelHUD(target string, ips, ns []string, isp, loc string, lat, lon float64) {
	fmt.Println(CyanText("\n──[ NODE INTELLIGENCE ]──"))
	fmt.Printf(" %-15s %s\n", WhiteText("TARGET:"), YellowText(target))
	fmt.Printf(" %-15s %s\n", WhiteText("IPs:"), strings.Join(ips, ", "))
	fmt.Printf(" %-15s %s\n", WhiteText("NS SERVERS:"), strings.Join(ns, ", "))
	fmt.Printf(" %-15s %s\n", WhiteText("ISP/ORG:"), isp)
	fmt.Printf(" %-15s %s\n", WhiteText("LOCATION:"), loc)
	fmt.Printf(" %-15s %s\n", WhiteText("COORDS:"), fmt.Sprintf("%.4f, %.4f", lat, lon))
}

func PrintPortScan(results []probes.ScanResult) {
	fmt.Println(CyanText("──[ OPEN SERVICES ]──"))
	if len(results) == 0 {
		fmt.Printf(" %s\n", YellowText("No common ports open."))
		return
	}
	for _, res := range results {
		fmt.Printf(" %-15s %-10d %s\n", WhiteText("PORT:"), res.Port, GreenText("[OPEN]"))
	}
	fmt.Println()
}
