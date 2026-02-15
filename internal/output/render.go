package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	
	// IP Stack
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	// Restored the Zip and full location string
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.ISP)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	if len(data.NameServers["NS"]) > 0 {
		for _, ns := range data.NameServers["NS"] {
			fmt.Printf("%s[-] %s\n", White, ns)
			// The tactical sub-bullet from your original version
			fmt.Printf("    %s↳ %s[ACTIVE_NODE]\n", NeonBlue, White)
		}
	}

	// Restored the Tactical Scan Log Vibe
	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] INFO: Initializing Tactical Scan: %s%s\n", White, NeonPink, data.TargetName)
		for _, p := range ports {
			// Re-added the [ACK/SYN] tag for that raw network feel
			fmt.Printf("%s[+] PORT %s: %sOPEN [ACK/SYN]\n", NeonGreen, p, White)
		}
		fmt.Printf("%s[*] INFO: Tactical scan complete.\n", White)
	}

	fmt.Printf("%s[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}
