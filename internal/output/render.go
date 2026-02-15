package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Node:      %s%s\n", White, NeonBlue, data.TargetName)
	
	for i, ip := range data.TargetIPs {
		fmt.Printf("%s[*] IP [%d]:    %s%s\n", White, i, NeonGreen, ip)
	}

	fmt.Printf("%s[*] Provider:   %s%s (%s)\n", White, NeonYellow, data.ISP, data.Org)
	fmt.Printf("%s[*] Location:   %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	
	if len(data.NameServers["NS"]) > 0 {
		fmt.Printf("%s[*] NS Nodes:   %s%v\n", White, NeonBlue, data.NameServers["NS"])
	}
	if len(data.NameServers["MX"]) > 0 {
		fmt.Printf("%s[*] MX Nodes:   %s%v\n", White, NeonBlue, data.NameServers["MX"])
	}

	fmt.Printf("%s[*] Map Link:   %s%s\n", White, NeonBlue, data.MapLink)
	fmt.Printf("%s------------------------------------------%s\n", NeonPink, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:     %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Carrier:    %s%s\n", White, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Vector:     %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("%s------------------------------------%s\n", NeonPink, Reset)
}
