package output

const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	White     = "\033[37m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)

func RedText(s string) string     { return Red + s + Reset }
func GreenText(s string) string   { return Green + s + Reset }
func YellowText(s string) string  { return Yellow + s + Reset }
func BlueText(s string) string    { return Blue + s + Reset }
func BoldText(s string) string    { return Bold + s + Reset }
func UnderlineText(s string) string { return Underline + s + Reset }
