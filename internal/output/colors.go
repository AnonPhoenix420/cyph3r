package output

// Neon Tech Color Palette
// Using 256-color ANSI escape sequences for high-vibrancy output
const (
	Reset      = "\033[0m"
	White      = "\033[37m"
	
	// Neon Blue (Banner & IP Addresses)
	NeonBlue   = "\033[38;5;45m"
	
	// Neon Pink (Web Links / Fatal Errors)
	NeonPink   = "\033[38;5;201m"
	
	// Neon Green (Port Status / Success)
	NeonGreen  = "\033[38;5;82m"
	
	// Neon Yellow (GeoIP / NS Nodes)
	NeonYellow = "\033[38;5;226m"
)
