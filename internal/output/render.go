package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PrintNodeIntel handles the primary identity display (IPs and NS Nodes)
func PrintNodeIntel(data models.IntelData, target string) {
	fmt.Printf("\n%s──[ NODE IDENTITY ANALYSIS ]──%s\n", White, Reset)
	fmt.Printf("%s TARGET URL:  %s%s\n", White, NeonBlue, target, Reset)
	
	// IP Addresses - Neon Blue
	fmt.Printf("%s IP ADDRESS:  ", White)
	if len(data.IPs) > 0 {
		for i, ip := range data.IPs {
			fmt.Printf("%s%s%s", NeonBlue, ip, Reset)
			if i < len(data.IPs)-1 {
				fmt.Printf(" %s|%s ", White, Reset)
			}
		}
	} else {
		fmt.Printf("%sNULL_NODE%s", NeonPink, Reset)
	}
	fmt.Println()

	// Nameserver Nodes - Neon Yellow
	if len(data.Nameservers) > 0 {
		for _, ns := range data.Nameservers {
			fmt.Printf("%s NS NODE:     %s%s%s\n", White, NeonYellow, ns, Reset)
		}
	} else {
		fmt.Printf("%s NS NODE:     %sNO_RECORDS_FOUND%s\n", White, NeonPink, Reset)
	}
}

// PrintGeoHUD handles the geographic metadata and the Neon Pink weblink
func PrintGeoHUD(data models.IntelData) {
	fmt.Printf("\n%s──[ GEOGRAPHIC INTELLIGENCE ]──%s\n", White, Reset)
	
	// Location details in Neon Yellow
	fmt.Printf("%s 📍 LOCATION:   %s%s, %s%s\n", White, NeonYellow, data.City, data.Country, Reset)
	fmt.Printf("%s 🛰️  REGION:     %s%s (%s)%s\n", White, NeonYellow, data.Region, data.CountryCode, Reset)
	fmt.Printf("%s 🌐 COORDS:     %s%.4f, %.4f%s\n", White, NeonYellow, data.Lat, data.Lon, Reset)
	fmt.Printf("%s 📡 PROVIDER:   %s%s%s\n", White, NeonYellow, data.ISP, Reset)

	// The Neon Pink Map Link (Your specific request)
	// We use the latitude and longitude from the models to build a direct search link
	mapURL := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%.4f,%.4f", data.Lat, data.Lon)
	
	fmt.Printf("\n%s 🗺️  MAP_LINK:    %s%s%s\n", White, NeonPink, mapURL, Reset)
	fmt.Printf("%s────────────────────────────────────────────────%s\n", White, Reset)
}
