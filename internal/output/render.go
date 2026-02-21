package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// --- MAIN IDENTITY BOX ---
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// --- NETWORK VECTORS ---
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs {
		fmt.Printf(" %s↳ [v] %-18s %s→ ---            %s[LINK_ACTIVE]%s\n", Cyan, White+ip, Gray, NeonGreen, Reset)
	}

	// --- GEO & TELEMETRY ---
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s• %sLOCATION: %s%s, %s, %s\n", Cyan, White, NeonYellow, data.City, data.Region, data.Country)
	fmt.Printf(" %s• %sPOSITION: %s%.4f° N, %.4f° E %s📡 (SIGNAL: %s%s%s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, NeonGreen, data.Latency, Amber)

	// --- INFRASTRUCTURE (With Leaked Debug Look) ---
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	if data.IsWAF {
		fmt.Printf(" %s[*] %sDEBUG: Node-ID [LEAKED]             %s[STATUS: OK]\n", Gray, White, NeonPink)
		fmt.Printf(" %s[*] %sSoftware: %-25s %s[]\n", NeonBlue, White, data.WAFType, Gray)
	}
	for _, res := range data.ScanResults {
		fmt.Printf(" %s[+] %-25s %s[ACTIVE]%s\n", NeonGreen, White+res, NeonBlue, Reset)
	}

	// --- DNS CLUSTERS (Verbose Tree) ---
	if verbose && len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s\n", Gray, White+ns)
			for _, ip := range ips {
				fmt.Printf("  %s↳ %-30s %s[NODE]%s\n", Cyan, ip, Gray, Reset)
			}
		}
	}
	fmt.Println()
}
