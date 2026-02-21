package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	target  := flag.String("t", "", "Target domain")
	testVec := flag.String("test", "", "Tactical Vector: SYN, UDP, HULK")
	pps     := flag.Int("pps", 10, "Packets Per Second")
	sec     := flag.Int("time", 30, "Duration in seconds")
	monitor := flag.Bool("monitor", false, "Keep running until manual stop")

	flag.Parse()
	output.Banner()

	shield := intel.CheckShield()
	if !shield.IsActive {
		fmt.Printf("\n\033[31m[!] GHOST_MODE FAILURE: VPN Required.\033[0m\n")
		os.Exit(1)
	}

	if *target != "" {
		// Run initial recon
		data, _ := intel.GetTargetIntel(*target)
		output.DisplayHUD(data, false)

		if *testVec != "" {
			var ctx context.Context
			var cancel context.CancelFunc

			if *monitor {
				fmt.Printf("\033[38;5;214m[!] MONITOR_MODE: Active. Press Ctrl+C to kill session.\033[0m\n")
				ctx, cancel = context.WithCancel(context.Background())
			} else {
				ctx, cancel = context.WithTimeout(context.Background(), time.Duration(*sec)*time.Second)
			}
			defer cancel()

			intel.RunTacticalTest(intel.TacticalConfig{
				Target: *target,
				Vector: *testVec,
				PPS:    *pps,
			}, ctx)
		}
	} else {
		flag.Usage()
	}
}
