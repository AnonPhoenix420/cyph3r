package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PulseNode and DisplayPhoneHUD remain unchanged from previous stable version

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.ISP)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	for _, ns := range data.NameServers["NS"] {
		fmt.Printf("%s[-] %s\n", White, ns)
		ips := data.NameServers["IP_"+ns]
		for _, ip := range ips {
			fmt.Printf("    %s↳ [%s]\n", NeonBlue, ip)
		}
	}

	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] INFO: Initializing Tactical Admin Scan: %s%s\n", White, NeonPink, data.TargetName)
		for _, p := range ports {
			fmt.Printf("%s[+] PORT %s: %sOPEN [ACK/SYN]\n", NeonGreen, p, White)
		}
		fmt.Printf("%s[*] INFO: Admin/Web scan complete.\n", White)
	}
	fmt.Printf("%s[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}
