package intel

import (
	"context"
	"net"
	"time"
)

// LookupNodes performs a forced DNS resolution to find A and NS records.
func LookupNodes(target string) ([]string, []string) {
	var ips []string
	var nsNodes []string

	// Force use of Google Public DNS to bypass local "Command Not Found" issues
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: time.Second * 5}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	// 1. Resolve IP Addresses (A Records)
	foundIPs, err := resolver.LookupIP(context.Background(), "ip", target)
	if err == nil {
		for _, ip := range foundIPs {
			ips = append(ips, ip.String())
		}
	}

	// 2. Resolve Nameservers (NS Records)
	foundNS, err := resolver.LookupNS(context.Background(), target)
	if err == nil {
		for _, ns := range foundNS {
			nsNodes = append(nsNodes, ns.Host)
		}
	}

	return ips, nsNodes
}
