package output

import "fmt"

// ================= COLOR CODES =================
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	Bold        = "\033[1m"
	BoldRed     = "\033[1;31m"
	BoldGreen   = "\033[1;32m"
	BoldYellow  = "\033[1;33m"
	BoldBlue    = "\033[1;34m"
	BoldMagenta = "\033[1;35m"
	BoldCyan    = "\033[1;36m"
)

// ================= STATUS HELPERS =================

func Up(msg string) {
	fmt.Println(BoldBlue + "[UP] " + Reset + msg)
}

func Down(msg string) {
	fmt.Println(BoldRed + "[DOWN] " + Reset + msg)
}

func Info(msg string) {
	fmt.Println(BoldYellow + "[INFO] " + Reset + msg)
}

func Success(msg string) {
	fmt.Println(BoldGreen + "[SUCCESS] " + Reset + msg)
}

// Optional colored text helpers
func RedText(s string) string { return Red + s + Reset }
func GreenText(s string) string { return Green + s + Reset }
func BlueText(s string) string { return Blue + s + Reset }
func YellowText(s string) string { return Yellow + s + Reset }
func MagentaText(s string) string { return Magenta + s + Reset }
func CyanText(s string) string { return Cyan + s + Reset }

// Bold helpers
func BoldText(s string) string { return Bold + s + Reset }
func BoldRedText(s string) string { return BoldRed + s + Reset }
func BoldGreenText(s string) string { return BoldGreen + s + Reset }
func BoldBlueText(s string) string { return BoldBlue + s + Reset }
func BoldYellowText(s string) string { return BoldYellow + s + Reset }
func BoldMagentaText(s string) string { return BoldMagenta + s + Reset }
func BoldCyanText(s string) string { return BoldCyan + s + Reset }
