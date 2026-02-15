package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PulseNode is now exported (Starts with Capital P)
func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Node:          %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_INTELLIGENCE ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] ISP:           %s%s\n", White, NeonGreen, data.ISP)
	fmt.Printf("%s[*] Organization:  %s%s\n", White, NeonGreen, data.Org)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] Coordinates:   %s%s, %s\n", White, NeonPink, data.Lat, data.Lon)
	fmt.Printf("%s[*] Tactical Map:  %s📍 %s\n", White, NeonBlue, data.MapLink)

	fmt.Printf("\n%s[ NAME_SERVER_CLUSTER ]%s\n", NeonYellow, Reset)
	for host, ips := range data.NameServers {
		fmt.Printf("%s[-] %-20s\n", White, host)
		for _, ip := range ips {
			fmt.Printf("    %s↳ [%s]%s\n", NeonBlue, ip, Reset)
		}
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ SATELLITE_TRIANGULATION_REPORT ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Target Number: %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Service:       %s%s (%s)\n", White, NeonGreen, p.Carrier, p.Type)
	fmt.Printf("%s[*] Country:       %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Region/State:  %s%s\n", White, NeonGreen, p.State)
	fmt.Printf("%s[*] Area/City:     %s%s\n", White, NeonGreen, p.Location)
	fmt.Printf("\n%s[ 📡 PINPOINT_DATA ]%s\n", NeonYellow, Reset)
	fmt.Printf("%s[*] Coordinates:   %s%s, %s\n", White, NeonPink, p.Lat, p.Lon)
	fmt.Printf("%s[*] Map Vector:    %s📍 %s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("───────────────────────────────────────\n")
}
