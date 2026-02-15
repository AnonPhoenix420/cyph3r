package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PulseNode handles the status line in main.go
func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

// DisplayHUD restores your original high-depth intelligence layout
func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	
	// IP Stack
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.ISP)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	if len(data.NameServers["NS"]) > 0 {
		for _, ns := range data.NameServers["NS"] {
			fmt.Printf("%s[-] %s\n", White, ns)
			fmt.Printf("    %s↳ %s[ACTIVE_NODE]\n", NeonBlue, White)
		}
	}

	// Tactical Scan Section
	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] INFO: Initializing Tactical Scan: %s%s\n", White, NeonPink, data.TargetName)
		for _, p := range ports {
			fmt.Printf("%s[+] PORT %s: %sOPEN [ACK/SYN]\n", NeonGreen, p, White)
		}
		fmt.Printf("%s[*] INFO: Tactical scan complete.\n", White)
	}

	fmt.Printf("%s[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}

// DisplayPhoneHUD restores the satellite tracking layout
func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:     %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Carrier:    %s%s\n", White, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Vector:     %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("%s------------------------------------%s\n", NeonPink, Reset)
}
