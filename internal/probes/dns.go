package probes

import (
	"context"
	"fmt"
	"net"
	"time"
)

func CheckDNS(target string) string {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: time.Second}
			return d.DialContext(ctx, "udp", fmt.Sprintf("%s:53", target))
		},
	}
	// Try a real lookup; if it works, the port is truly open.
	_, err := resolver.LookupHost(context.Background(), "google.com")
	if err != nil {
		return "FILTERED/NO_RECURSION"
	}
	return "DNS_ALIVE"
}
