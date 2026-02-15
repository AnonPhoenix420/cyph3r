package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Node:      %s%s\n", White, NeonBlue, data.TargetName)
	
	if len(data.TargetIPs) > 0 {
		fmt.Printf("%s[*] IP Vector:  %s%s\n", White, NeonGreen, data.TargetIPs[0])
	}

	fmt.Printf("%s[*] Provider:   %s%s (%s)\n", White, NeonYellow, data.ISP, data.Org)
	fmt.Printf("%s[*] Location:   %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	fmt.Printf("%s[*] Lat/Lon:    %s%s, %s\n", White, NeonPink, data.Lat, data.Lon)
	
	if len(data.NameServers["DNS"]) > 0 {
		fmt.Printf("%s[*] DNS Nodes:  %s%s\n", White, NeonBlue, data.NameServers["DNS"][0])
	}

	fmt.Printf("%s[*] Map Link:   %s%s\n", White, NeonBlue, data.MapLink)
	fmt.Printf("%s------------------------------------------%s\n", NeonPink, Reset)
}
