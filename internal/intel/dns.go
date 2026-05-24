package intel

import (
	"net"
	"strings"
)

// DNSRecordSet holds comprehensive DNS intelligence for a target
type DNSRecordSet struct {
	MXRecords  []string `json:"mx_records"`
	ARecords   []string `json:"a_records"`
	NSRecords  []string `json:"ns_records"`
	TXTRecords []string `json:"txt_records"`
	ReverseDNS string   `json:"reverse_dns"`
}

// LookupMXRecords returns host lookup configurations for target nodes (your original function)
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

// GetFullDNSRecords performs comprehensive DNS reconnaissance
func GetFullDNSRecords(target string) DNSRecordSet {
	domain := target
	if strings.Contains(target, "@") {
		domain = strings.Split(target, "@")[1]
	}

	dns := DNSRecordSet{}

	// MX Records
	if mx, err := LookupMXRecords(domain); err == nil {
		dns.MXRecords = mx
	}

	// A Records
	if ips, err := net.LookupIP(domain); err == nil {
		for _, ip := range ips {
			dns.ARecords = append(dns.ARecords, ip.String())
		}
	}

	// NS Records
	if ns, err := net.LookupNS(domain); err == nil {
		for _, record := range ns {
			dns.NSRecords = append(dns.NSRecords, record.Host)
		}
	}

	// TXT Records
	if txt, err := net.LookupTXT(domain); err == nil {
		dns.TXTRecords = txt
	}

	// Reverse DNS (if we have an IP)
	if len(dns.ARecords) > 0 {
		names, err := net.LookupAddr(dns.ARecords[0])
		if err == nil && len(names) > 0 {
			dns.ReverseDNS = names[0]
		}
	}

	return dns
}
