package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ... (Previous animation/color codes remain same)

func DisplayHUD(data models.IntelData, verbose bool) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-41s %s║", Cyan, NeonPink+data.TargetName, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" • ENTITY:   %s%s\n", NeonYellow, data.Org)
	fmt.Printf(" • POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)
	fmt.Printf(" • Location: %s%s, %s%s\n", NeonGreen, data.City, data.Country, Reset)

	if verbose {
		fmt.Printf("\n%s[ RAW_GEO_DATA ]%s\n%s%s%s\n", Red, Reset, Amber, data.RawGeo, Reset)
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software: %s%s%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "STACK: "), Reset)
		} else {
			fmt.Printf("%s[+] %s\n", NeonGreen, res)
		}
	}
}
