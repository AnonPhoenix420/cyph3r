package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ TARGET_HUD ] ---\n%s[*] Node: %s%s\n", NeonPink, White, NeonBlue, data.TargetName)
	fmt.Printf("%s[*] Loc:  %s%s, %s\n", White, NeonGreen, data.City, data.Country)
	fmt.Printf("%s[*] Map:  %s%s\n", White, NeonBlue, data.MapLink)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ SATELLITE_HUD ] ---\n%s[*] Target: %s%s\n", NeonPink, White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Carrier: %s%s\n", White, NeonGreen, p.Carrier)
	fmt.Printf("%s[*] Vector:  %s%s\n", White, NeonBlue, p.MapLink)
}
