package intel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)

	// 1. IP Resolution
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// 2. Recursive DNS
	ns, _ := net.LookupNS(input)
	for _, n := range ns {
		data.NameServers["NS"] = append(data.NameServers["NS"], n.Host)
		nsIPs, _ := net.LookupIP(n.Host)
		for _, nip := range nsIPs {
			data.NameServers["IP_"+n.Host] = append(data.NameServers["IP_"+n.Host], nip.String())
		}
	}

	// 3. WHOIS (Direct Socket - No External Libs)
	data.Org = fetchWhois(input)

	// 4. Tactical Scan
	data.NameServers["PORTS"] = probes.ScanPorts(input)

	// 5. Geo-Intel
	if len(data.TargetIPs) > 0 {
		fetchGeoData(&data)
	}

	return data, nil
}

func fetchWhois(domain string) string {
	// Standard WHOIS server for most domains
	whoisServer := "whois.iana.org" 
	
	// Step 1: Query IANA to find the real registrar server
	conn, err := net.DialTimeout("tcp", whoisServer+":43", 5*time.Second)
	if err != nil { return "OFFLINE" }
	
	fmt.Fprintf(conn, domain+"\r\n")
	scanner := bufio.NewScanner(conn)
	var registrar, created string
	
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, "registrar:") {
			registrar = strings.TrimSpace(strings.Split(line, ":")[1])
		}
		if strings.Contains(line, "created:") {
			created = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}
	conn.Close()

	if registrar != "" {
		return fmt.Sprintf("%s (Born: %s)", registrar, created)
	}
	return "DATA_MASKED"
}

func fetchGeoData(data *models.IntelData) {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0])
	if err != nil { return }
	defer resp.Body.Close()

	var t struct {
		Country, RegionName, City, Zip, Isp, Org string
	}
	json.NewDecoder(resp.Body).Decode(&t)
	data.Country, data.State, data.City, data.Zip = t.Country, t.RegionName, t.City, t.Zip
	data.ISP = t.Isp
	if data.Org == "" { data.Org = t.Org }
}
