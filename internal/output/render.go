package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData, verbose bool) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, NeonBlue)
	if data.IsWAF { fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, NeonBlue) }
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs { fmt.Printf(" ↳ %s[v]%s %-18s %s[LINK_ACTIVE]%s\n", Cyan, NeonBlue, NeonGreen, ip, NeonBlue, Reset) }

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	fmt.Printf(" • ENTITY:   %s%s\n", NeonYellow, data.Org)
	fmt.Printf(" • POSITION: %s%.4f° N, %.4f° E %s📡 (SIGNAL: %s)\n", Cyan, data.Lat, data.Lon, Amber, data.Latency)

	if verbose && len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", NeonBlue, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s\n", Gray, ns)
			for _, ip := range ips { fmt.Printf("  %s↳ %-30s %s[ONLINE]%s\n", Cyan, ip, NeonGreen, Reset) }
		}
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") { fmt.Printf("%s[+] %-25s %s[ACTIVE]%s\n", NeonGreen, res, NeonBlue, Reset)
		} else { fmt.Printf("%s[*] %-25s %s[]%s\n", NeonBlue, res, NeonBlue, Reset) }
	}

	if verbose && data.RawGeo != "" {
		fmt.Printf("\n%s[ RAW_METADATA ]%s\n%s%s%s\n", Gray, Reset, Gray, data.RawGeo, Reset)
	}
}
