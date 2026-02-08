package output

import (
	"fmt"
	"time"
)

// ScanAnimation creates a high-tech loading sequence
func ScanAnimation() {
	frames := []string{"[■□□□□□□□□□]", "[■■□□□□□□□□]", "[■■■□□□□□□□]", "[■■■■□□□□□□]", "[■■■■■□□□□□]", "[■■■■■■□□□□]", "[■■■■■■■□□□]", "[■■■■■■■■□□]", "[■■■■■■■■■□]", "[■■■■■■■■■■]"}
	
	fmt.Printf("%s[*] INITIALIZING NEON_LINK...%s\n", White, Reset)
	
	for _, frame := range frames {
		fmt.Printf("\r%s %s LOADING_CORES %s", NeonBlue, frame, Reset)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// PulseNode simulates a connection ping to the target
func PulseNode(target string) {
	fmt.Printf("%s[*] ESTABLISHING PULSE WITH: %s%s%s\n", White, NeonPink, target, Reset)
	
	for i := 0; i < 3; i++ {
		fmt.Printf("%s . %s", NeonGreen, Reset)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("%s [ONLINE]%s\n", NeonGreen, Reset)
	time.Sleep(500 * time.Millisecond)
}

// StatusUpdate provides a small tactical update line
func StatusUpdate(msg string) {
	fmt.Printf("%s[>] %s...%s\n", White, msg, Reset)
}
