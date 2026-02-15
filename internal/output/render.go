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
	fmt.Printf("%s[*] Node:          %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}
	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] Coordinates:   %s%s, %s\n", White, NeonPink, data.Lat, data.Lon)
	fmt.Printf("%s[*] Map Vector:    %s📍 %s\n", White, NeonBlue, data.MapLink)

	fmt.Printf("\n%s[ NAME_SERVER_CLUSTER ]%s\n", NeonYellow, Reset)
	for host, ips := range data.NameServers {
		fmt.Printf("%s[-] %-20s\n", White, host)
		for _, ip := range ips {
			fmt.Printf("    %s↳ [%s]%s\n", NeonBlue, ip, Reset)
		}
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_TRIANGULATION ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Target Number: %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Country:       %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Region/State:  %s%s\n", White, NeonGreen, p.State)
	fmt.Printf("%s[*] Location:      %s%s\n", White, NeonGreen, p.Location)
	fmt.Printf("%s[*] Carrier:       %s%s\n", White, NeonGreen, p.Carrier)
	fmt.Printf("\n%s[ 📡 PINPOINT_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Coordinates:   %s%s, %s\n", White, NeonPink, p.Lat, p.Lon)
	fmt.Printf("%s[*] Map Vector:    %s📍 %s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("───────────────────────────────────────\n")
}
