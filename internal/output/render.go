package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func RenderReport(report *models.ComprehensiveReport) {
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ [!] TARGET_NODE: %s\n", report.Target)
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝\n")

	fmt.Println("[ REGISTRATION & GEO INTEL ]")
	fmt.Printf(" • COUNTRY:       %s (%s)\n", report.Location.Country, report.Location.CountryCode)
	fmt.Printf(" • STATE:         %s\n", report.Location.State)
	fmt.Printf(" • CITY:          %s\n", report.Location.City)
	fmt.Printf(" • ZIP / AREA:    %s / %s\n", report.Location.ZIP, report.Location.AreaCode)
	fmt.Printf(" • COORDINATES:   %s (≈ %.1f km radius)\n", report.Location.Coordinates, report.Location.RadiusKM)

	if report.ReverseDNS != "N/A" && report.ReverseDNS != "" {
		fmt.Printf(" • REVERSE DNS:   %s\n", report.ReverseDNS)
	}

	fmt.Println("\n[ ASSOCIATED CONTACTS ]")
	for _, contact := range report.Associated {
		fmt.Printf(" • %s\n", contact)
	}

	fmt.Println("\n[ SOCIAL MEDIA ASSOCIATIONS ]")
	for _, social := range report.SocialProfiles {
		fmt.Printf(" • %s: %s (%d%% confidence)\n", social.Platform, social.ProfileURL, social.Confidence)
	}

	fmt.Printf("\n[ RISK SCORE: %d/100 ]\n", report.RiskScore)
	fmt.Println("═══════════════════════════════════════════════════════════════\n")
}

// Legacy compatibility wrapper (for existing phone command)
func RenderPhoneReport(target string, lineStatus, carrier, locale string) {
	fmt.Printf("[+] TELEPHONY INTELLIGENCE VECTOR: %s\n", target)
	fmt.Println("[-] TARGET MATRIX CLASSIFICATION: CEL_TRACKING_REPORT")
	fmt.Printf(" • LINE STATUS:       %s\n", lineStatus)
	fmt.Printf(" • CARRIER PROVIDER:  %s\n", carrier)
	fmt.Printf(" • CELL LOCALE:       %s\n", locale)
}
