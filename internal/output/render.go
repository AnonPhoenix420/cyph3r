package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// 1. IDENTITY BOX (Border Fix: ANSI-Aware Padding)
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	
	// We use %-42s on the string before the color is applied to keep the border locked
	targetLine := fmt.Sprintf("[!] TARGET_NODE: %s", data.TargetName)
	fmt.Printf("\n║ %s%-61s %s║", Cyan, targetLine, NeonBlue)
	
	if data.IsWAF {
		shieldLine := fmt.Sprintf("[!] SHIELD:      %s", data.WAFType)
		fmt.Printf("\n║ %s%-61s %s║", Amber, shieldLine, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// 2. TACTICAL TELEMETRY
	fmt.Printf("\n%s[ TACTICAL_TELEMETRY ]%s\n", Electric, Reset)
	mStatus, pStatus, hStatus := "OFF", "OFF", "OFF"
	if data.IsMobile { mStatus = "ON (CELLULAR)" }
	if data.IsProxy { pStatus = "ON (VPN/PROXY)" }
	if data.IsHosting { hStatus = "ON (DATACENTER)" }
	
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "MOBILE_NET:", White, mStatus)
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "PROXY_NODE:", White, pStatus)
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "HOSTING:", White, hStatus)

	// 3. GEO INTEL (Restored)
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s• %sLOCATION: %s%s, %s, %s\n", Cyan, White, NeonYellow, data.City, data.Region, data.Country)
	fmt.Printf(" %s• %sPOSITION: %s%.4f° N, %.4f° E %s📡 (%s%s%s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, NeonGreen, data.Latency, Amber)
	fmt.Printf(" %s• %sOPERATOR: %s%s\n", Cyan, White, Gray, data.Org)

	// 4. NETWORK VECTORS
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs {
		fmt.Printf(" %s↳ [v] %-18s %s→ %s[LINK_ACTIVE]%s\n", Cyan, White+ip, Gray, NeonGreen, Reset)
	}

	// 5. UNIFORM CLUSTERS
	if verbose && len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			for _, ip := range ips {
				fmt.Printf(" %s[*] %-25s %s→ %s%-18s %s[NODE]%s\n", Gray, White+ns, Gray, Cyan, ip, Gray, Reset)
			}
		}
	}

	// 6. INFRASTRUCTURE STACK
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	if data.IsWAF {
		fmt.Printf(" %s[*] %sDEBUG: Node-ID [LEAKED]             %s[STATUS: OK]%s\n", Gray, White, NeonPink, Reset)
	}
	fmt.Println()
}
