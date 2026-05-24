package output

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// drawBoxLine handles the layout formatting specifically for infrastructure mode
func drawBoxLine(label, value, labelCol, valCol string) {
	visibleText := fmt.Sprintf("[!] %s: %s", label, value)
	width := 59 
	padding := width - len(visibleText)
	if padding < 0 { padding = 0 }
	fmt.Printf("\n║ %s[!] %s: %s%s%s %s║", labelCol, label, valCol, value, strings.Repeat(" ", padding), NeonBlue)
}

// DisplayHUD automatically branches the UI based on the detected Vector Type
func DisplayHUD(data models.IntelData, verbose bool) {
	// 1. Clear any running terminal artifact lines
	fmt.Print(ClearLine)

	// 2. Call your signature banner from banner.go
	Banner()

	// 3. Render the correct unboxed style or infra box layout
	if data.VectorType == "EMAIL_STEALTH_VECTOR" {
		renderEmailHUD(data)
	} else {
		renderInfrastructureHUD(data)
	}
}

func renderEmailHUD(data models.IntelData) {
	// Your custom unboxed open-matrix stream layout
	fmt.Printf("\n%s[+] CYPH3R GHOST ELITE INTEL REPORT FOR: %s%s", NeonGreen, data.TargetName, Reset)
	fmt.Printf("\n%s[-] TARGET VECTOR MATRIX CLASSIFICATION: %s%s\n", NeonPink, data.VectorType, Reset)
	
	stealthStr := "FALSE"; if data.StealthStatus { stealthStr = "TRUE_STEALTH_VERIFIED" }
	dispStr := "FALSE"; if data.IsDisposable { dispStr = "TRUE" }

	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "STEALTH STATUS:", NeonGreen+Bold, stealthStr+Reset)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "USER VECTOR:", Cyan, data.UserVector)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "HOST ROUTE:", Amber, data.HostRoute)
	fmt.Printf("\n %s•%s %-15s %s%s", NeonPink, Reset, "DISPOSABLE:", Red, dispStr)
	fmt.Printf("\n %s•%s %-15s %s%s\n", NeonPink, Reset, "AVATAR TRACE:", Gray, data.AvatarTrace)
	
	fmt.Printf("\n%s[ RESOLVED MX STEALTH PATHS ]%s", NeonYellow, Reset)
	for _, mx := range data.MXPaths {
		fmt.Printf("\n  %s↳ %s%s", Electric, Reset, mx)
	}
	fmt.Println("\n")
}

func renderInfrastructureHUD(data models.IntelData) {
	// Legacy box mapping layout
	fmt.Printf("\n%s╔═══════════════════════════════════════════════════════════════╗", NeonBlue)
	drawBoxLine("TARGET_NODE", data.TargetName, Cyan, NeonPink)
	
	v4 := "NOT_DETECTED"; if len(data.TargetIPs) > 0 { v4 = data.TargetIPs[0] }
	drawBoxLine("TARGET_IPv4", v4, Amber, NeonGreen)

	v6 := "NOT_DETECTED"; if len(data.TargetIPv6s) > 0 { v6 = data.TargetIPv6s[0] }
	drawBoxLine("TARGET_IPv6", v6, Amber, Cyan)

	if data.IsWAF { drawBoxLine("SHIELD     ", data.WAFType, Amber, NeonYellow) }
	fmt.Printf("\n╚═══════════════════════════════════════════════════════════════╝%s\n", Reset)

	fmt.Printf("\n%s[ ORGANIZATION_DOX ]%s\n", NeonPink, Reset)
	if data.Org != "" { fmt.Printf(" • %-18s %s%s\n", "ENTITY_NAME:", NeonGreen, data.Org) }
	fmt.Printf(" • %-18s %s%s\n", "DESCRIPTION:", Gray, data.ISP)
	fmt.Printf(" • %-18s %s%s\n", "NETWORK_ASN:", NeonYellow, data.AS)

	fmt.Printf("\n%s[ GEO_ENTITY ]%s\n", NeonBlue, Reset)
	loc := fmt.Sprintf("%s, %s, %s", data.City, data.RegionName, data.Country)
	fmt.Printf(" • %-18s %s%s\n", "LOCATION:", NeonYellow, loc)

	fmt.Printf("\n%s[ INFRASTRUCTURE_STACK ]%s\n", NeonBlue, Reset)
	for _, res := range data.ScanResults {
		if strings.Contains(res, "PORT") {
			fmt.Printf(" [+] %-18s %s[ACTIVE]%s\n", NeonYellow+res, NeonGreen, Reset)
		}
	}
	fmt.Println()
}
