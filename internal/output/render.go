package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ NODE_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Registrar: %s%s\n", White, NeonGreen, data.Registrar)
	fmt.Printf("%s[*] Provider:  %s%s\n", White, NeonGreen, data.ISP)
	fmt.Printf("%s[*] Location:  %s%s, %s\n", White, NeonGreen, data.City, data.Country)
	
	// Map Link
	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f", data.Lat, data.Lon)
	fmt.Printf("%s[*] Geo-Map:   %s%s%s\n", White, NeonBlue, mapURL, Reset)
}

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}
