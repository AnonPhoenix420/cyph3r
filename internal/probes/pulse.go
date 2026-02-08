package output

import (
	"fmt"
	"time"
)

// ScanAnimation creates a pulsing "system calibration" effect in Neon Blue
func ScanAnimation() {
	frames := []string{"[■□□□□□□□□□]", "[■■□□□□□□□□]", "[■■■□□□□□□□]", "[■■■■□□□□□□]", "[■■■■■□□□□□]", "[■■■■■■□□□□]", "[■■■■■■■□□□]", "[■■■■■■■■□□]", "[■■■■■■■■■□]", "[■■■■■■■■■■]"}
	
	fmt.Printf("\n%s[*] INITIALIZING NEON TECH PULSE SEQUENCE...%s\n", White, Reset)
	
	for _, frame := range frames {
		fmt.Printf("\r%s%s CALIBRATING_SENSORS...%s", NeonBlue, frame, Reset)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n" + White + "[+] SYSTEM_READY: LINK_ESTABLISHED" + Reset)
}

// PulseNode creates a small delay to simulate a real-time "ping" to the target
func PulseNode(target string) {
	fmt.Printf("%s[*] PULSING TARGET NODE: %s%s%s", White, NeonBlue, target, Reset)
	for i := 0; i < 3; i++ {
		fmt.Printf("%s.%s", NeonBlue, Reset)
		time.Sleep(400 * time.Millisecond)
	}
	fmt.Println(" [ONLINE]")
}
