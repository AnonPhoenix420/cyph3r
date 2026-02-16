package intel

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)

	// 1. IP & Recursive DNS
	ips, _ := net.LookupIP(input)
	for _, ip := range ips { data.TargetIPs = append(data.TargetIPs, ip.String()) }

	ns, _ := net.LookupNS(input)
	for _, n := range ns {
		data.NameServers["NS"] = append(data.NameServers["NS"], n.Host)
		nsIPs, _ := net.LookupIP(n.Host)
		for _, nip := range nsIPs {
			data.NameServers["IP_"+n.Host] = append(data.NameServers["IP_"+n.Host], nip.String())
		}
	}

	// 2. Subdomain Discovery (Passive + Active)
	data.Subdomains = discoverSubdomains(input)

	// 3. Built-in WHOIS (Direct Socket)
	data.Org = fetchWhois(input)

	// 4. Tactical Service Scan
	data.NameServers["PORTS"] = probes.ScanPorts(input)

	// 5. Geo-Intelligence
	if len(data.TargetIPs) > 0 { fetchGeoData(&data) }

	return data, nil
}

func discoverSubdomains(domain string) []string {
	var found []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	wordlist := []string{"www", "mail", "vpn", "dev", "api", "staff", "portal", "ssh", "ftp"}

	for _, sub := range wordlist {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			target := s + "." + domain
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			if _, err := net.DefaultResolver.LookupHost(ctx, target); err == nil {
				mu.Lock()
				found = append(found, target)
				mu.Unlock()
			}
		}(sub)
	}
	wg.Wait()
	return found
}

func fetchWhois(domain string) string {
	server := "whois.iana.org"
	if strings.HasSuffix(domain, ".ir") { server = "whois.nic.ir" }
	conn, err := net.DialTimeout("tcp", server+":43", 5*time.Second)
	if err != nil { return "RESTRICTED" }
	defer conn.Close()
	fmt.Fprintf(conn, domain+"\r\n")
	scanner := bufio.NewScanner(conn)
	var reg, born string
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, "registrar:") || strings.Contains(line, "source:") {
			reg = strings.TrimSpace(strings.Split(line, ":")[1])
		}
		if strings.Contains(line, "created:") || strings.Contains(line, "last-updated:") {
			born = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}
	return fmt.Sprintf("%s (Update: %s)", reg, born)
}

func fetchGeoData(data *models.IntelData) {
	resp, err := http.Get("http://ip-api.com/json/" + data.TargetIPs[0])
	if err != nil { return }
	defer resp.Body.Close()
	var t struct{ Country, RegionName, City, Zip, Isp, Org string }
	json.NewDecoder(resp.Body).Decode(&t)
	data.Country, data.State, data.City, data.Zip, data.ISP = t.Country, t.RegionName, t.City, t.Zip, t.Isp
	if data.Org == "" || strings.Contains(data.Org, "RESTRICTED") { data.Org = t.Org }
}
