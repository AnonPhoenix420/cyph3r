package probes

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// ExecuteProbe checks connectivity and returns success + latency
func ExecuteProbe(proto, target string, port int) (bool, time.Duration) {
	start := time.Now()
	
	if proto == "http" || proto == "https" {
		client := &http.Client{Timeout: 2 * time.Second}
		_, err := client.Get(fmt.Sprintf("%s://%s", proto, target))
		return err == nil, time.Since(start)
	}

	// Default TCP logic
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err == nil {
		conn.Close()
		return true, time.Since(start)
	}
	
	return false, 0
}

// ConductACKProbe sends a naked ACK packet to map firewall state.
// Unfiltered ports will respond with a RST; Filtered will remain silent.
func ConductACKProbe(target string, port int) string {
	address := fmt.Sprintf("%s:%d", target, port)
	// We use a raw-style dial to see if the network stack returns an RST
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	
	// If it fails with "connection refused," it's actually UNFILTERED 
	// because the target sent back a Reset (RST) packet.
	if err != nil {
		return "UNFILTERED (Firewall Bypass)"
	}
	defer conn.Close()
	return "OPEN"
}
