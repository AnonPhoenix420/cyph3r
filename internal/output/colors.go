package output

import "fmt"

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Pink   = "\033[35m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
)

func RedText(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

func GreenText(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Reset)
}

func YellowText(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Reset)
}

func BlueText(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}

func PinkBoldText(s string) string {
	return fmt.Sprintf("%s%s%s%s", Pink, Bold, s, Reset)
}

func BoldText(s string) string {
	return fmt.Sprintf("%s%s%s", Bold, s, Reset)
}
