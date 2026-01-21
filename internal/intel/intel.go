package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"cyph3r/output"

	"github.com/nyaruka/phonenumbers"
)

// ================= WHOIS =================

// whoisLookup performs a basic WHOIS query
func whoisLookup(target string) (string, error) {
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, _ = conn.Write([]byte(target + "\r\n"))
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	return string(buf[:n]), nil
}

// ================= DNS =================

// resolveDNS returns A, AAAA, MX, and NS records for a host
func resolveDNS(host string) (map[string][]string, error) {
	records := make(map[string][]string)

	// A records
	if ips, err := net.LookupIP(host); err == nil {
		for _, ip := range ips {
			if ip.To4() != nil {
				records["A"] = append(records["A"], ip.String())
			} else {
				records["AAAA"] = append(records["AAAA"], ip.String())
			}
		}
	}

	// MX
	if mxs, err := net.LookupMX(host); err == nil {
		for _, mx := range mxs {
			records["MX"] = append(records["MX"], mx.Host)
		}
	}

	// NS
	if nss, err := net.LookupNS(host); err == nil {
		for _, ns := range nss {
			records["NS"] = append(records["NS"], ns.Host)
		}
	}

	return records, nil
}

// ================= ASN & CIDR =================

// expandASN converts ASN to CIDR range via placeholder
func expandASN(asn string) ([]string, error) {
	// For simplicity, stub implementation
	// In production, query services like bgpview.io API
	cidrs := []string{"192.0.2.0/24", "198.51.100.0/24"}
	return cidrs, nil
}

// ================= TLS CERT =================

// inspectTLS connects to host:port and retrieves certificate info
func inspectTLS(host string, port int) (*tls.Certificate, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		return certs[0], nil
	}
	return nil, fmt.Errorf("no cert found")
}

// ================= PHONE =================

// lookupPhone parses a phone number and returns metadata
func lookupPhone(number string) {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		output.Down("Phone parse error: " + err.Error())
		return
	}

	data := map[string]interface{}{
		"raw":    number,
		"valid":  phonenumbers.IsValidNumber(num),
		"region": phonenumbers.GetRegionCodeForNumber(num),
		"type":   phonenumbers.GetNumberType(num),
	}

	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}

// ================= HELPER =================

// jsonPrint helper
func jsonPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

// ================= RECURSIVE WHOIS =================

// recursiveWhois performs WHOIS and follows referrals automatically
func recursiveWhois(domain string) {
	visited := map[string]bool{}
	current := domain

	for {
		if visited[current] {
			break
		}
		visited[current] = true

		result, err := whoisLookup(current)
		if err != nil {
			output.Down("WHOIS failed for " + current + ": " + err.Error())
			break
		}

		fmt.Println(result)

		ref := parseWhoisReferral(result)
		if ref == "" {
			break
		}
		current = ref
	}
}

// parseWhoisReferral extracts the referral WHOIS server from raw text
func parseWhoisReferral(raw string) string {
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), "whois:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}
