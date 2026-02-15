package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_TRIANGULATION ] ---%s\n", NeonPink, Reset)
	
	status := "STABLE_UPLINK"
	if p.State == "Triangulating Region..." { status = "SCANNING_GLOBAL_DB" }

	fmt.Printf("%s[*] System Status:  %s%s\n", White, NeonGreen, status)
	fmt.Printf("%s[*] Target Number: %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Service:       %s%s\n", White, NeonGreen, p.Carrier)
	fmt.Printf("%s[*] Country:       %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Region:        %s%s\n", White, NeonGreen, p.State)

	fmt.Printf("\n%s[ 📡 OSINT_GEO_COORDINATES ]%s\n", NeonYellow, Reset)
	if p.Lat != "" {
		fmt.Printf("%s[*] Lat/Lon:       %s%s, %s\n", White, NeonPink, p.Lat, p.Lon)
		fmt.Printf("%s[*] Map Vector:    %s📍 %s\n", White, NeonBlue, p.MapLink)
	} else {
		fmt.Printf("%s[*] Vector:        %sSEARCHING GLOBAL PREFIXES...\n", White, NeonPink)
	}
	fmt.Printf("───────────────────────────────────────\n")
}
