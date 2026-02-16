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
	fmt.Printf("%s[*] Org/Owner:     %s%s\n", White, NeonYellow, data.Org)

	if len(data.Subdomains) > 0 {
		fmt.Printf("\n%s[ IDENTIFIED_SUB_NODES ]\n", NeonPink)
		for _, s := range data.Subdomains {
			fmt.Printf("%s[+] Node: %s%s\n", NeonGreen, White, s)
		}
	}

	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] Tactical Admin Scan: %s\n", White, data.TargetName)
		for _, p := range ports {
			fmt.Printf("%s[+] PORT %s: %sOPEN\n", NeonGreen, p, White)
		}
	}
	fmt.Printf("%s\n[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	riskColor := NeonGreen
	if p.BreachAlert || strings.Contains(p.Risk, "CRITICAL") { riskColor = NeonPink }

	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:      %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Risk Level:  %s%s\n", White, riskColor, p.Risk)
	
	if p.BreachAlert {
		fmt.Printf("%s[!] BREACH:     %sMATCH FOUND IN PUBLIC LEAKS\n", NeonPink, White)
		fmt.Printf("%s[*] Alias Hint:  %s%s\n", White, NeonYellow, p.HandleHint)
	}

	if len(p.AliasMatches) > 0 {
		fmt.Printf("%s[*] Footprint:   %s%s\n", White, NeonBlue, strings.Join(p.AliasMatches, ", "))
	}

	fmt.Printf("%s[*] Social:      %s%s\n", White, NeonGreen, strings.Join(p.SocialPresence, ", "))
	fmt.Printf("%s[*] Carrier:     %s%s\n", White, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Location:    %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Map Vector:  %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("%s------------------------------------%s\n", NeonPink, Reset)
}
