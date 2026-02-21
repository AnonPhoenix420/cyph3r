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
	fmt.Printf("\n%s[GHOST_MODE] ENGAGING %s OMNI-VECTOR -> %s:%s%s\n", output.NeonPink, cfg.Vector, cfg.Target, cfg.Port, output.Reset)
	fmt.Printf("%s[!] FORCE_MULTIPLIER: %dx ACTIVE (%d Total Req/Sec)%s\n", output.Amber, cfg.Power, cfg.PPS*cfg.Power, output.Reset)

	ticker := time.NewTicker(time.Second / time.Duration(cfg.PPS))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s[+] Mission complete. All workers collapsed.%s\n", output.NeonGreen, output.Reset)
			return
		case <-ticker.C:
			// DYNAMIC POWER: Spawns workers based on the Power variable
			for i := 0; i < cfg.Power; i++ {
				go executeVector(cfg, client)
			}
		}
	}
}

// ... rest of executeVector stays the same ...
