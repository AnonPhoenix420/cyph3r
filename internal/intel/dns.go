package intel

import (
	"net"
)

// LookupNodes resolves the A and NS records for a target.
func LookupNodes(target string) ([]string, []string) {
	var ips []string
	var nss []string

	// Resolve IPs
	addr, err := net.LookupHost(target)
	if err == nil {
		ips = addr
	}

	// Resolve Nameservers
	nsRecords, err := net.LookupNS(target)
	if err == nil {
		for _, ns := range nsRecords {
			nss = append(nss, ns.Host)
		}
	}

	return ips, nss
}
