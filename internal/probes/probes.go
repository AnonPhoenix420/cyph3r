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
