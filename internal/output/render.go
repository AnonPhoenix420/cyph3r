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
	// --- HEADER: SHIELD & TARGET ---
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", Electric)
	fmt.Printf("\n║ %s[!] TARGET_NODE: %-41s %s║", Cyan, NeonPink+data.TargetName, Electric)
	if data.IsWAF {
		fmt.Printf("\n║ %s[!] SHIELD:      %-41s %s║", Amber, NeonYellow+data.WAFType, Electric)
	} else {
		fmt.Printf("\n║ %s[!] SHIELD:      %-41s %s║", Gray, "UNPROTECTED / DIRECT_IP", Electric)
	}
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	// --- NETWORK VECTORS (With PTR Support) ---
	fmt.Printf("\n%s[ NETWORK_VECTORS ]%s\n", Cyan, Reset)
	for i, ip := range data.TargetIPs {
		v := "v4"; if strings.Contains(ip, ":") { v = "v6" }
		ptr := "---"
		if i < len(data.ReverseDNS) && data.ReverseDNS[i] != "NO_PTR" { ptr = data.ReverseDNS[i] }
		fmt.Printf(" %s↳ %s[%-2s]%s %-16s %s→ %-25s %s[LINK_ACTIVE]%s\n", 
			Cyan, NeonBlue, v, NeonGreen, ip, Gray, "→", NeonPink+ptr, Electric, Reset)
	}

	// --- GEO ENTITY (Added ISP/ASN/Time for "IP-Tracer" parity) ---
	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s ISP/ASN: %s%s\n", Cyan, White, NeonBlue, data.Latency) // Repurposing Latency slot for ISP if Org is empty
	fmt.Printf(" %s•%s POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)
	fmt.Printf(" %s•%s Location: %s%s, %s%s\n", Cyan, White, NeonGreen, data.City, data.Country, Reset)

	// --- RESTORED: AUTHORITATIVE_CLUSTERS ---
	if len(data.NameServers) > 0 {
		fmt.Printf("\n%s[ AUTHORITATIVE_CLUSTERS ]%s\n", Cyan, Reset)
		for ns, ips := range data.NameServers {
			fmt.Printf(" %s[-] %s%s\n", NeonPink, White, ns)
			for _, ip := range ips {
				fmt.Printf(" %s↳ %s%-35s %s[ONLINE]%s\n", Electric, NeonGreen, ip, NeonGreen, Reset)
			}
		}
	}

	// --- RAW DATA ---
	if verbose && data.RawGeo != "" {
		fmt.Printf("\n%s[ RAW_GEO_DATA ]%s\n%s%s%s\n", Red, Reset, Amber, data.RawGeo, Reset)
	}

	// --- RESTORED: INFRASTRUCTURE STACK (The "Admin Ports") ---
	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "USAGE:") {
			fmt.Printf("%s[*] INFRA_TYPE: %s%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "USAGE: "))
			continue
		}
		// RESTORED: Full Port/Software Logic
		if strings.Contains(res, "PORT") {
			fmt.Printf("%s[+] %s%-30s %s[ACTIVE]%s\n", NeonGreen, White, res, Electric, Reset)
		} else if strings.Contains(res, "Software") || strings.Contains(res, "ArvanCloud") {
			fmt.Printf("%s[*] Software:   %s%-15s %s[]%s\n", NeonBlue, NeonYellow, res, NeonBlue, Reset)
		} else {
			fmt.Printf("%s[+] %s%s\n", NeonGreen, White, res)
		}
	}
}
