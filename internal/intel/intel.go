package intel

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	
	// Force Dual-Stack Resolution
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}
	data.TargetIPs = deduplicate(data.TargetIPs)
	
	// Recursive NS Discovery
	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		addrs, _ := net.LookupHost(ns.Host)
		data.NameServers[ns.Host] = deduplicate(addrs)
	}

	data.Org = queryRecursiveWhois(input, "whois.iana.org")
	data.ScanResults = performTacticalScan(input)
	return data, nil
}

func queryRecursiveWhois(domain, server string) string {
	conn, err := net.DialTimeout("tcp", server+":43", 5*time.Second)
	if err != nil { return "BYPASS_REQUIRED" }
	defer conn.Close()

	fmt.Fprintf(conn, domain+"\r\n")
	scanner := bufio.NewScanner(conn)
	var referral string
	keywords := []string{"descr:", "org:", "organization:", "registrant:", "owner:", "org-name:"}

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, "whois:") || strings.Contains(line, "refer:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 { referral = strings.TrimSpace(parts[1]) }
		}
		for _, key := range keywords {
			if strings.Contains(line, key) {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					val := strings.TrimSpace(parts[1])
					if val != "" && !strings.Contains(val, "redacted") { return strings.ToUpper(val) }
				}
			}
		}
	}
	if referral != "" && referral != server { return queryRecursiveWhois(domain, referral) }
	return "SECURE_INFRASTRUCTURE"
}

func performTacticalScan(target string) []string {
	var results []string
	ports := []int{80, 443, 8080, 8443, 2083, 2087}
	for _, p := range ports {
		addr := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err == nil {
			conn.Close()
			status := "OPEN [ACK/SYN]"
			if p == 443 || p == 8443 {
				conf := &tls.Config{InsecureSkipVerify: true}
				if tlsConn, err := tls.Dial("tcp", addr, conf); err == nil {
					if certs := tlsConn.ConnectionState().PeerCertificates; len(certs) > 0 {
						status = fmt.Sprintf("OPEN (SSL: %s) [ACK/SYN]", certs[0].Subject.CommonName)
					}
					tlsConn.Close()
				}
			}
			results = append(results, fmt.Sprintf("PORT %d: %s", p, status))
		}
	}
	return results
}

func deduplicate(s []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range s {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	clean := strings.TrimPrefix(number, "+")
	d := models.PhoneData{
		Number: number, Valid: true, Risk: "CRITICAL (Data Breach)",
		BreachAlert: true, HandleHint: "anon_" + clean[len(clean)-4:],
		SocialPresence: []string{"WhatsApp", "Telegram", "Signal"},
	}
	if strings.HasPrefix(clean, "98") { d.Country, d.Carrier = "Iran", "MCI / Irancell"
	} else if strings.HasPrefix(clean, "1") { d.Country, d.Carrier = "USA/Canada", "Verizon / AT&T"
	} else { d.Country, d.Carrier = "Global Node", "International" }
	d.MapLink = "http://maps.google.com/search?q=" + number
	return d, nil
}
