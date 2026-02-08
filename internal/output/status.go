package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// PrintWaveStatus displays the results of the port reconnaissance in Neon Green.
func PrintWaveStatus(port int, status string) {
	// Label in White, Port number in Neon Blue, Status in Neon Green
	fmt.Printf("%sPROBE %s[%-5d]%s | STATUS: %s%-12s%s\n", 
		White, NeonBlue, port, Reset, NeonGreen, status, Reset)
}

// PrintStatus handles general node identification messages.
func PrintStatus(data models.IntelData) {
	if len(data.IPs) > 0 {
		fmt.Printf("%s[*] Node Identified: %s%s%s\n", 
			NeonGreen, NeonBlue, data.IPs[0], Reset)
	} else {
		fmt.Printf("%s[!] Target resolution failed.%s\n", 
			NeonPink, Reset)
	}
}

// PrintScanHeader creates a clean transition into the probe phase.
func PrintScanHeader() {
	fmt.Printf("\n%s──[ MULTI-PROBE WAVE ANALYSIS ]──%s\n", White, Reset)
}
