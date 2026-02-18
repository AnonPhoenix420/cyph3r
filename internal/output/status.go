package output

import (
	"fmt"
)

// PrintShieldStatus displays the high-fidelity VPN session data
func PrintShieldStatus(active bool, location string, isp string) {
	fmt.Printf("%s[*] INFO: %sVerifying Shield Status... ", White, Cyan)
	
	if !active {
		fmt.Printf("%s[UNSECURED]%s\n", Red, Reset)
		return
	}
	
	fmt.Printf("%s[SECURE]%s\n", NeonGreen, Reset)
	
	// Sub-line with session details using Gray for the separators
	fmt.Printf("    %s↳ %sNode: %s%s %s| %sISP: %s%s%s\n\n", 
		Cyan, White, NeonBlue, location, Gray, White, NeonYellow, isp, Reset)
}

// PrintScanStart is a small helper for the initial uplink message
func PrintScanStart(target string) {
	fmt.Printf("%s[*] INFO: %sInitializing Satellite Uplink to %s%s%s...\n", 
		White, Cyan, NeonPink, target, Reset)
}
