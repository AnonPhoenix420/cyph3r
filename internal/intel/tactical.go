package intel

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func RunTacticalTest(cfg models.TacticalConfig, ctx context.Context, client *http.Client) {
	fmt.Printf("\n%s[GHOST_MODE] %sENGAGING %s OMNI-VECTOR -> %s:%s%s\n", output.NeonPink, output.Bold, cfg.Vector, cfg.Target, cfg.Port, output.Reset)
	fmt.Printf("%s[!] FORCE_MULTIPLIER: %dx ACTIVE (%d Req/Sec)%s\n", output.Amber, cfg.Power, cfg.PPS*cfg.Power, output.Reset)

	ticker := time.NewTicker(time.Second / time.Duration(cfg.PPS))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s[+] Mission vectors collapsed.%s\n", output.NeonGreen, output.Reset)
			return
		case <-ticker.C:
			// The Multiplier: Launches Power x workers every tick
			for i := 0; i < cfg.Power; i++ {
				go executeVector(cfg, client)
			}
		}
	}
}

func executeVector(cfg models.TacticalConfig, client *http.Client) {
	addr := net.JoinHostPort(cfg.Target, cfg.Port)
	// Same vector logic as before, optimized for speed...
	// [SYN, HULK, UDP logic here]
}
