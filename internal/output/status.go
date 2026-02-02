package output

import (
	"fmt"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func PrintIntelHUD(target string, d *intel.NodeIntel) {
	currentTime := time.Now().Format("January 2, 2006, 3:04 pm")

	fmt.Println(CyanText("\n──[ NODE INTELLIGENCE ]──"))
	fmt.Printf(" %-15s %s\n", WhiteText("TARGET:"), YellowText(target))
	fmt.Printf(" %-15s %s\n", WhiteText("IP ADDRESSES:"), strings.Join(d.IPs, " | "))
	fmt.Printf(" %-15s %s\n", WhiteText("NS NODES:"), strings.Join(d.NSIPs, "\n                "))
	
	fmt.Println(CyanText("──[ GEOGRAPHIC METADATA ]──"))
	fmt.Printf(" %-15s %s (%s)\n", WhiteText("COUNTRY:"), d.Country, d.CountryCode)
	fmt.Printf(" %-15s %s (%s)\n", WhiteText("REGION:"), d.Region, d.RegionCode)
	fmt.Printf(" %-15s %s, %s\n", WhiteText("CITY/ZIP:"), d.City, d.Zip)
	fmt.Printf(" %-15s %s\n", WhiteText("TIMEZONE:"), d.TZ)
	fmt.Printf(" %-15s %s\n", WhiteText("SCAN DATE:"), currentTime)
	
	fmt.Println(CyanText("──[ NETWORK INFRASTRUCTURE ]──"))
	fmt.Printf(" %-15s %s\n", WhiteText("ISP:"), d.ISP)
	fmt.Printf(" %-15s %s\n", WhiteText("ORG:"), d.Org)
	fmt.Printf(" %-15s %s\n", WhiteText("ASN:"), d.ASN)
	fmt.Printf(" %-15s %.4f, %.4f\n", WhiteText("LOCATION:"), d.Lat, d.Lon)
	fmt.Println()
}

// This is the function that was missing causing the build error:
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
