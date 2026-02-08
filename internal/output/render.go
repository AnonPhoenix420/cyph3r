package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PrintNodeIntel displays the primary network identity of the target
func PrintNodeIntel(data models.IntelData, target string) {
	fmt.Printf("\n%s──[ NETWORK IDENTITY ]──%s\n", White, Reset)
	fmt.Printf("%s TARGET:    %s%s\n", White, NeonBlue, target, Reset)
	
	// IP Resolution in Neon Blue
	fmt.Printf("%s IP NODES:   ", White)
	if len(data.IPs) > 0 {
		for i, ip := range data.IPs {
			fmt.Printf("%s%s%s", NeonBlue, ip, Reset)
			if i < len(data.IPs)-1 {
				fmt.Printf(" %s|%s ", White, Reset)
			}
		}
	} else {
		fmt.Printf("%sNOT_RESOLVED%s", NeonPink, Reset)
	}
	fmt.Println()

	// Nameservers in Neon Yellow
	if len(data.Nameservers) > 0 {
		for _, ns := range data.Nameservers {
			fmt.Printf("%s NS NODE:    %s%s%s\n", White, NeonYellow, ns, Reset)
		}
	}
}

// PrintRegistryHUD displays WHOIS data and ownership info
func PrintRegistryHUD(data models.IntelData) {
	fmt.Printf("\n%s──[ REGISTRY INTEL ]──%s\n", White, Reset)
	
	registrar := data.Registrar
	if registrar == "" || registrar == "UNKNOWN" {
		registrar = "PROTECTED_OR_HIDDEN"
	}

	fmt.Printf("%s REGISTRAR:  %s%s%s\n", White, NeonYellow, registrar, Reset)
	
	// If you want a small snippet of the raw WHOIS, we show the first 3 lines of actual data
	if data.WhoisRaw != "" {
		fmt.Printf("%s STATUS:     %sACTIVE_RECORD_RETRIVED%s\n", White, NeonGreen, Reset)
	}
}

// PrintGeoHUD displays location metadata and the Neon Pink Map Link
func PrintGeoHUD(data models.IntelData) {
	fmt.Printf("\n%s──[ GEOGRAPHIC HUD ]──%s\n", White, Reset)
	
	fmt.Printf("%s LOCATION:   %s%s, %s%s\n", White, NeonYellow, data.City, data.Country, Reset)
	fmt.Printf("%s PROVIDER:   %s%s%s\n", White, NeonYellow, data.ISP, Reset)
	fmt.Printf("%s ASN:        %s%s%s\n", White, NeonYellow, data.ASN, Reset)
	fmt.Printf("%s COORDS:     %s%.4f, %.4f%s\n", White, NeonYellow, data.Lat, data.Lon, Reset)

	// Construct the dynamic Google Maps link (Neon Pink)
	// We use the latitude and longitude from our data model
	mapURL := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%.4f,%.4f", data.Lat, data.Lon)
	
	fmt.Printf("\n%s 🗺️  MAP_LINK:  %s%s%s\n", White, NeonPink, mapURL, Reset)
	fmt.Printf("%s────────────────────────────────────────────────%s\n", White, Reset)
}

// PrintScanHeader prepares the user for the port scan phase
func PrintScanHeader() {
	fmt.Printf("\n%s──[ EXECUTING PORT RECONNAISSANCE ]──%s\n", White, Reset)
}
