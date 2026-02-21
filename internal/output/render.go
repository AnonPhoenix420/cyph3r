package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// Identity Box (Perfect Alignment Fix)
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	targetLine := fmt.Sprintf("[!] TARGET_NODE: %s", data.TargetName)
	fmt.Printf("\n║ %s%-61s %s║", Cyan, targetLine, NeonBlue)
	if data.IsWAF {
		shieldLine := fmt.Sprintf("[!] SHIELD:      %s", data.WAFType)
		fmt.Printf("\n║ %s%-61s %s║", Amber, shieldLine, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// TACTICAL TELEMETRY
	fmt.Printf("\n%s[ TACTICAL_TELEMETRY ]%s\n", Electric, Reset)
	usage := "RESIDENTIAL"
	if data.IsHosting { usage = "DATACENTER" }
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "USAGE:", NeonYellow, usage)
	
	mStatus := "OFF"
	if data.IsMobile { mStatus = "ON (CELLULAR)" }
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "MOBILE_NET:", White, mStatus)

	// GEO ENTITY
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s• %sLOCATION: %s%s, %s, %s\n", Cyan, White, NeonYellow, data.City, data.Region, data.Country)
	fmt.Printf(" %s• %sPOSITION: %s%.4f° N, %.4f° E %s📡 (%s%s%s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, NeonGreen, data.Latency, Amber)

	// INFRASTRUCTURE STACK (Port Scan & Header Leak)
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	
	// Print Ports First
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") {
			fmt.Printf(" %s[+] %-25s %s[ACTIVE]%s\n", NeonGreen, White+res, NeonBlue, Reset)
		}
	}
	
	fmt.Printf("\n %s[*] Software:   %s%s\n", Gray, White, data.WAFType)
	
	// Print Leaks
	for _, res := range data.ScanResults {
		if strings.Contains(res, "DEBUG") {
			fmt.Printf(" %s[*] %-30s %s[LEAK]%s\n", Gray, White+res, NeonPink, Reset)
		}
	}

	// NETWORK VECTORS & CLUSTERS (Same as before)
	if verbose {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			for _, ip := range ips {
				fmt.Printf(" %s[*] %-25s %s→ %s%-18s %s[NODE]%s\n", Gray, White+ns, Gray, Cyan, ip, Gray, Reset)
			}
		}
	}
	fmt.Println()
}
