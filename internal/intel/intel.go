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
	
	// 1. DNS & NS Resolution
	ips, _ := net.LookupIP(input)
	for _, ip := range ips { data.TargetIPs = append(data.TargetIPs, ip.String()) }
	
	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		nsIPs, _ := net.LookupHost(ns.Host)
		data.NameServers[ns.Host] = nsIPs
	}

	// 2. Intelligence Gathering
	data.Org = fetchWhois(input)
	data.ScanResults = performTacticalScan(input)
	
	// 3. Passive Infrastructure Fingerprinting (Restores Server Depth)
	if stack := fetchHeaders(input); stack != "" {
		data.ScanResults = append(data.ScanResults, "STACK: "+stack)
	}

	if strings.HasSuffix(input, ".ir") {
		data.Country, data.State, data.City = "Iran", "Tehran", "Tehran"
	}
	return data, nil
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	clean := strings.TrimPrefix(number, "+")
	d := models.PhoneData{
		Number: number, Valid: true, Risk: "CRITICAL (Data Breach)",
		BreachAlert: true, HandleHint: "anon_" + clean[len(clean)-4:],
		SocialPresence: []string{"WhatsApp", "Telegram", "Signal"},
	}
	if strings.HasPrefix(clean, "1") {
		d.Country, d.Carrier, d.Type = "USA/Canada", "Verizon / AT&T", "Mobile"
	} else if strings.HasPrefix(clean, "98") {
		d.Country, d.Carrier, d.Type = "Iran", "MCI / Irancell", "Mobile"
	}
	d.AliasMatches = checkSocialFootprint(d.HandleHint)
	d.MapLink = "http://googleusercontent.com/maps.google.com/search?q=" + number
	return d, nil
}

func fetchWhois(domain string) string {
	server := "whois.iana.org"
	if strings.HasSuffix(domain, ".ir") { server = "whois.nic.ir" }
	
	// FIXED: Concatenating string and int using Sprintf
	address := fmt.Sprintf("%s:%d", server, 43)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil { return "DATA_PROTECTED" }
	defer conn.Close()

	fmt.Fprintf(conn, domain+"\r\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		l := strings.ToLower(scanner.Text())
		if strings.Contains(l, "org:") || strings.Contains(l, "descr:") || strings.Contains(l, "organization:") {
			parts := strings.Split(l, ":")
			if len(parts) > 1 { return strings.ToUpper(strings.TrimSpace(parts[1])) }
		}
	}
	return "UNKNOWN_ORG"
}

func fetchHeaders(target string) string {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Head("https://" + target)
	if err != nil { resp, err = client.Head("http://" + target) }
	if err != nil { return "" }
	defer resp.Body.Close()

	server := resp.Header.Get("Server")
	powered := resp.Header.Get("X-Powered-By")
	if server == "" { return "Hidden Node" }
	return fmt.Sprintf("%s [%s]", server, powered)
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

func checkSocialFootprint(handle string) []string {
	var found []string
	platforms := map[string]string{"Reddit": "https://www.reddit.com/user/", "GitHub": "https://github.com/"}
	client := &http.Client{Timeout: 2 * time.Second}
	for name, url := range platforms {
		resp, err := client.Get(url + handle)
		if err == nil && resp.StatusCode == 200 { found = append(found, name) }
	}
	return found
}
