package output

import (
	"fmt"
	"os"
)

// Success logs a positive result with a Neon Green indicator
func Success(message string) {
	fmt.Printf("%s[+] SUCCESS: %s%s%s\n", NeonGreen, White, message, Reset)
}

// Warning logs a non-fatal issue with a Neon Yellow indicator
func Warning(message string) {
	fmt.Printf("%s[!] WARNING: %s%s%s\n", NeonYellow, White, message, Reset)
}

// Error logs a fatal issue with a Neon Pink indicator
func Error(message string) {
	fmt.Printf("%s[!] FATAL: %s%s%s\n", NeonPink, White, message, Reset)
}

// Critical exits the program after displaying a high-alert message
func Critical(message string) {
	fmt.Printf("\n%s[☠] CRITICAL_FAILURE: %s%s%s\n", NeonPink, White, message, Reset)
	os.Exit(1)
}

// Progress prints a subtle status line without a newline for "live" updates
func Progress(message string) {
	fmt.Printf("\r%s[*] %s...%s", NeonBlue, message, Reset)
}
