package output

import "fmt"

func Error(msg string) {
	fmt.Printf("%s[!] ERROR: %s%s%s\n", NeonPink, White, msg, Reset)
}
