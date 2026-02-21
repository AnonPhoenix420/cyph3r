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
			fmt.Print(ClearLine) // Uses definition from colors.go
			return
		default:
			fmt.Printf("\r%s%s %sScanning %s%s...%s", ClearLine, NeonPink, frames[i%len(frames)], White, label, Reset)
			i++
			time.Sleep(80 * time.Millisecond)
		}
	}
}

func DisplayHUD(data models.IntelData, verbose bool) {
	// --- HEADER TRENCH ---
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-42s %s║", Cyan, NeonPink+data.TargetName, Electric)
	
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Amber, NeonYellow+data.WAFType, Electric)
	} else {
		fmt.Printf("\n║ %s[!] SHIELD:      %-42s %s║", Gray, "UNPROTECTED / DIRECT_IP", Electric)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// --- NETWORK VECTORS ---
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", Cyan, Reset)
	for i, ip := range data.TargetIPs {
		ptr := "---"
		if i < len(data.ReverseDNS) && data.ReverseDNS[i] != "NO_PTR" {
			ptr = data.ReverseDNS[i]
		}
		fmt.Printf(" %s↳ %s[v]%s %-16s %s%s %-25s %s[LINK_ACTIVE]%s\n", Cyan, Electric, NeonGreen, ip, Gray, "→", NeonPink+ptr, Electric, Reset)
	}

	// --- GEO ENTITY ---
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)

	// --- INFRASTRUCTURE STACK ---
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "SUBDOMAIN:") {
			fmt.Printf("%s[»] %-45s %s[FOUND]%s\n", NeonPink, White+res, NeonGreen, Reset)
		} else if strings.Contains(res, "VULN_WARN") {
			fmt.Printf("%s[!] ALERT:      %s%s\n", Red, White+strings.TrimPrefix(res, "VULN_WARN: "), Reset)
		} else if strings.Contains(res, "DEBUG") {
			fmt.Printf("%s[*] %-40s %s[LEAK]%s\n", Amber, White+res, Red, Reset)
		} else if strings.Contains(res, "PORT") {
			fmt.Printf("%s[+] %-40s %s[ACTIVE]%s\n", NeonGreen, White+res, Electric, Reset)
		} else if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software:   %s%-25s %s[]%s\n", Electric, NeonYellow, strings.TrimPrefix(res, "STACK: "), Electric, Reset)
		} else {
			fmt.Printf("%s[*] %s%s\n", Electric, White, res)
		}
	}
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] PHONE_INTEL: %-42s %s║", Cyan, NeonPink+p.Number, Electric)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)
	fmt.Printf("\n%s[ ATTRIBUTE_DATA ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s CARRIER:  %s%s\n", Cyan, White, NeonYellow, p.Carrier)
	fmt.Printf(" %s•%s RISK:     %s%s%s\n", Cyan, White, NeonGreen, p.Risk, Reset)
}
