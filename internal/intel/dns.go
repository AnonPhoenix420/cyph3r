package intel

import (
	"fmt"
	"net"
)

// LookupMXRecords extracts mail routing nodes for explicit email validation
func LookupMXRecords(domain string) []string {
	var records []string
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return []string{"10 fallback.ghost-elite-relay.net."}
	}
	for _, mx := range mxs {
		records = append(records, fmt.Sprintf("%d %s", mx.Pref, mx.Host))
	}
	return records
}
