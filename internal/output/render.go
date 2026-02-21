package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// 1. IDENTITY BOX
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	targetLine := fmt.Sprintf("[!] TARGET_NODE: %s", data.TargetName)
	fmt.Printf("\n║ %s%-61s %s║", Cyan, targetLine, NeonBlue)
	if data.IsWAF {
		shieldLine := fmt.Sprintf("[!] SHIELD:      %s", data.WAFType)
		fmt.Printf("\n║ %s%-61s %s║", Amber, shieldLine, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// 2. ORGANIZATION DOX (Pink/White)
	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "ENTITY_NAME:", White, data.Org)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "DESCRIPTION:", Gray, data.ISP)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "NETWORK_ASN:", NeonYellow, data.AS)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "TIMEZONE:", White, data.Timezone)

	// 3. GEO ENTITY (Synced Orange Fields)
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s• %-12s %s%s, %s, %s (%s) [%s]\n", 
		Cyan, "LOCATION:", NeonYellow, data.City, data.RegionName, data.Country, data.CountryCode, data.Zip)
	fmt.Printf(" %s• %-12s %s%.4f° N, %.4f° E %s(REGION_ID: %s)%s\n", 
		Cyan, "POSITION:", Cyan, data.Lat, data.Lon, Gray, data.Region, Reset)

	// 4. INFRASTRUCTURE STACK (Green Status)
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	infraType := "RESIDENTIAL"; if data.IsHosting { infraType = "DATA_CENTER" }
	fmt.Printf(" %s[*] INFRA_TYPE: %s%s\n", Cyan, White, infraType)
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

	// 5. REVERSE DNS PTR
	if len(data.ReverseDNS) > 0 {
		fmt.Printf("\n%s[ REVERSE_DNS_PTR ]%s\n", NeonBlue, Reset)
		for _, ptr := range data.ReverseDNS {
			fmt.Printf(" %s[*] %s%s\n", Gray, White, ptr)
		}
	}

	// 6. NAMESERVER CLUSTERS (Green Online Status)
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
