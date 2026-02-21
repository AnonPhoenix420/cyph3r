package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// 1. IDENTITY HEADER
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// 2. ORGANIZATION DOX (Conditional Entity Name)
	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	if data.Org != "" {
		fmt.Printf(" %s• %-15s %s%s\n", Cyan, "ENTITY_NAME:", White, data.Org)
	}
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "DESCRIPTION:", Gray, data.ISP)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "NETWORK_ASN:", NeonYellow, data.AS)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "TIMEZONE:", White, data.Timezone)

	// 3. GEO ENTITY (Conditional ZIP/ID)
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s, %s (%s)", data.City, data.RegionName, data.Country, data.CountryCode)
	if data.Zip != "" {
		loc += fmt.Sprintf(" [%s]", data.Zip)
	}
	fmt.Printf(" %s• %-12s %s%s\n", Cyan, "LOCATION:", NeonYellow, loc)
	fmt.Printf(" %s• %-12s %s%.4f° N, %.4f° E %s(ID: %s)%s\n", 
		Cyan, "POSITION:", Cyan, data.Lat, data.Lon, Gray, data.Region, Reset)

	// 4. INFRASTRUCTURE STACK (Green Active Check)
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	infra := "RESIDENTIAL"; if data.IsHosting { infra = "DATA_CENTER" }
	fmt.Printf(" %s[*] INFRA_TYPE: %s%s\n", Cyan, White, infra)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") {
			fmt.Printf(" %s[+] %-20s %s[ACTIVE]%s\n", NeonGreen, White+res, NeonGreen, Reset)
		}
	}
	fmt.Printf("\n %s[*] Software:   %s%s\n", Gray, White, data.WAFType)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "DEBUG") {
			fmt.Printf(" %s[*] %-30s %s[LEAK]%s\n", Gray, White+res, NeonPink, Reset)
		}
	}

	// 5. PTR & AUTHORITATIVE CLUSTERS
	if len(data.ReverseDNS) > 0 {
		fmt.Printf("\n%s[ REVERSE_DNS_PTR ]%s\n", NeonBlue, Reset)
		for _, ptr := range data.ReverseDNS {
			fmt.Printf(" %s[*] %s%s\n", Gray, White, ptr)
		}
	}

	if verbose {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s%s\n", NeonPink, White, ns)
			for _, ip := range ips {
				fmt.Printf("  %s↳ %-18s %s[ONLINE]%s\n", Cyan, ip, NeonGreen, Reset)
			}
		}
	}
	fmt.Println()
}
