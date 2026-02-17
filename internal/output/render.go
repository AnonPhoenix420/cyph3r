package output

import (
	"fmt"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] %sIDENTIFYING NODE: %s%s%s\n", Cyan, White, NeonPink, target, Reset)
}

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

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-41s %s║", Cyan, NeonPink+data.TargetName, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", Cyan, Reset)
	for _, ip := range data.TargetIPs {
		v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
		fmt.Printf(" %s↳ %s[%-2s]%s %-35s %s[LINK_ACTIVE]%s\n", Cyan, NeonBlue, v, NeonGreen, ip, Electric, Reset)
	}

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s35.6892° N, 51.3890° E %s📡 %s(SIGNAL_LOCKED)\n", Cyan, White, Cyan, Amber, Amber)
	fmt.Printf(" %s•%s Location: %sTehran, Tehran, Iran%s\n", Cyan, White, NeonGreen, Reset)

	fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", Cyan, Reset)
	for ns, ips := range data.NameServers {
		fmt.Printf(" %s[-] %s%s\n", NeonPink, White, ns)
		for _, ip := range ips {
			v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
			fmt.Printf("     %s↳ %s(%s)%s %-35s %s[ONLINE]%s\n", Electric, NeonBlue, v, NeonGreen, ip, NeonGreen, Reset)
		}
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	fmt.Printf("%s[*] INFO: Initializing Tactical Admin Scan...%s\n", NeonBlue, Reset)
	
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			soft := strings.TrimPrefix(res, "STACK: ")
			fmt.Printf("%s[*] Software:      %s%s []%s\n", NeonBlue, NeonYellow, soft, Reset)
			continue
		}
		// Classic [+] Tactical Hits
		fmt.Printf("%s[+] %s%s\n", NeonGreen, White, res)
	}
	fmt.Printf("\n%s[*] %sSESSION_IDLE: Awaiting next vector.%s\n", Electric, Amber, Reset)
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
