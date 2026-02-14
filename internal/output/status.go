package output

import "fmt"

func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s\n", Red, White, msg)
}

func Info(msg string) {
	fmt.Printf("%s[*] INFO: %s%s\n", NeonBlue, White, msg)
}

func Warning(msg string) {
	fmt.Printf("%s[!] WARNING: %s%s\n", NeonYellow, White, msg)
}
