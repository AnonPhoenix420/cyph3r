package probes

import (
	"fmt"
	"net"
	"time"
)

// ConductWave performs a high-speed TCP handshake probe on a specific port.
func ConductWave(target string, port int) (string, string, string) {
	address := fmt.Sprintf("%s:%d", target, port)
	
	// Setting a 2-second timeout to keep the "Wave" moving fast
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	
	if err != nil {
		// If the connection is refused or times out
		return "TCP_STEALTH", "OFFLINE", "NO_RESPONSE"
	}
	
	// If we successfully connected, the port is open
	defer conn.Close()
	
	return "TCP_FULL_HANDSHAKE", "ALIVE", "SYN_ACK_RECEIVED"
}
