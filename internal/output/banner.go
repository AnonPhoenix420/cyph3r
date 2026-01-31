package output

import (
	"fmt"
	"runtime"
)

func Banner() {
	ascii := `
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM`

	fmt.Println(Cyan + Bold + ascii + Reset)
	fmt.Println(Gray + " ────────────────────────────────────────────────────────" + Reset)
	fmt.Printf(" %s[SYS]%s OS: %-7s | ARCH: %-7s | BUILD: 2.6\n", Bold, Reset, runtime.GOOS, runtime.GOARCH)
}
