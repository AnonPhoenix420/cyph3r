package output

import (
	"fmt"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func LoadingAnimation(done chan bool, label string) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print(ClearLine)
			return
		default:
			fmt.Printf("\r%s%s %sScanning %s%s...%s", ClearLine, NeonPink, frames[i%len(frames)], White, label, Reset)
			i++
			time.Sleep(80 * time.Millisecond)
		}
	}
}

func DisplayHUD(data models.IntelData, verbose bool) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-41s %s║", Cyan, NeonPink+data.TargetName, Electric)
	
	// NEW: Integrated WAF Shield into your box header
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-41s %s║", Amber, NeonYellow+data.WAFType, Electric)
	} else {
		fmt.Printf("\n║ %s[!] SHIELD:      %-41s %s║", Gray, "UNPROTECTED / DIRECT_IP", Electric)
	}
	
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// UPGRADED: NETWORK_VECTORS now shows PTR (Reverse DNS)
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", Cyan, Reset)
	for i, ip := range data.TargetIPs {
		v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
		
		// Align PTR data if it exists
		ptr := "---"
		if i < len(data.ReverseDNS) && data.ReverseDNS[i] != "NO_PTR" {
			ptr = data.ReverseDNS[i]
		}
		
		fmt.Printf(" %s↳ %s[%-2s]%s %-16s %s→ %-25s %s[LINK_ACTIVE]%s\n", 
			Cyan, NeonBlue, v, NeonGreen, ip, White, NeonPink+ptr, Electric, Reset)
	}

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY: %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)
	fmt.Printf(" %s•%s Location: %s%s, %s%s\n", Cyan, White, NeonGreen, data.City, data.Country, Reset)

	if len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", Cyan, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s%s\n", NeonPink, White, ns)
			for _, ip := range ips {
				fmt.Printf(" %s↳ %s%-35s %s[ONLINE]%s\n", Electric, NeonGreen, ip, NeonGreen, Reset)
			}
		}
	}

	if verbose && data.RawGeo != "" {
		fmt.Printf("\n%s[ RAW_GEO_DATA ]%s\n%s%s%s\n", Red, Reset, Amber, data.RawGeo, Reset)
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "USAGE:") {
			fmt.Printf("%s[*] INFRA_TYPE: %s%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "USAGE: "))
			continue
		}
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software:   %s%-15s %s[]%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "STACK: "), NeonBlue, Reset)
			continue
		}
		fmt.Printf("%s[+] %s%s\n", NeonGreen, White, res)
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] PHONE_INTEL: %-42s %s║", Cyan, NeonPink+p.Number, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)
	
	fmt.Printf("\n%s[ ATTRIBUTE_DATA ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s CARRIER:  %s%s\n", Cyan, White, NeonYellow, p.Carrier)
	fmt.Printf(" %s•%s LOCATION: %s%s\n", Cyan, White, NeonGreen, p.Country)
	fmt.Printf(" %s•%s RISK:     %s%s%s\n", Cyan, White, NeonGreen, p.Risk, Reset)
	
	fmt.Printf("\n%s[ DIGITAL_FOOTPRINT ]%s\n", Cyan, Reset)
	fmt.Printf(" %s»%s ALIAS: %s%s\n", Cyan, White, Amber, p.HandleHint)
	fmt.Printf(" %s»%s SOCIAL: %s%s\n", Cyan, White, NeonGreen, strings.Join(p.SocialPresence, ", "))
	
	fmt.Printf("\n%s[*] %sMAP_VECTOR: %s%s%s\n", White, Cyan, NeonBlue, p.MapLink, Reset)
}
