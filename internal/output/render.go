package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Render handles legacy IntelPayload rendering (backward compatibility)
func Render(payload *models.IntelPayload) {
	fmt.Print(ClearLine) // Use global from colors.go or banner.go
	Banner()

	fmt.Printf("%s[!] TARGET_NODE: %s%s\n\n", NeonBlue, payload.Target, Reset)

	if payload.Geo.City != "" || payload.Geo.Country != "" {
		fmt.Printf("%s[ GEO_ENTITY ]%s\n", Cyan, Reset)
		fmt.Printf(" • LOCATION:     %s, %s\n", payload.Geo.City, payload.Geo.Country)
	}

	if payload.OwnerName != "" {
		fmt.Printf(" • ENTITY OWNER: %s\n", payload.OwnerName)
	}

	if payload.ASN != "" {
		fmt.Printf(" • NETWORK_ASN:  %s\n", payload.ASN)
	}

	if len(payload.OpenPorts) > 0 {
		fmt.Printf("%s[ ACTIVE PORTS ]%s\n", Cyan, Reset)
		fmt.Printf(" • OPEN PORTS:   %v\n", payload.OpenPorts)
	}

	if len(payload.Vulnerabilities) > 0 {
		fmt.Printf("%s[ SECURITY EXPOSURES ]%s\n", Red, Reset)
		for _, vuln := range payload.Vulnerabilities {
			fmt.Printf(" • %s\n", vuln)
		}
	}

	fmt.Printf("\n%s[ RISK PROFILING: %d/100 ]%s\n", NeonYellow, 42, Reset)
	fmt.Println("═══════════════════════════════════════════════════════════════\n")
}

// RenderReport handles the new elite ComprehensiveReport (used with --full / -v)
func RenderReport(report *models.ComprehensiveReport) {
	fmt.Print(ClearLine)
	Banner()

	fmt.Printf("%s╔═══════════════════════════════════════════════════════════════╗%s\n", NeonBlue, Reset)
	fmt.Printf("%s║ [!] TARGET_NODE: %-45s ║%s\n", NeonBlue, report.Target, Reset)
	fmt.Printf("%s╚═══════════════════════════════════════════════════════════════╝%s\n\n", NeonBlue, Reset)

	// Geo & Location Intelligence
	fmt.Printf("%s[ GEO & REGISTRATION INTEL ]%s\n", Cyan, Reset)
	fmt.Printf(" • COUNTRY:      %s (%s)\n", report.Location.Country, report.Location.CountryCode)
	fmt.Printf(" • STATE:        %s\n", report.Location.State)
	fmt.Printf(" • CITY:         %s\n", report.Location.City)
	fmt.Printf(" • ZIP:          %s\n", report.Location.ZIP)
	fmt.Printf(" • AREA CODE:    %s\n", report.Location.AreaCode)
	fmt.Printf(" • COORDINATES:  %s (≈ %.1f km radius)\n", report.Location.Coordinates, report.Location.RadiusKM)

	if report.ReverseDNS != "" && report.ReverseDNS != "N/A" {
		fmt.Printf(" • REVERSE DNS:  %s\n", report.ReverseDNS)
	}

	// Associated Contacts & DNS
	fmt.Printf("\n%s[ ASSOCIATED CONTACTS & DNS ]%s\n", Cyan, Reset)
	for _, contact := range report.Associated {
		fmt.Printf(" • %s\n", contact)
	}

	// Social Media Intelligence
	if len(report.SocialProfiles) > 0 {
		fmt.Printf("\n%s[ SOCIAL MEDIA ASSOCIATIONS ]%s\n", Cyan, Reset)
		for _, social := range report.SocialProfiles {
			fmt.Printf(" • %s: %s %s(%d%% confidence)%s\n", 
				social.Platform, social.ProfileURL, Gray, social.Confidence, Reset)
		}
	}

	// Risk & Summary
	fmt.Printf("\n%s[ RISK SCORE: %d/100 ]%s\n", NeonYellow, report.RiskScore, Reset)
	fmt.Println("═══════════════════════════════════════════════════════════════\n")
}

// RenderPhoneReport maintains legacy phone output
func RenderPhoneReport(target, lineStatus, carrier, locale string) {
	fmt.Print(ClearLine)
	Banner()

	fmt.Printf("%s[+] TELEPHONY INTELLIGENCE VECTOR: %s%s\n", NeonBlue, target, Reset)
	fmt.Printf("%s[-] TARGET MATRIX CLASSIFICATION: CEL_TRACKING_REPORT%s\n\n", Cyan, Reset)

	fmt.Printf(" • LINE STATUS:       %s\n", lineStatus)
	fmt.Printf(" • CARRIER PROVIDER:  %s\n", carrier)
	fmt.Printf(" • CELL LOCALE:       %s\n", locale)
	fmt.Printf(" • RISK PROFILING:    12/100\n\n")
	fmt.Println("═══════════════════════════════════════════════════════════════\n")
}
