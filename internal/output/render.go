package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs {
		fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.Org)

	if len(data.Subdomains) > 0 {
		fmt.Printf("\n%s[ IDENTIFIED_SUB_NODES ]\n", NeonPink)
		for _, s := range data.Subdomains {
			fmt.Printf("%s[+] Node: %s%s\n", NeonGreen, White, s)
		}
	}

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	for _, ns := range data.NameServers["NS"] {
		fmt.Printf("%s[-] %s\n", White, ns)
		for _, ip := range data.NameServers["IP_"+ns] {
			fmt.Printf("    %s↳ [%s]\n", NeonBlue, ip)
		}
	}

	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] Tactical Admin Scan: %s\n", White, data.TargetName)
		for _, p := range ports {
			if strings.Contains(p, "!") {
				fmt.Printf("%s[!] ALERT: PORT %s: %sCRITICAL_VERSION_FOUND\n", NeonYellow, p, White)
			} else {
				fmt.Printf("%s[+] PORT %s: %sOPEN\n", NeonGreen, p, White)
			}
		}
	}
	fmt.Printf("%s[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:   %s%s\n[*] Carrier:  %s%s\n", White, NeonBlue, p.Number, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Vector:   %s%s\n------------------------------------\n", White, NeonBlue, p.MapLink)
}
