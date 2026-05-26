package output

import (
	"fmt"
	"strings"
	"time"
)

// DrawProgressBar paints a clean, high-visibility cyberpunk loading tracking metric
func DrawProgressBar(label string, current, total int) {
	if total <= 0 {
		return
	}

	width := 30
	percentage := float64(current) / float64(total)
	filledLength := int(percentage * float64(width))

	if filledLength > width {
		filledLength = width
	}
	if filledLength < 0 {
		filledLength = 0
	}

	// Create bar visualization components
	filled := strings.Repeat("█", filledLength)
	empty := strings.Repeat("░", width-filledLength)

	// Clean rendering string using native color constants from colors.go
	fmt.Printf("\r%s[%s]%s %-25s %s[%s%s]%s %3d%%",
		NeonPink, label, Reset,
		"",
		Cyan, filled, empty, Reset,
		int(percentage*100),
	)

	if current == total {
		fmt.Println()
	}
}

// DisplayStatusMessage renders a timed operational update to the telemetry screen
func DisplayStatusMessage(message string, isError bool) {
	timestamp := time.Now().Format("15:04:05")
	if isError {
		fmt.Printf("%s[%s] %s[-] ALERT: %s%s\n", Gray, timestamp, Red, message, Reset)
	} else {
		fmt.Printf("%s[%s] %s[+] INTEL: %s%s\n", Gray, timestamp, NeonGreen, message, Reset)
	}
}
