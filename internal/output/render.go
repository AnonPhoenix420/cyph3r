package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PulseNode initializes the scan visualizer
func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

// DisplayHUD renders the Network/Domain intelligence report
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
	}

	if ports := data.NameServers["PORTS"]; len(ports) > 0 {
		fmt.Printf("\n%s[*] Tactical Admin Scan: %s\n", White, data.TargetName)
		for _, p := range ports {
			fmt.Printf("%s[+] PORT %s: %sOPEN\n", NeonGreen, p, White)
		}
	}
	fmt.Printf("%s[+] SUCCESS: Operation Complete.\n%s", NeonGreen, Reset)
}

// DisplayPhoneHUD renders the Deep OSINT Satellite report with Alias Hunting
func DisplayPhoneHUD(p models.PhoneData) {
	// Risk Logic: NeonPink for Breaches or High Risk
	riskColor := NeonGreen
	if p.BreachAlert || strings.Contains(p.Risk, "HIGH") || strings.Contains(p.Risk, "CRITICAL") {
		riskColor = NeonPink
	}

	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:      %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Risk Level:  %s%s\n", White, riskColor, p.Risk)

	// Deep Social OSINT & Breach Logic
	if p.BreachAlert {
		fmt.Printf("%s[!] BREACH:     %sMATCH FOUND IN PUBLIC LEAKS\n", NeonPink, White)
		if p.HandleHint != "" {
			fmt.Printf("%s[*] Alias Hint:  %s%s\n", White, NeonYellow, p.HandleHint)
		}
	}

	// NEW: Display discovered Alias Footprint (GitHub, Reddit, etc.)
	if len(p.AliasMatches) > 0 {
		footprint := strings.Join(p.AliasMatches, ", ")
		fmt.Printf("%s[*] Footprint:   %s%s\n", White, NeonBlue, footprint)
	}

	if len(p.SocialPresence) > 0 {
		socials := strings.Join(p.SocialPresence, ", ")
		fmt.Printf("%s[*] Social:      %s%s\n", White, NeonGreen, socials)
	}

	fmt.Printf("%s[*] Status:      %s%t\n", White, NeonGreen, p.Valid)
	fmt.Printf("%s[*] Type:        %s%s\n", White, NeonYellow, p.Type)
	fmt.Printf("%s[*] Carrier:     %s%s\n", White, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Location:    %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Map Vector:  %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("%s------------------------------------%s\n", NeonPink, Reset)
}
