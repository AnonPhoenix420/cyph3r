package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// Identity
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// ORGANIZATION DOX
	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "ENTITY_NAME:", White, data.Org)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "DESCRIPTION:", Gray, data.ISP)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "NETWORK_ASN:", NeonYellow, data.AS)
	fmt.Printf(" %s• %-15s %s%s\n", Cyan, "OWNER_LEAK:", NeonPink, "Admin/NOC Auth: "+data.Org)

	// GEO ENTITY
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s• %sLOCATION: %s%s, %s, %s [%s]\n", Cyan, White, NeonYellow, data.City, data.Region, data.Country, data.Zip)
	fmt.Printf(" %s• %sPOSITION: %s%.4f° N, %.4f° E\n", Cyan, White, Cyan, data.Lat, data.Lon)

	// INFRASTRUCTURE
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") { fmt.Printf(" %s[+] %-25s %s[ACTIVE]%s\n", NeonGreen, White+res, NeonBlue, Reset) }
	}
	fmt.Printf("\n %s[*] Software:   %s%s\n", Gray, White, data.WAFType)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "DEBUG") { fmt.Printf(" %s[*] %-30s %s[LEAK]%s\n", Gray, White+res, NeonPink, Reset) }
	}

	// REVERSE DNS
	if len(data.ReverseDNS) > 0 {
		fmt.Printf("\n%s[ REVERSE_DNS_PTR ]%s\n", NeonBlue, Reset)
		for _, ptr := range data.ReverseDNS { fmt.Printf(" %s[*] %s%s\n", Gray, White, ptr) }
	}

	// CLUSTERS
	if verbose {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			for _, ip := range ips { fmt.Printf(" %s[*] %-25s %s→ %s%-18s %s[NODE]%s\n", Gray, White+ns, Gray, Cyan, ip, Gray, Reset) }
		}
	}
	fmt.Println()
}
