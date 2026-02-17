package output

import (
	"fmt"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ANSI Definitions
const (
	NeonPink = "\033[38;5;206m"
	NeonGreen = "\033[38;5;121m"
	NeonBlue = "\033[38;5;123m"
	NeonYellow = "\033[38;5;227m"
	Electric = "\033[38;5;45m"
	Cyan = "\033[0;36m"
	Amber = "\033[38;5;214m"
	White = "\033[0;37m"
	Red = "\033[0;31m"
	Reset = "\033[0m"
	ClearLine = "\033[2K\r"
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
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
	fmt.Printf(" %s•%s ENTITY:   %s%s\n", Cyan, White, NeonYellow, data.Org)
	fmt.Printf(" %s•%s POSITION: %s%.4f° N, %.4f° E %s📡 %s(SIGNAL: %s)\n", Cyan, White, Cyan, data.Lat, data.Lon, Amber, Amber, data.Latency)
	fmt.Printf(" %s•%s Location: %s%s, %s%s\n", Cyan, White, NeonGreen, data.City, data.Country, Reset)

	if verbose {
		fmt.Printf("\n%s[ RAW_GEO_DATA ]%s\n%s%s%s\n", Red, Reset, Amber, data.RawGeo, Reset)
	}

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", Cyan, Reset)
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "STACK:") {
			fmt.Printf("%s[*] Software:      %s%-15s %s[]%s\n", NeonBlue, NeonYellow, strings.TrimPrefix(res, "STACK: "), NeonBlue, Reset)
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
	fmt.Printf(" • CARRIER:  %s%s\n", NeonYellow, p.Carrier)
	fmt.Printf(" • LOCATION: %s%s\n", NeonGreen, p.Country)
	fmt.Printf(" • RISK:     %sLOW (Clearnet)%s\n", NeonGreen, Reset)

	fmt.Printf("\n%s[ DIGITAL_FOOTPRINT ]%s\n", Cyan, Reset)
	fmt.Printf(" » ALIAS:    %s%s\n", Amber, p.HandleHint)
	fmt.Printf(" » SOCIAL:   %s%s\n", NeonGreen, strings.Join(p.SocialPresence, ", "))
	fmt.Printf("\n%s[*] %sMAP_VECTOR: %s%s%s\n", White, Cyan, NeonBlue, p.MapLink, Reset)
}
