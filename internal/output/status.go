package output

import "fmt"

// Error displays a critical failure message in red
func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s%s\n", Red, White, msg, Reset)
}

// Info displays a tactical information message in blue
func Info(msg string) {
	fmt.Printf("%s[*] INFO: %s%s%s\n", NeonBlue, White, msg, Reset)
}

// Warning displays a non-critical alert in yellow
func Warning(msg string) {
	fmt.Printf("%s[!] WARNING: %s%s%s\n", NeonYellow, White, msg, Reset)
}

// Success displays a positive result in green
func Success(msg string) {
	fmt.Printf("%s[+] SUCCESS: %s%s%s\n", NeonGreen, White, msg, Reset)
}
