package output

import "fmt"

// Status indicators and UI elements
func PulseNode(target string) {
	fmt.Printf("%s[⚡] PULSING NODE: %s%s%s\n", NeonYellow, White, target, Reset)
}

func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s\n", NeonPink, msg, Reset)
}

func PrintScanHeader() {
	fmt.Printf("\n%s[#] STARTING TACTICAL PORT PROBE...%s\n", NeonBlue, Reset)
}
