package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PrintScanHeader() {
	fmt.Printf("\n%s[#] STARTING TACTICAL PORT PROBE...%s\n", NeonBlue, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ NODE_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Registrar: %s%s\n", White, NeonGreen, data.Registrar, Reset)
	fmt.Printf("%s[*] Location:  %s%s, %s%s\n", White, NeonGreen, data.City, data.Country, Reset)
	fmt.Printf("%s[*] Provider:  %s%s (%s)%s\n", White, NeonGreen, data.ISP, data.ASN, Reset)
	fmt.Printf("%s[*] Geo:       %s%f, %f%s\n", White, NeonBlue, data.Lat, data.Lon, Reset)
}
