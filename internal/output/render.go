package output

import (
	"fmt"
	"cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ REMOTE_TARGET_INTELLIGENCE_HUD ] ---\n%s[*] Node: %s%s\n", NeonPink, White, NeonBlue, data.TargetName)
	fmt.Printf("%s[*] Location: %s%s, %s, %s %s\n", White, NeonGreen, data.City, data.State, data.Country, data.Zip)
	fmt.Printf("%s[*] Coordinates: %s%s, %s\n", White, NeonPink, data.Lat, data.Lon)
	fmt.Printf("%s[*] Map Vector: %s📍 %s\n", White, NeonBlue, data.MapLink)
}

func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_TRIANGULATION ] ---\n%s[*] Number: %s%s\n", NeonPink, White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Region: %s%s (%s)\n", White, NeonGreen, p.State, p.Location)
	fmt.Printf("%s[*] Lat/Lon: %s%s, %s\n", White, NeonPink, p.Lat, p.Lon)
	fmt.Printf("%s[*] Map Vector: %s📍 %s\n", White, NeonBlue, p.MapLink)
}
