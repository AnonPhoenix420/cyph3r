package output

import "fmt"

func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s%s\n", NeonPink, White, msg, Reset)
}

func PrintStatus(label, msg string) {
	fmt.Printf("%s[*] %s: %s%s%s\n", White, label, NeonBlue, msg, Reset)
}
