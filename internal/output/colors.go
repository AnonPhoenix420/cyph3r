package output

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Cyan   = "\033[38;5;51m" 
	Gray   = "\033[38;5;244m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

// Helper functions for easy colorizing
func BlueText(s string) string { return Cyan + Bold + s + Reset }
func RedText(s string) string  { return Red + s + Reset }
