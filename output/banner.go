package output

import "fmt"

// Banner prints a neon 3D-style ASCII logo spelling CYPH3R
func Banner() {
	logo := []string{
		`  _____  __     ____   ____  _   _ ____  _____ `,
		` / ____| \ \   / ___| |  _ \| | | |  _ \|  __ \`,
		`| |      \ \ / /     | |_) | | | | |_) | |__) |`,
		`| |       \ V /      |  _ <| | | |  _ <|  ___/ `,
		`| |____    | |       | |_) | |_| | |_) | |     `,
		` \_____|   |_|       |____/ \___/|____/|_|     `,
	}

	// Gradient / neon effect colors
	colors := []string{BoldCyan, BoldBlue, BoldMagenta, BoldCyan, BoldBlue, BoldMagenta}

	for i, line := range logo {
		fmt.Println(colors[i%len(colors)] + line + Reset)
	}

	fmt.Println(BoldYellow + "⚡ CYPH3R — High-Fidelity Go Network Tester" + Reset)
	fmt.Println(BoldGreen + "💡 Educational / Professional Use Only" + Reset)
	fmt.Println()
}
