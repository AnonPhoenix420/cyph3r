package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
)

// ================= WHOIS =================

func whoisLookup(target string) (string, error) {
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	_, _ = conn.Write([]byte(target + "\n"))
	buf, _ := io.ReadAll(conn)
	return string(buf), nil
}

// ================= DNS =================

func resolveDNS(target string) ([]string, error) {
	ips, err := net.LookupHost(target)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

// ================= ASN / CIDR =================

type ASNInfo struct {
	ASN   string
	Org   string
	CIDRs []string
}

func lookupASN(ip string) (*ASNInfo, error) {
	// placeholder, replace with real ASN service lookup
	return &ASNInfo{
		ASN:   "AS12345",
		Org:   "Example Org",
		CIDRs: []string{ip + "/24"},
	}, nil
}

// ================= TLS =================

func inspectTLS(target string) (*tls.ConnectionState, error) {
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", fmt.Sprintf("%s:443", target), &tls.Config{})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return &conn.ConnectionState(), nil
}

// ================= ICMP =================

func pingICMP(target string) (latency time.Duration, err error) {
	// placeholder: implement with golang.org/x/net/icmp if needed
	return 0, nil
}

func jitterICMP(target string, count int) (avg, max, min time.Duration, loss float64, err error) {
	// placeholder: loop pingICMP 'count' times and calculate metrics
	return 0, 0, 0, 0, nil
}

// ================= PHONE =================

func lookupPhone(number string) map[string]interface{} {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	return map[string]interface{}{
		"raw":    number,
		"valid":  phonenumbers.IsValidNumber(num),
		"region": phonenumbers.GetRegionCodeForNumber(num),
		"type":   phonenumbers.GetNumberType(num),
	}
}

// ================= HELPER =================

func printJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
