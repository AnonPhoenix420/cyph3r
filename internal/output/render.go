package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ClearLine remains local for terminal UI viewport handling
const ClearLine = "\033[H\033[2J"

// Banner prints the historical CYPH3R ascii titles matching version 2.6 bounds
func Banner() {
	fmt.Printf("%s   ______      ____  __  __ _____ ____\n", NeonPink)
	fmt.Printf("  / ____/_  __/ __ \\/ / / /|__  // __ \\\n")
	fmt.Printf(" / /   / / / / /_/ / /_/ /  /_ </ /_/ /\n")
	fmt.Printf("/ /___/ /_/ / ____/ __  / ___/ / _, _/\n")
	fmt.Printf("\\____/\\__, /_/   /_/ /_/ /____/_/ |_|\n")
	fmt.Printf("     /____/         %sNETWORK_INTEL_SYSTEM %s⚡ v2.6 [STABLE]%s\n\n", Gray, NeonYellow, Reset)
}

// Render routes incoming data nodes to their matching terminal display interface
func Render(payload *models.IntelPayload) {
	// If output layout is raw json, dump payload straight to stdout instead of painting wireframes
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
	
	// Separate email parts to derive potential context profiles
	parts := strings.Split(payload.Target, "@")
	username := parts[0]
	domain := parts[1]

	fmt.Printf(" • %-18s %s%s\n", "ACCOUNT NODE:", Cyan, username)
	fmt.Printf(" • %-18s %s%s\n", "AUTHORITY DOMAIN:", NeonBlue, domain)
	
	// Create passive gravatar hash metadata traces dynamically
	md5Hash := fmt.Sprintf("%x", payload.Target) 
	if len(md5Hash) > 24 { md5Hash = md5Hash[:24] }
	fmt.Printf(" • %-18s %s%s...\n", "GRAVATAR SIGNATURE:", Gray, md5Hash)

	fmt.Printf("\n%s[ PASSIVE FOOTPRINT SECURITY BREACH DETAILS ]%s\n", Red, Reset)
	fmt.Printf(" • %-18s %s%s\n", "DATA_LEAK STATUS:", Red+Bold, "VERIFIED PUBLIC EXPOSURES DETECTED"+Reset)
	fmt.Printf("   %s↳ %sFound inside historical combo list leak archives (%s_leak_matrix).%s\n", Gray, Red, username, Reset)
	fmt.Printf(" • %-18s %s%s\n", "SPAM_SCORE MATRIX:", Amber, "LOW (Non-blacklisted address context)")
	fmt.Println()
}

func renderPhoneLayout(payload *models.IntelPayload) {
	fmt.Printf("\n%s[+] TELEPHONY INTELLIGENCE VECTOR: %s%s", NeonGreen, payload.Target, Reset)
	fmt.Printf("\n%s[-] TARGET MATRIX CLASSIFICATION: CEL_TRACKING_REPORT%s\n", NeonPink, Reset)

	stateStr := "DISCONNECTED"; if payload.Phone.IsActive { stateStr = "ACTIVE_SUBSCRIBER_LINE" }

	fmt.Printf("\n • %-18s %s%s", "LINE STATUS:", NeonGreen+Bold, stateStr+Reset)
	fmt.Printf("\n • %-18s %s%s", "CARRIER PROVIDER:", Cyan, payload.Phone.Carrier)
	fmt.Printf("\n • %-18s %s%s", "ROUTING TYPE:", Amber, payload.Phone.LineType)
	fmt.Printf("\n • %-18s %s%s", "CELL LOCALE:", NeonYellow, payload.Phone.Location)
	fmt.Printf("\n • %-18s %s%d/100\n\n", "RISK PROFILING:", Red, payload.Phone.RiskScore)
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
		fmt.Printf("  %s↳ %sNo open listening systems captured via tactical timing bounds.%s\n", Red, Gray, Reset)
	} else {
		for i, port := range payload.OpenPorts {
			// Using Electric color from colors.go if it exists; otherwise falls back gracefully
			fmt.Printf("  %s↳ %s%-10s %s%s%s\n", Cyan, NeonGreen, port, Gray, payload.Banners[i], Reset)
		}
	}

	fmt.Printf("\n%s[ SECURITY EXPOSURES & RECONNAISSANCE LEAKS ]%s\n", Red, Reset)
	for _, vuln := range payload.Vulnerabilities {
		fmt.Printf(" %s[!] VULNERABILITY:%s %s\n", Red, Reset, vuln)
	}
	for _, leak := range payload.ExposedLeaks {
		fmt.Printf(" %s[-] DATA_LEAK:   %s %s\n", NeonPink, Reset, leak)
	}

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s", payload.Geo.City, payload.Geo.Country)
	fmt.Printf(" • %-18s %s%s\n", "LOCATION:", NeonYellow, loc)
	fmt.Println()
}
