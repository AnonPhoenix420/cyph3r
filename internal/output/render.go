package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Node:      %s%s\n", White, NeonBlue, data.TargetName)
	
	// IP Stack
	for i, ip := range data.TargetIPs {
		fmt.Printf("%s[*] IP [%d]:    %s%s\n", White, i, NeonGreen, ip)
	}

	// Location Stack
	fmt.Printf("%s[*] Provider:   %s%s\n", White, NeonYellow, data.ISP)
	fmt.Printf("%s[*] Org:        %s%s\n", White, NeonYellow, data.Org)
	fmt.Printf("%s[*] Location:   %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	
	// DNS Stack
	if len(data.NameServers["NS"]) > 0 {
		fmt.Printf("%s[*] NS Nodes:   %s%v\n", White, NeonBlue, data.NameServers["NS"])
	}
	if len(data.NameServers["MX"]) > 0 {
		fmt.Printf("%s[*] Mail Nodes: %s%v\n", White, NeonBlue, data.NameServers["MX"])
	}

	fmt.Printf("%s[*] Map Link:   %s%s\n", White, NeonBlue, data.MapLink)
	fmt.Printf("%s------------------------------------------%s\n", NeonPink, Reset)
}
