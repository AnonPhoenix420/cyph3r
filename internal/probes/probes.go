package probes

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func ExecuteProbe(proto, target string, port int) (bool, time.Duration) {
	start := time.Now()
	addr := net.JoinHostPort(target, fmt.Sprintf("%d", port))

	switch proto {
	case "tcp", "ack":
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	case "udp":
		conn, err := net.DialTimeout("udp", addr, 2*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	case "http", "https":
		client := http.Client{Timeout: 5 * time.Second}
		url := fmt.Sprintf("%s://%s:%d", proto, target, port)
		resp, err := client.Get(url)
		if err == nil {
			resp.Body.Close()
			return true, time.Since(start)
		}
	case "ping":
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	}
	return false, 0
}
