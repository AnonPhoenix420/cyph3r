cat <<EOF > internal/output/banner.go
package output

import (
	"fmt"
	"time"
)

func Banner() {
	banner := \`
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \\
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM\`

	fmt.Println(CyanText(banner))
	fmt.Printf("\n %s\n", WhiteText("⚡ v2.6 [STABLE] // Wireframe HUD Edition"))
	fmt.Println(CyanText(" ───────────────────────────────────────"))
}

func ScanAnimation() {
	fmt.Print(WhiteText("[*] Calibrating HUD Sensors... "))
	time.Sleep(600 * time.Millisecond)
	fmt.Println(GreenText("[READY]"))
}
EOF
