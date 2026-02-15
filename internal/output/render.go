package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	
	fmt.Printf("%s[*] Target Name:  %s%s\n", White, NeonBlue, data.TargetName)
	for _, tip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, tip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Organization:  %s%s\n", White, NeonGreen, data.Org)
	fmt.Printf("%s[*] ISP:           %s%s\n", White, NeonGreen, data.ISP)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.RegionName, data.Country, data.Zip)
	fmt.Printf("%s[*] Coordinates:   %s%f, %f\n", White, NeonGreen, data.Lat, data.Lon)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]%s\n", NeonYellow, Reset)
	for host, ip := range data.NameServers {
		fmt.Printf("%s[-] %-22s %s[%s]%s\n", White, host, NeonBlue, ip, Reset)
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ PHONE_INTELLIGENCE_REPORT ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Number:       %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Location:     %s%s\n", White, NeonGreen, p.Location)
	fmt.Printf("%s[*] Carrier:      %s%s\n", White, NeonGreen, p.Carrier)
	fmt.Printf("%s[*] Line Type:    %s%s\n", White, NeonGreen, p.Type)
	fmt.Printf("───────────────────────────────────────\n")
}
