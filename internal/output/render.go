package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	// --- LOCALHOST IDENTITY (FOR NOTATION) ---
	fmt.Printf("\n%s[ LOCAL_HOST_IDENTITY ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Hostname:  %s%s\n", White, NeonBlue, data.LocalHost)
	for _, ip := range data.LocalIPs {
		fmt.Printf("%s[*] Local IP:  %s%s\n", White, NeonGreen, ip)
	}

	// --- FULL RECON HUD ---
	fmt.Printf("\n%s--- [ FULL_RECON_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Organization: %s%s\n", White, NeonGreen, data.Org)
	fmt.Printf("%s[*] Registrar:    %s%s\n", White, NeonGreen, data.Registrar)
	fmt.Printf("%s[*] Location:     %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.RegionName, data.Country, data.Zip)
	fmt.Printf("%s[*] Country Code: %s%s\n", White, NeonGreen, data.CountryCode)
	fmt.Printf("%s[*] Coordinates:  %s%f, %f\n", White, NeonGreen, data.Lat, data.Lon)

	// --- DNS / NAME SERVERS ---
	fmt.Printf("\n%s[ NAME_SERVERS ]%s\n", NeonYellow, Reset)
	for _, ns := range data.NameServers {
		fmt.Printf("%s[-] %s\n", White, ns)
	}

	// --- GEO-MAP LINK ---
	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f", data.Lat, data.Lon)
	fmt.Printf("\n%s[*] Geo-Map: %s%s%s\n", White, NeonBlue, mapURL, Reset)
}
