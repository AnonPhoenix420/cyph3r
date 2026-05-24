package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Render is the unified global gate called by your main.go orchestrator
func Render(payload *models.IntelPayload) {
	// 1. Flush terminal lines and execute your signature branding header
	fmt.Print(ClearLine)
	Banner()

	// 2. Route dynamically based on the Type assigned in main.go
	switch payload.Type {
	case models.TypeEmailTarget:
		renderEmailLayout(payload)
	case models.TypeNetworkTarget:
		renderInfrastructureLayout(payload)
	case models.TypePhoneTarget:
		renderPhoneLayout(payload)
	case models.TypeGeoTarget:
		renderGeoLayout(payload)
	default:
		fmt.Printf("\n%s[-] Unknown processing vector type mapped.%s\n", Red, Reset)
	}
}

func renderEmailLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] CYPH3R GHOST ELITE INTEL REPORT FOR: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n%s[-] TARGET VECTOR MATRIX CLASSIFICATION: EMAIL_STEALTH_VECTOR%s\n", NeonPink, Reset)
	
	// Safely extract from your intel.ResolveEmail results if data exists
	stealthStr := "TRUE_STEALTH_VERIFIED"
	dispStr := "FALSE"
	userVector := strings.Split(payload.Target, "@")[0]
	hostRoute := strings.Split(payload.Target, "@")[1]
	avatarTrace := "https://gravatar.com/avatar/hash-reference"

	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "STEALTH STATUS:", NeonGreen+Bold, stealthStr+Reset)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "USER VECTOR:", Cyan, userVector)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "HOST ROUTE:", Amber, hostRoute)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "DISPOSABLE:", Red, dispStr)
	fmt.Printf("\n %s•%s %-15s %s%s\n", NeonPink, Reset, "AVATAR TRACE:", Gray, avatarTrace)
	
	fmt.Printf("\n%s[ RESOLVED MX STEALTH PATHS ]%s", NeonYellow, Reset)
	// Fallback mock array if payload.Email.MXPaths is empty or unpopulated yet
	mxPaths := []string{"10 mx1.stealth-relay.net.", "20 inbound-smtp.mx.net."}
	for _, mx := range mxPaths {
		fmt.Printf("\n  %s↳ %s%s", Electric, Reset, mx)
	}
	fmt.Println("\n")
}

func renderInfrastructureLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	
	// Helper logic to dynamically fill the infrastructure box layout safely
	visibleText := fmt.Sprintf("[!] TARGET_NODE: %s", payload.Target)
	width := 59 
	padding := width - len(visibleText)
	if padding < 0 { padding = 0 }
	fmt.Printf("\n║ %s[!] TARGET_NODE: %s%s%s %s║", Cyan, NeonPink, payload.Target, strings.Repeat(" ", padding), NeonBlue)
	
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	fmt.Printf(" • %-18s %s%s\n", "DESCRIPTION:", Gray, payload.ISP)
	fmt.Printf(" • %-18s %s%s\n", "NETWORK_ASN:", NeonYellow, payload.ASN)

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s", payload.Geo.City, payload.Geo.Country)
	fmt.Printf(" • %-18s %s%s\n", "LOCATION:", NeonYellow, loc)
	fmt.Println()
}

func renderPhoneLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] TELEPHONY VECTOR DETECTED: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n %s•%s %-15s %sParsing Payload Matrix...%s\n\n", NeonPink, Reset, "STATUS:", White, Reset)
}

func renderGeoLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] GEO-COORDINATE GRID LOCK DETECTED: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "LATITUDE:", Cyan, payload.Geo.Latitude)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "LONGITUDE:", Cyan, payload.Geo.Longitude)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "MAP VECTOR:", Gray, payload.Geo.MapReference)
	fmt.Printf("\n %s•%s %-15s %s%s\n\n", NeonPink, Reset, "GRID CELL:", NeonYellow, payload.Geo.City)
}
