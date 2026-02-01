package intel

import (
	"fmt"
	"net"
	"strings"
	"github.com/likexian/whois-go"
)

// NodeIntel holds the unified intelligence profile
type NodeIntel struct {
	IPs       []string
	NS        []string
	Registrar string
	BornOn    string
	ISP       string
	Location  string
	Coords    string
}

// GetFullIntel performs a deep-dive reconnaissance on a target
func GetFullIntel(target string) (NodeIntel, error) {
	data := NodeIntel{}

	// 1. Resolve IP Addresses
	ips, _ := net.LookupIP(target)
	for _, ip := range ips {
		data.IPs = append(data.IPs, ip.String())
	}

	// 2. Resolve Name Servers
	ns, _ := net.LookupNS(target)
	for _, nameserver := range ns {
		data.NS = append(data.NS, nameserver.Host)
	}

	// 3. WHOIS Lookup
	w, err := whois.Whois(target)
	if err == nil {
		lines := strings.Split(w, "\n")
		for _, line := range lines {
			lowerLine := strings.ToLower(line)
			if strings.Contains(lowerLine, "creation date:") || strings.Contains(lowerLine, "created:") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 { data.BornOn = strings.TrimSpace(parts[1]) }
			}
			if strings.Contains(lowerLine, "registrar:") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 { data.Registrar = strings.TrimSpace(parts[1]) }
			}
		}
	}

	// Fallback for empty WHOIS
	if data.BornOn == "" { data.BornOn = "Protected/Unknown" }
	if data.Registrar == "" { data.Registrar = "Private" }

	return data, nil
}
