package output

import (
	"fmt"
	"time"
)

// Banner prints the high-fidelity CYPH3R ASCII signature
func Banner() {
	// We use a raw string literal (backticks) to preserve the slashes
	banner := `
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM`

	fmt.Println(CyanText(banner))
	fmt.Printf("  %s\n", WhiteText("v2.6 [STABLE] // Go 1.23 Edition"))
	fmt.Println(CyanText("  ───────────────────────────────────────"))
}

// ScanAnimation provides the visual sensor calibration effect
func ScanAnimation() {
	frames := []string{"◒", "◐", "◓", "◑"}
	fmt.Print(WhiteText("[*] Calibrating HUD Sensors... "))
	
	// Quick 1-second loop for the animation
	for i := 0; i < 10; i++ {
		fmt.Printf("\r[*] Calibrating HUD Sensors... %s ", CyanText(frames[i%len(frames)]))
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Printf("\r[*] Calibrating HUD Sensors... %s\n", GreenText("[READY]"))
}
