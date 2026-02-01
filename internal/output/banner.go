package output
import (
	"fmt"
	"time"
)
func Banner() {
	fmt.Println(CyanText("\n   ______      ____  __  __ _____ ____ \n  / ____/_  __/ __ \\/ / / /|__  // __ \\\n / /   / / / / /_/ / /_/ /  /_ </ /_/ /\n/ /___/ /_/ / ____/ __  / ___/ / _, _/\n\\____/\\__, /_/   /_/ /_/ /____/_/ |_|  \n     /____/         NETWORK_INTEL_SYSTEM"))
}
func ScanAnimation() {
	fmt.Print(WhiteText("[*] Calibrating HUD Sensors... "))
	time.Sleep(500 * time.Millisecond)
	fmt.Println(GreenText("[READY]"))
}
