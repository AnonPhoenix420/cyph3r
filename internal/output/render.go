package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// drawBoxLine handles surgical alignment by calculating visible text length
func drawBoxLine(label, value, labelCol, valCol string) {
	visibleText := fmt.Sprintf("[!] %s: %s", label, value)
	width := 61 // Total inner width of the box
	padding := width - len(visibleText)
	if padding < 0 { padding = 0 }
	
	fmt.Printf("\n║ %s[!] %s: %s%s%s %s║", labelCol, label, valCol, value, strings.Repeat(" ", padding), NeonBlue)
}

func DisplayHUD(data models.IntelData, verbose bool) {
	// 1. IDENTITY HEADER
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	
	drawBoxLine("TARGET_NODE", data.TargetName, Cyan, NeonPink)
	
	v4 := "NOT_DETECTED"; if len(data.TargetIPs) > 0 { v4 = data.TargetIPs[0] }
	drawBoxLine("TARGET_IPv4", v4, Amber, NeonGreen)

	v6 := "NOT_DETECTED"; if len(data.TargetIPv6s) > 0 { v6 = data.TargetIPv6s[0] }
	drawBoxLine("TARGET_IPv6", v6, Amber, Cyan)

	if data.IsWAF {
		drawBoxLine("SHIELD     ", data.WAFType, Amber, NeonYellow)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// 2. ORGANIZATION DOX
	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	if data.Org != "" { fmt.Printf(" %s• %-15s %s%s\n", Amber, "ENTITY_NAME:", NeonGreen, data.Org) }
	fmt.Printf(" %s• %-15s %s%s\n", Amber, "DESCRIPTION:", Gray, data.ISP)
	fmt.Printf(" %s• %-15s %s%s\n", Amber, "NETWORK_ASN:", NeonYellow, data.AS)
	fmt.Printf(" %s• %-15s %s%s\n", Amber, "TIMEZONE:", NeonGreen, data.Timezone)

	// 3. GEO ENTITY
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s, %s (%s)", data.City, data.RegionName, data.Country, data.CountryCode)
	if data.Zip != "" { loc += fmt.Sprintf(" [%s]", data.Zip) }
	fmt.Printf(" %s• %-12s %s%s\n", Amber, "LOCATION:", NeonYellow, loc)
	
	// POSITION + REGION ID (The "23" for Tehran)
	fmt.Printf(" %s• %-12s %s%.4f° N, %.4f° E %s(ID: %s)%s\n", 
		Amber, "POSITION:", Cyan, data.Lat, data.Lon, Gray, data.Region, Reset)

	// 4. INFRASTRUCTURE STACK
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	infra := "RESIDENTIAL"; if data.IsHosting || data.IsWAF { infra = "DATA_CENTER / CLOUD_PROXY" }
	fmt.Printf(" %s[*] INFRA_TYPE: %s%s\n", Amber, NeonGreen, infra)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") {
			fmt.Printf(" %s[+] %-13s %s[ACTIVE]%s\n", NeonGreen, NeonYellow+res, NeonGreen, Reset)
		}
	}
	fmt.Printf("\n %s[*] Software:   %s%s\n", Gray, NeonGreen, data.WAFType)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "DEBUG") {
			fmt.Printf(" %s[*] %-30s %s[LEAK]%s\n", Gray, NeonGreen+res, NeonPink, Reset)
		}
	}

	// 5. PTR & AUTHORITATIVE CLUSTERS
	if len(data.ReverseDNS) > 0 {
		fmt.Printf("\n%s[ REVERSE_DNS_PTR ]%s\n", NeonBlue, Reset)
		for _, ptr := range data.ReverseDNS { fmt.Printf(" %s[*] %s%s\n", Gray, NeonGreen, ptr) }
	}

	if verbose {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s%s\n", NeonPink, NeonGreen, ns)
			for _, ip := range ips { fmt.Printf("  %s↳ %-18s %s[ONLINE]%s\n", Cyan, ip, NeonGreen, Reset) }
		}
	}
	fmt.Println()
}
