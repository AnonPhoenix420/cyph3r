package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// --- TACTICAL TELEMETRY (Utilizing Electric) ---
	fmt.Printf("\n%s[ TACTICAL_TELEMETRY ]%s\n", Electric, Reset)
	mStatus, pStatus, hStatus := "OFF", "OFF", "OFF"
	if data.IsMobile { mStatus = "ON (CELLULAR)" }
	if data.IsProxy { pStatus = "ON (VPN/PROXY)" }
	if data.IsHosting { hStatus = "ON (DATACENTER)" }
	
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "MOBILE_NET:", White, mStatus)
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "PROXY_NODE:", White, pStatus)
	fmt.Printf(" %s• %-12s %s%s\n", Electric, "HOSTING:", White, hStatus)

	// --- NETWORK VECTORS ---
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs {
		fmt.Printf(" %s↳ [v] %-18s %s→ %s[LINK_ACTIVE]%s\n", Cyan, White+ip, Gray, NeonGreen, Reset)
	}

	// --- UNIFORM CLUSTER STRINGS ---
	if verbose && len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			for _, ip := range ips {
				fmt.Printf(" %s[*] %-28s %s→ %s%-15s %s[NODE]%s\n", Gray, White+ns, Gray, Cyan, ip, Gray, Reset)
			}
		}
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	if data.IsWAF {
		fmt.Printf(" %s[*] %sDEBUG: Node-ID [LEAKED]             %s[STATUS: OK]\n", Gray, White, NeonPink)
	}
	fmt.Println()
}
