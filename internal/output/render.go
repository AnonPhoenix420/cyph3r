package output

import (
	"fmt"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// LoadingAnimation creates a pulsing "Searching" effect in a goroutine
func LoadingAnimation(done chan bool, label string) {
	chars := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("\r%s[+] %s: COMPLETE%s\n", NeonGreen, label, Reset)
			return
		default:
			fmt.Printf("\r%s[%s] %s...%s", NeonPink, chars[i%len(chars)], label, Reset)
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonPink)
	fmt.Printf("\n║ %s[!] TARGET IDENTIFIED: %-36s %s║", White, data.TargetName, NeonPink)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ NETWORK_NODES ]%s\n", NeonBlue, Reset)
	for _, ip := range data.TargetIPs {
		fmt.Printf(" %s↳%s IP_ADDR: %-15s %s[AUTHORIZED]%s\n", NeonBlue, White, ip, NeonGreen, Reset)
	}

	fmt.Printf("\n%s[ GEO_INTEL ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s•%s LOC: %s, %s, %s\n", NeonBlue, White, data.City, data.State, data.Country)
	fmt.Printf(" %s•%s ORG: %s%s%s\n", NeonBlue, White, NeonYellow, data.Org, Reset)
	fmt.Printf(" %s•%s MAP: %s35.6892° N, 51.3890° E (LINK_SENT)%s\n", NeonBlue, White, NeonBlue, Reset)

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf(" %s» %s%-20s%s\n", NeonYellow, White, "SOFTWARE:", strings.TrimPrefix(res, "STACK: "))
			continue
		}
		fmt.Printf(" %s»%s %s\n", NeonGreen, White, res)
	}
	fmt.Printf("\n%s[*] SYSTEM_IDLE: Awaiting next command.%s\n", NeonPink, Reset)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s[!] SATELLITE_UPLINK_ESTABLISHED%s\n", NeonPink, Reset)
	fmt.Printf("%s┌──────────────────────────────────────────────────────────────┐%s\n", White, Reset)
	fmt.Printf("│ %sTARGET: %-15s %s|%s RISK: %-21s %s│\n", White, p.Number, NeonPink, White, p.Risk, White)
	fmt.Printf("│ %sSTATUS: %-15t %s|%s CARRIER: %-18s %s│\n", White, p.Valid, NeonPink, White, p.Carrier, White)
	fmt.Printf("%s└──────────────────────────────────────────────────────────────┘%s\n", White, Reset)
	
	fmt.Printf("%s[ SOCIAL_FOOTPRINT ]%s\n", NeonBlue, Reset)
	fmt.Printf(" %s»%s ALIAS_HINT: %s%s%s\n", NeonBlue, White, NeonYellow, p.HandleHint, Reset)
	fmt.Printf(" %s»%s PLATFORMS:  %s%s%s\n", NeonBlue, White, NeonGreen, strings.Join(p.SocialPresence, ", "), Reset)
	fmt.Printf("\n%s[*] GPS_VECTOR: %s%s%s\n", White, NeonBlue, p.MapLink, Reset)
}
