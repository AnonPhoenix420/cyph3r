package output

import "fmt"

func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s%s\n", NeonPink, White, msg, Reset)
}

func PulseNode(target string) {
	fmt.Printf("%s[*] ESTABLISHING PULSE WITH: %s%s%s\n", White, NeonPink, target, Reset)
}
