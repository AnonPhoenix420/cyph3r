package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	// --- HEADER ---
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue)
	} else {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Gray, "UNPROTECTED / DIRECT_IP", NeonBlue)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// --- VECTORS ---
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs {
		fmt.Printf(" %s↳ %s[v]%s %-18s %s[LINK_ACTIVE]%s\n", Cyan, NeonBlue, NeonGreen, ip, NeonBlue, Reset)
	}

	// --- GEO ---
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)

	// --- STACK ---
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software:   %s%-25s %s[]%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "STACK: "), NeonBlue, Reset)
		} else {
			fmt.Printf("%s[*] %s%s\n", Electric, White, res)
		}
	}
}
