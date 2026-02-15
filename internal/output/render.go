package output

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func PulseNode(target string) {
	fmt.Printf("\n%s[!] Identifying Node: %s%s%s\n", White, NeonPink, target, Reset)
}

func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ TARGET_HUD ] ---\n%s[*] Node: %s%s\n", NeonPink, White, NeonBlue, data.TargetName)
}

// Ensure internal/output/banner.go has: func DisplayBanner()
