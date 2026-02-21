package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, 
		Timeout:   6 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// 1. VPN KILL-SWITCH
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] PROTON VPN DISCONNECTED. EMERGENCY HALT.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{
		TargetName:  input,
		NameServers: make(map[string][]string),
	}

	// 2. NETWORK VECTORS & PTR
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		ipStr := ip.String()
		data.TargetIPs = append(data.TargetIPs, ipStr)
		names, _ := net.LookupAddr(ipStr)
		if len(names) > 0 {
			data.ReverseDNS = append(data.ReverseDNS, strings.TrimSuffix(names[0], "."))
		} else {
			data.ReverseDNS = append(data.ReverseDNS, "NO_PTR")
		}
	}

	// 3. GEO & ISP INTELLIGENCE
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country, data.Lat, data.Lon = geo.Org, geo.City, geo.Country, geo.Lat, geo.Lon
		data.RawGeo, data.Latency = raw, pingTarget(data.TargetIPs[0])
		usage := "RESIDENTIAL"; if geo.Hosting { usage = "DATA_CENTER" }; if geo.Proxy { usage += "/PROXY" }
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
	}

	// 4. AUTHORITATIVE CLUSTERS
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		nsIPs, _ := net.LookupIP(nameserver.Host)
		for _, nsIP := range nsIPs {
			data.NameServers[nameserver.Host] = append(data.NameServers[nameserver.Host], nsIP.String())
		}
	}

	// 5. EXPLOIT FINGERPRINTING (VULN SCANNER)
	analyzeExploitSurface(input, &data)

	// 6. PORT VECTORS
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)

	return data, nil
}

func analyzeExploitSurface(target string, data *models.IntelData) {
	client := GetClient()
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	srv := resp.Header.Get("Server")
	if srv != "" {
		data.ScanResults = append(data.ScanResults, "STACK: "+srv)
		s := strings.ToLower(srv)
		
		// OS Detection
		if strings.Contains(s, "ubuntu") { data.ScanResults = append(data.ScanResults, "OS_HINT: Linux (Ubuntu)") }
		if strings.Contains(s, "centos") { data.ScanResults = append(data.ScanResults, "OS_HINT: Linux (CentOS)") }
		if strings.Contains(s, "win") || strings.Contains(s, "iis") { data.ScanResults = append(data.ScanResults, "OS_HINT: Windows Server") }

		// Vuln Logic: Flag old versions
		if strings.Contains(s, "nginx/1.10") || strings.Contains(s, "apache/2.2") || strings.Contains(s, "php/5.") {
			data.ScanResults = append(data.ScanResults, "VULN_WARN: OUTDATED_SOFTWARE_DETECTED")
		}
	}

	// WAF Detection
	if resp.Header.Get("CF-RAY") != "" || strings.ToLower(srv) == "cloudflare" {
		data.IsWAF, data.WAFType = true, "Cloudflare (Global CDN)"
	} else if strings.Contains(strings.ToLower(srv), "arvancloud") || resp.Header.Get("ArvanCloud-Trace") != "" {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
	}

	// Framework Fingerprint
	if pwr := resp.Header.Get("X-Powered-By"); pwr != "" {
		data.ScanResults = append(data.ScanResults, "TECH: "+pwr)
	}
}

func fetchGeo(ip string) (models.GeoResponse, string) {
	client := GetClient()
	resp, err := client.Get("http://ip-api.com/json/" + ip + "?fields=status,country,city,isp,org,as,lat,lon,proxy,hosting,query")
	if err != nil { return models.GeoResponse{}, "{}" }
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var r models.GeoResponse
	json.Unmarshal(body, &r)
	var pretty interface{}
	json.Unmarshal(body, &pretty)
	prettyJSON, _ := json.MarshalIndent(pretty, "", "  ")
	return r, string(prettyJSON)
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var results []string
	ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087}
	for _, p := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(target, fmt.Sprintf("%d", p)), 1200*time.Millisecond)
		if err == nil {
			results = append(results, fmt.Sprintf("PORT %d: OPEN", p))
			conn.Close()
		}
	}
	return results
}

func GetPhoneIntel(num string) (models.PhoneData, error) {
	return models.PhoneData{Number: num, Carrier: "MCI/Irancell", Risk: "LOW", SocialPresence: []string{"TG", "WA"}}, nil
}
