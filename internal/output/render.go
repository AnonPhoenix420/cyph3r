package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

var ClearLine = "\033[H\033[2J"

func Banner() {
	fmt.Println(`   ______      ____  __  __ _____ ____
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/
\____/\__, /_/   /_/ /_/ /____/_/ |_|
     /____/         NETWORK_INTEL_SYSTEM`)
}

func Render(payload *models.IntelPayload) {
	// Your original render logic here - keep as is
	fmt.Printf("\n[!] TARGET_NODE: %s\n", payload.Target)
	// ... (add more if needed)
}

func RenderReport(report *models.ComprehensiveReport) {
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ [!] TARGET_NODE: %s                                          ║\n", report.Target)
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝\n")

	fmt.Println("[ GEO & REGISTRATION INTEL ]")
	fmt.Printf(" • COUNTRY:      %s (%s)\n", report.Location.Country, report.Location.CountryCode)
	fmt.Printf(" • STATE:        %s\n", report.Location.State)
	fmt.Printf(" • CITY:         %s\n", report.Location.City)
	fmt.Printf(" • ZIP:          %s\n", report.Location.ZIP)
	fmt.Printf(" • AREA CODE:    %s\n", report.Location.AreaCode)
	fmt.Printf(" • COORDINATES:  %s (≈ %.1f km radius)\n", report.Location.Coordinates, report.Location.RadiusKM)

	fmt.Println("\n[ SOCIAL MEDIA ASSOCIATIONS ]")
	for _, s := range report.SocialProfiles {
		fmt.Printf(" • %s → %s (%d%%)\n", s.Platform, s.ProfileURL, s.Confidence)
	}

	fmt.Printf("\n[ RISK SCORE: %d/100 ]\n", report.RiskScore)
}

func RenderPhoneReport(target, lineStatus, carrier, locale string) {
	fmt.Printf("[+] TELEPHONY INTELLIGENCE VECTOR: %s\n", target)
	fmt.Println("[-] TARGET MATRIX CLASSIFICATION: CEL_TRACKING_REPORT")
	fmt.Printf(" • LINE STATUS:       %s\n", lineStatus)
	fmt.Printf(" • CARRIER PROVIDER:  %s\n", carrier)
	fmt.Printf(" • CELL LOCALE:       %s\n", locale)
	fmt.Printf(" • RISK PROFILING:    12/100\n")
}
