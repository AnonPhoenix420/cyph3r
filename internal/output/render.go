package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-41s %s║", Cyan, NeonPink+data.TargetName, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", Cyan, Reset)
	for _, ip := range data.TargetIPs {
		v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
		fmt.Printf(" %s↳ %s[%-2s]%s %s\n", Cyan, NeonBlue, v, NeonGreen, ip)
	}

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s35.6892° N, 51.3890° E %s(LOCKED)\n", Cyan, White, Cyan, Amber)

	fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", Cyan, Reset)
	for ns, ips := range data.NameServers {
		fmt.Printf(" %s[-] %s%s\n", NeonPink, White, ns)
		for _, ip := range ips {
			v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
			fmt.Printf("     %s↳ %s(%s)%s %s\n", Electric, NeonBlue, v, NeonGreen, ip)
		}
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf(" %s» %sCORE_OS:      %s%s\n", Amber, White, NeonYellow, strings.TrimPrefix(res, "STACK: "))
			continue
		}
		fmt.Printf(" %s»%s %s\n", NeonGreen, White, res)
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] PHONE_INTEL: %-42s %s║", Cyan, NeonPink+p.Number, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ ATTRIBUTE_DATA ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s CARRIER:  %s%s\n", Cyan, White, NeonYellow, p.Carrier)
	fmt.Printf(" %s•%s LOCATION: %s%s\n", Cyan, White, NeonGreen, p.Country)
	fmt.Printf(" %s•%s RISK:     %s%s\n", Cyan, White, Red, p.Risk)

	fmt.Printf("\n%s[ DIGITAL_FOOTPRINT ]%s\n", Cyan, Reset)
	fmt.Printf(" %s»%s ALIAS:    %s%s\n", Cyan, White, Amber, p.HandleHint)
	fmt.Printf(" %s»%s SOCIAL:   %s%s\n", Cyan, White, NeonGreen, strings.Join(p.SocialPresence, ", "))
	fmt.Printf("\n%s[*] %sMAP_VECTOR: %s%s%s\n", White, Cyan, NeonBlue, p.MapLink, Reset)
}
