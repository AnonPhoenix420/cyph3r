package output

import (
	"fmt"
	"time"
)

func Banner() {
	banner := `
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM`
	
	fmt.Println(CyanText(banner))
	fmt.Printf("  %s\n", WhiteText("v2.6 [STABLE] // Wireframe HUD Edition"))
	fmt.Println(CyanText("  ───────────────────────────────────────"))
}

func ScanAnimation() {
	frames := []string{"◒", "◐", "◓", "◑"}
	fmt.Print(WhiteText("[*] Calibrating HUD Sensors... "))
	for i := 0; i < 10; i++ {
		fmt.Printf("\r[*] Calibrating HUD Sensors... %s ", CyanText(frames[i%len(frames)]))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r[*] Calibrating HUD Sensors... %s\n", GreenText("[READY]"))
}
