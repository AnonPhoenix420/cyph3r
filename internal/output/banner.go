package output

import (
	"fmt"
	"runtime"
	"time"
)

// Banner prints the main Wireframe HUD logo
func Banner() {
	ascii := `
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM`

	fmt.Println(Cyan + Bold + ascii + Reset)
	fmt.Println(Gray + " ────────────────────────────────────────────────────────" + Reset)
	
	// System Metadata
	fmt.Printf(" %s[SYS]%s OS: %-7s | ARCH: %-7s | BUILD: 2.6\n", Bold, Reset, runtime.GOOS, runtime.GOARCH)
}

// ScanAnimation provides the high-speed "calibrating" HUD effect
func ScanAnimation() {
	// Braille or Cyber characters for the spinner
	chars := []string{"◐", "◓", "◑", "◒"}
	fmt.Print(Cyan + " [*] Calibrating HUD Sensors... ")
	
	for i := 0; i < 12; i++ {
		// \r returns the cursor to the start of the line to overwrite it
		fmt.Printf("\r %s[*] Calibrating HUD Sensors... %s", Cyan, chars[i%len(chars)])
		time.Sleep(80 * time.Millisecond)
	}
	
	// Print the final READY status in Green
	fmt.Println(Green + " [READY]" + Reset)
}
