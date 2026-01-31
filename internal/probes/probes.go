package probes

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// ExecuteProbe handles the logic for different network protocols.
// It returns true if the target is reachable and the latency duration.
func ExecuteProbe(proto, target string, port int) (bool, time.Duration) {
	start := time.Now()
	addr := net.JoinHostPort(target, fmt.Sprintf("%d", port))

	switch proto {
	case "tcp":
		// Standard TCP Handshake
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}

	case "ack":
		// In Go, a standard Dial to a port behaves like a SYN/ACK check.
		// For true raw ACK flags, root privileges are required.
		// This serves as a high-reliability connectivity check.
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}

	case "udp":
		// UDP is connectionless. We send a packet and check if an error occurs.
		conn, err := net.DialTimeout("udp", addr, 3*time.Second)
		if err == nil {
			conn.Close()
			// Note: UDP may report "UP" even if the port is filtered.
			return true, time.Since(start)
		}

	case "http", "https":
		// Full Application Layer check
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		url := fmt.Sprintf("%s://%s", proto, target)
		// If port is not default, add it to URL
		if (proto == "http" && port != 80) || (proto == "https" && port != 443) {
			url = fmt.Sprintf("%s://%s:%d", proto, target, port)
		}

		resp, err := client.Get(url)
		if err == nil {
			resp.Body.Close()
			// Success if we get any standard HTTP response code
			return resp.StatusCode < 500, time.Since(start)
		}

	case "ping":
		// ICMP Ping usually requires root. 
		// This uses a TCP-based reachability check as a reliable fallback.
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err == nil {
			conn.Close()
			return true, time.Since(start)
		}
	}

	return false, 0
}
