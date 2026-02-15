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
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	for _, tip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, tip)
	}
	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.RegionName, data.Country, data.Zip)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonGreen, data.ISP)
	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]%s\n", NeonYellow, Reset)
	for host, ips := range data.NameServers {
		fmt.Printf("%s[-] %-20s\n", White, host)
		for _, ip := range ips {
			fmt.Printf("    %s↳ [%s]%s\n", NeonBlue, ip, Reset)
		}
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ SATELLITE_TRIANGULATION_REPORT ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Status:       %sUPLINK_STABLE\n", White, NeonGreen)
	fmt.Printf("%s[*] Target Num:   %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Country:      %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Geolocation:  %s%s\n", White, NeonGreen, p.Location)
	fmt.Printf("%s[*] Carrier:      %s%s\n", White, NeonGreen, p.Carrier)
	
	fmt.Printf("\n%s[ 📡 PINPOINT_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Map Vector:   %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("───────────────────────────────────────\n")
}
