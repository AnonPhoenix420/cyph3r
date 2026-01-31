package output

import (
	"fmt"
	"runtime"
)

// Banner prints the updated CYPH3R branding to the console
func Banner() {
	// Standard CYPH3R Blue
	color := Blue 

	ascii := `
  _____   __     __   _____   _    _    _____   _____
 / ____|  \ \   / /  | ___ \ | |  | |  |___ |  | ___ \
| |       \ \_/ /   | |_/ /  | |__| |    / /   | |_/ /
| |        \   /    |  __/   |  __  |  |_ \    |  _  \
| |____     | |     | |      | |  | |  ___) |  |  | \ \
 \_____|    |_|     \_|      |_|  |_|  |____/   \_|  \_|
`

	fmt.Print(color + Bold + ascii + Reset)
	fmt.Println(BlueText(" │          CYPH3R: Network Diagnostics & Intel Tool"))
	fmt.Println(BlueText(" └──────────────────────────────────────────────────┘"))
	
	// Dynamic Metadata
	fmt.Printf(" [OS: %s] | [Arch: %s] | [Status: Active]\n", runtime.GOOS, runtime.GOARCH)
	
	// The Hacker Mascot
	fmt.Println(`
            [ Hacker at Work ]
                 _________
                |  _____  |
                | |     | |
                | | CLI | |
                | |_____| |
                |_________|
                    ||
        (•_•)      ||
       <|   |>=====||
        / \        ||
                  /__\
	`)
	
	fmt.Println(RedText(" [!] Educational & Professional Use Only"))
	fmt.Println(fmt.Sprintf(" [%s] Initializing internal modules...\n", "✔"))
}
