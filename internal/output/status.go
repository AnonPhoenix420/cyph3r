package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PrintShieldStatus displays high-fidelity VPN/OPSEC session data
func PrintShieldStatus(shield intel.ShieldStatus) {
	fmt.Printf("%s[*] INFO: %sVerifying Shield Status... ", White, Cyan)
	
	if !shield.IsActive {
		fmt.Printf("%s[UNSECURED]%s\n", Red, Reset)
		return
	}
	
	if shield.VPNDetected {
		fmt.Printf("%s[SECURE - VPN ACTIVE]%s\n", NeonGreen, Reset)
	} else {
		fmt.Printf("%s[SECURE]%s\n", NeonGreen, Reset)
	}
	
	fmt.Printf("    %s↳ %sNode: %s%s %s| %sISP: %s%s %s| %sRisk: %d/100%s\n\n", 
		Cyan, White, NeonBlue, shield.Location, Gray, White, NeonYellow, shield.ISP, Gray, 
		func() string {
			if shield.VPNDetected {
				return NeonGreen
			}
			return Red
		}(), shield.Recommendation == "Active shield / VPN detected - Good OPSEC" ? 25 : 75, Reset)
}

// PrintScanStart is a helper for scan initialization
func PrintScanStart(target string) {
	fmt.Printf("%s[*] INFO: %sInitializing Satellite Uplink to %s%s%s...\n", 
		White, Cyan, NeonPink, target, Reset)
}

// RenderShieldReport renders full shield status using the new model
func RenderShieldReport() {
	report := intel.GetShieldReport()
	RenderReport(&report)
}
