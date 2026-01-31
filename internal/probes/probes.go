package probes

import (
	"net"
	"net/http"
	"time"
)

func RunProbe(proto, target string, port int) (bool, time.Duration) {
	start := time.Now()
	addr := fmt.Sprintf("%s:%d", target, port)
	
	switch proto {
	case "tcp", "ack": // ACK usually requires raw sockets; TCP Dial is the safe alternative
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	case "udp":
		conn, err := net.DialTimeout("udp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	case "http", "https":
		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(proto + "://" + target)
		if err == nil {
			resp.Body.Close()
			return resp.StatusCode < 400, time.Since(start)
		}
	}
	return false, 0
}
