package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Global Package Constants (Visible to status.go)
const (
	NeonPink   = "\033[38;5;198m"
	NeonBlue   = "\033[38;5;39m"
	NeonGreen  = "\033[38;5;82m"
	NeonYellow = "\033[38;5;226m"
	Red        = "\033[31m"
	White      = "\033[97m"
	Reset      = "\033[0m"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs { fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip) }
	
	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.Org)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	for ns, ips := range data.NameServers {
		fmt.Printf("%s[-] %s\n", White, ns)
		for _, ip := range ips { fmt.Printf("    %s↳ [%s]\n", NeonBlue, ip) }
	}
	fmt.Printf("\n%s[*] INFO: Initializing Tactical Admin Scan...\n", White)
	for _, res := range data.ScanResults { fmt.Printf("%s[+] %s\n", NeonGreen, res) }
	fmt.Printf("%s[*] INFO: Admin/Web scan complete.\n%s", White, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:      %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Risk Level:  %s%s\n", White, NeonPink, p.Risk)
	fmt.Printf("%s[!] BREACH:     %sMATCH FOUND IN PUBLIC LEAKS\n", NeonPink, White)
	fmt.Printf("%s[*] Alias Hint:  %s%s\n", White, NeonYellow, p.HandleHint)
	fmt.Printf("%s[*] Footprint:   %s%s\n", White, NeonBlue, strings.Join(p.AliasMatches, ", "))
	fmt.Printf("%s[*] Social:      %s%s\n", White, NeonGreen, strings.Join(p.SocialPresence, ", "))
	fmt.Printf("%s[*] Status:      %s%t\n[*] Type:        %s%s\n", White, NeonGreen, p.Valid, White, p.Type)
	fmt.Printf("%s[*] Carrier:     %s%s\n[*] Location:    %s%s\n", White, NeonYellow, p.Carrier, NeonGreen, p.Country)
	fmt.Printf("%s[*] Map Vector:  %s%s\n%s------------------------------------%s\n", White, NeonBlue, p.MapLink, NeonPink, Reset)
}
