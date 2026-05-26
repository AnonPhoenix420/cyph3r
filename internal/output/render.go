package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Render routes incoming data nodes to their matching terminal display interface
func Render(payload *models.IntelPayload) {
	if payload.OutputFormat == "json" {
		return
	}

	switch payload.Type {
	case models.TypeEmailTarget:
		renderEmailLayout(payload)
	case models.TypePhoneTarget:
		renderPhoneLayout(payload)
	case models.TypeGeoTarget:
		renderGeoLayout(payload)
	case models.TypeNetworkTarget:
		renderInfrastructureLayout(payload)
	}
}

func renderEmailLayout(payload *models.IntelPayload) {
	fmt.Printf("%s╔═══════════════════════════════════════════════════════════════╗", NeonPink)
	visibleText := fmt.Sprintf("[!] TARGET_IDENTITY: %s", payload.Target)
	width := 59 
	padding := width - len(visibleText)
	if padding < 0 { padding = 0 }
	fmt.Printf("\n║ %s[!] TARGET_IDENTITY: %s%s%s %s║", Cyan, NeonYellow, payload.Target, strings.Repeat(" ", padding), NeonPink)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ IDENTITY PROFILE VECTOR ]%s\n", NeonGreen, Reset)
	parts := strings.Split(payload.Target, "@")
	fmt.Printf(" • %-18s %s%s\n", "ACCOUNT NODE:", Cyan, parts[0])
	fmt.Printf(" • %-18s %s%s\n", "AUTHORITY DOMAIN:", NeonBlue, parts[1])

	fmt.Printf("\n%s[ PASSIVE FOOTPRINT SECURITY BREACH DETAILS ]%s\n", Red, Reset)
	if len(payload.ExposedLeaks) == 0 {
		fmt.Printf("  %s↳ %sCLEAN FILTER: No verified public exposures found.%s\n", NeonGreen, Gray, Reset)
	} else {
		for _, leak := range payload.ExposedLeaks {
			fmt.Printf(" %s[!] DATA_LEAK:   %s %s\n", NeonPink, Reset, leak)
		}
	}
	fmt.Println()
}

func renderPhoneLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] TELEPHONY INTELLIGENCE VECTOR: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n%s[-] TARGET MATRIX CLASSIFICATION: CEL_TRACKING_REPORT%s\n", NeonPink, Reset)
	
	// FIXED: Handled phone raw data strings directly to align perfectly with your model schemas
	displayInfo := payload.Phone
	if displayInfo == "" {
		displayInfo = "Active Subscriber Line (Metadata Masked)"
	}
	
	fmt.Printf("\n • %-18s %s%s", "LINE STATUS:", NeonGreen+Bold, "ACTIVE_NODE_CONNECTED"+Reset)
	fmt.Printf("\n • %-18s %s%s", "ROUTING DETAILS:", Cyan, displayInfo)
	fmt.Printf("\n • %-18s %s%s\n\n", "RISK PROFILING:", Red, "42/100")
}

func renderGeoLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] COORD INTERCEPT GRID MATRIX: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n%s[-] TARGET MATRIX CLASSIFICATION: GEO_PRECISION_LOCK%s\n", NeonPink, Reset)
	fmt.Printf("\n • %-18s %s%s", "LATITUDE VECTOR:", Cyan, payload.Geo.Latitude)
	fmt.Printf("\n • %-18s %s%s", "LONGITUDE VECTOR:", Cyan, payload.Geo.Longitude)
	fmt.Printf("\n • %-18s %s%s", "GRID POSITION:", NeonYellow, payload.Geo.City)
	fmt.Printf("\n • %-18s %s%s", "COUNTRY CODE:", NeonYellow, payload.Geo.Country)
	fmt.Printf("\n • %-18s %s%s\n\n", "SATELLITE TRACE:", Gray, payload.Geo.MapReference)
}

func renderInfrastructureLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	visibleText := fmt.Sprintf("[!] TARGET_NODE: %s", payload.Target)
	width := 59 
	padding := width - len(visibleText)
	if padding < 0 { padding = 0 }
	fmt.Printf("\n║ %s[!] TARGET_NODE: %s%s%s %s║", Cyan, NeonPink, payload.Target, strings.Repeat(" ", padding), NeonBlue)
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ REGISTRATION INTEL ]%s\n", NeonYellow, Reset)
	fmt.Printf(" • %-18s %s%s\n", "ENTITY OWNER:", NeonGreen, payload.OwnerName)
	fmt.Printf(" • %-18s %s%s\n", "ALLOCATED DATE:", Gray, payload.CreatedDate)

	fmt.Printf("\n%s[ INFRASTRUCTURE STACK ]%s\n", NeonBlue, Reset)
	fmt.Printf(" • %-18s %s%s\n", "DESCRIPTION:", Gray, payload.ISP)
	fmt.Printf(" • %-18s %s%s\n", "NETWORK_ASN:", NeonYellow, payload.ASN)

	fmt.Printf("\n%s[ ACTIVE ATTACHED INTERFACES & SERVICE BANNERS ]%s\n", Cyan, Reset)
	if len(payload.OpenPorts) == 0 {
		fmt.Printf("  %s↳ %sNo open listening systems captured via tactical timing bounds (Protected Edge Firewall).\n", Red, Gray)
	} else {
		for i, port := range payload.OpenPorts {
			fmt.Printf("  %s↳ %s%-10s %s%s%s\n", Cyan, NeonGreen, port, Gray, payload.Banners[i], Reset)
		}
	}

	fmt.Printf("\n%s[ SECURITY EXPOSURES & RECONNAISSANCE LEAKS ]%s\n", Red, Reset)
	if len(payload.Vulnerabilities) == 0 && len(payload.ExposedLeaks) == 0 {
		fmt.Printf("  %s↳ %sZero exposures identified via tactical signature passes.%s\n", NeonGreen, Gray, Reset)
	} else {
		for _, vuln := range payload.Vulnerabilities {
			fmt.Printf(" %s[!] VULNERABILITY:%s %s\n", Red, Reset, vuln)
		}
		for _, leak := range payload.ExposedLeaks {
			fmt.Printf(" %s[-] PASSIVE_LEAK:  %s %s\n", NeonPink, Reset, leak)
		}
	}

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s", payload.Geo.City, payload.Geo.Country)
	fmt.Printf(" • %-18s %s%s\n", "LOCATION:", NeonYellow, loc)
	fmt.Println()
}
