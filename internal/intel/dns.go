package intel

import (
	"net"
)

// LookupMXRecords returns host lookup configurations for target nodes
func LookupMXRecords(domain string) ([]string, error) {
	var mxRecords []string
	mx, err := net.LookupMX(domain)
	if err != nil {
		return mxRecords, err
	}
	for _, record := range mx {
		mxRecords = append(mxRecords, record.Host)
	}
	return mxRecords, nil
}
