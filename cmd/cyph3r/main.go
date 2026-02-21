package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// 1. Tactical Flag Definition
	target := flag.String("t", "", "Target Domain/IP")
	vector := flag.String("test", "", "Vector: HULK, SYN, UDP, ACK, TCP, HTTP, HTTPS")
	port := flag.String("port", "443", "Target Port")
	pps := flag.Int("pps", 50, "Packets Per Second Base")
	power := flag.Int("power", 100, "God-Mode Multiplier (Workers per tick)")
	monitor := flag.Bool("monitor", false, "Infinite Execution Mode")
	verbose := flag.Bool("v", false, "Enable Deep Cluster/Reverse DNS Recon")

	flag.Parse()

	// 2. Initialize HUD
	output.Banner()
	if *target == "" {
		flag.Usage()
		return
	}

	// 3. Deep Recon Phase
	data, err := intel.GetTargetIntel(*target)
	if err != nil {
		return
	}
	output.DisplayHUD(data, *verbose)

	// 4. Tactical Execution Phase
	if *vector != "" {
		ctx := context.Background()
		if !*monitor {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
		}

		// High-Capacity Transport to support the 100x Multiplier
		transport := &http.Transport{
			MaxIdleConns:        10000,
			MaxIdleConnsPerHost: 5000,
			IdleConnTimeout:     30 * time.Second,
			DisableKeepAlives:   false, // Critical for HULK/HTTPS vector speed
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}

		httpClient := &http.Client{
			Transport: transport,
			Timeout:   2 * time.Second,
		}

		// Engage the Engine
		intel.RunTacticalTest(models.TacticalConfig{
			Target: *target,
			Vector: *vector,
			PPS:    *pps,
			Port:   *port,
			Power:  *power,
		}, ctx, httpClient)
	}
}
