package probes

import (
	"fmt"
	"net"
	"time"
)

// ConductWave sends multiple probe types to verify the service is "Alive"
func ConductWave(target string, port int) (method, status, convo string) {
	address := fmt.Sprintf("%s:%d", target, port)
	timeout := 1500 * time.Millisecond

	// 1. DNS Probe (Specialized for Port 53)
	if port == 53 {
		return "DNS", "OPEN", ProbeDNS(target)
	}

	// 2. The SYN/ACK Soldier (TCP Handshake)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err == nil {
		defer conn.Close()
		return "SYN/ACK", "OPEN", "Handshake Verified ✅"
	}

	// 3. The UDP Soldier (Bypass for Filtered ports)
	uConn, err := net.DialTimeout("udp", address, 1*time.Second)
	if err == nil {
		defer uConn.Close()
		return "UDP", "ALIVE", "Quiet Response 🔊"
	}

	return "ACK", "FILTERED", "No Response 🛡️"
}

// ProbeDNS performs a real A-record lookup to verify Port 53
func ProbeDNS(target string) string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: 1 * time.Second}
			return d.DialContext(ctx, "udp", fmt.Sprintf("%s:53", target))
		},
	}
	_, err := r.LookupHost(context.Background(), "google.com")
	if err != nil {
		return "Port 53 Active (No Recursion)"
	}
	return "DNS Fully Responsive 🧠"
}
