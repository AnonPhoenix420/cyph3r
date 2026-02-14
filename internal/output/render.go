package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PulseNode initializes the UI for the target
func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

// DisplayHUD renders only remote target intelligence
func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	
	fmt.Printf("%s[*] Target Name:  %s%s\n", White, NeonBlue, data.TargetName)
	for _, tip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, tip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Organization:  %s%s\n", White, NeonGreen, data.Org)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.RegionName, data.Country, data.Zip)
	fmt.Printf("%s[*] Coordinates:   %s%f, %f\n", White, NeonGreen, data.Lat, data.Lon)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]%s\n", NeonYellow, Reset)
	for host, ip := range data.NameServers {
		fmt.Printf("%s[-] %-22s %s[%s]%s\n", White, host, NeonBlue, ip, Reset)
	}

	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f", data.Lat, data.Lon)
	fmt.Printf("\n%s[*] Geo-Map Link: %s%s%s\n", White, NeonBlue, mapURL, Reset)
}
