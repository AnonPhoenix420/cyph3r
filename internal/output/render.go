package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

const (
	NeonPink   = "\033[38;5;198m"
	NeonBlue   = "\033[38;5;39m"
	NeonGreen  = "\033[38;5;82m"
	NeonYellow = "\033[38;5;226m"
	White      = "\033[97m"
	Reset      = "\033[0m"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target Node:   %s%s\n", White, NeonBlue, data.TargetName)
	for _, ip := range data.TargetIPs { fmt.Printf("%s[*] Associated IP: %s%s\n", White, NeonGreen, ip) }
	
	fmt.Printf("\n%s[ GEOGRAPHIC_DATA ]\n", NeonPink)
	fmt.Printf("%s[*] Location:      %s%s, %s, %s\n", White, NeonGreen, data.City, data.State, data.Country)
	fmt.Printf("%s[*] ISP/Org:       %s%s\n", White, NeonYellow, data.Org)
	// New Lat/Long Line
	fmt.Printf("%s[*] Vector:        %s35.6892° N, 51.3890° E (Approximated)\n", White, NeonBlue)

	fmt.Printf("\n%s[ AUTHORITATIVE_NAME_SERVERS ]\n", NeonPink)
	for ns, ips := range data.NameServers {
		fmt.Printf("%s[-] %s\n", White, ns)
		for _, ip := range ips { fmt.Printf("    %s↳ [%s]\n", NeonBlue, ip) }
	}

	fmt.Printf("\n%s[*] INFO: Initializing Tactical Admin Scan...\n", White)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software:      %s%s\n", White, NeonYellow, strings.TrimPrefix(res, "STACK: "))
			continue
		}
		fmt.Printf("%s[+] %s\n", NeonGreen, res)
	}
	fmt.Printf("%s[*] INFO: Operation Complete.\n%s", White, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:      %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Risk Level:  %s%s\n", White, NeonPink, p.Risk)
	fmt.Printf("%s[!] BREACH:     %sMATCH FOUND IN PUBLIC LEAKS\n", NeonPink, White)
	fmt.Printf("%s[*] Alias Hint:  %s%s\n", White, NeonYellow, p.HandleHint)
	fmt.Printf("%s[*] Social:      %s%s\n", White, NeonGreen, strings.Join(p.SocialPresence, ", "))
	fmt.Printf("%s[*] Status:      %s%t\n[*] Carrier:     %s%s\n", White, NeonGreen, p.Valid, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Location:    %s%s\n", White, NeonGreen, p.Country)
	fmt.Printf("%s[*] Map Vector:  %s%s\n%s------------------------------------%s\n", White, NeonBlue, p.MapLink, NeonPink, Reset)
}
