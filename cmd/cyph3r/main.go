package main

import (
	"context"
	"flag"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	target := flag.String("t", "", "Target Domain")
	vector := flag.String("test", "", "Vector: HULK, SYN, UDP, ACK, DNS, ICMP, TCP, HTTP, HTTPS")
	port := flag.String("port", "443", "Target Port")
	pps := flag.Int("pps", 20, "Speed (PPS)")
	monitor := flag.Bool("monitor", false, "Infinite Mode")

	flag.Parse()
	output.Banner()

	if *target == "" {
		flag.Usage()
		return
	}

	// 1. RECON PHASE (Persistence Restored)
	data, _ := intel.GetTargetIntel(*target)
	output.DisplayHUD(data, false)

	// 2. TACTICAL PHASE (Omni-Vector Engine)
	if *vector != "" {
		ctx := context.Background()
		if !*monitor {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
		}

		intel.RunTacticalTest(models.TacticalConfig{
			Target: *target,
			Vector: *vector,
			PPS:    *pps,
			Port:   *port,
		}, ctx)
	}
}
