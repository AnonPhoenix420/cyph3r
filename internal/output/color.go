package output

import "fmt"

const (
	Blue          = "\033[34m"
	Red           = "\033[31m"
	Reset         = "\033[0m"
	MetallicGreen = "\033[1;38;5;82m" // bold + 256-color green for a "metallic" look
)

func BlueText(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}

func RedText(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

// Status helpers used by main.go
func Up(msg string) {
	fmt.Printf("%s[UP]%s %s\n", Blue, Reset, msg)
}

func Down(msg string) {
	fmt.Printf("%s[DOWN]%s %s\n", Red, Reset, msg)
}

// Tool banner
func Banner() {
	// ASCII-art header for "CYPH3R" in a FIGlet-like "standard" style, printed in metallic green
	fmt.Println(MetallicGreen + `
  _____    __     __   ______    _    _    _____   ______ 
 / ____|  \ \   / /  | ___ \  | |  | |  |___  |  | ___ \
| |       \ \_/ /   | |_/ /  | |__| |    / /   | |_/ / 
| |        \   /    |  __/   |  __  |  |_ \   |  _  \ 
| |____     | |     | |      | |  | |  ___) |  | | \ \
 \_____|    |_|     \_|      |_|  |_|  |____/   \_|  \_|
` + Reset)

	// Small title under the ASCII-art
	fmt.Println("        CYPH3R — Network Diagnostics Utility")
	fmt.Println("     ⚠ Educational & Professional Use Only ⚠")
	fmt.Println()
}
