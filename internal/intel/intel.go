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
		Timeout:   7 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] PROTON VPN DISCONNECTED. EMERGENCY HALT.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// 1. RESOLVE & PTR
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

	// 2. GEO & INFRA (Fixed Mislabeling)
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country, data.Lat, data.Lon = geo.Org, geo.City, geo.Country, geo.Lat, geo.Lon
		data.RawGeo, data.Latency = raw, pingTarget(data.TargetIPs[0])
		usage := "RESIDENTIAL"
		if geo.Hosting || strings.Contains(strings.ToLower(geo.Org), "arvan") || strings.Contains(strings.ToLower(geo.Isp), "anycast") {
			usage = "DATA_CENTER/CDN"
		}
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
	}

	// 3. AUTHORITATIVE CLUSTERS
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		nsIPs, _ := net.LookupIP(nameserver.Host)
		for _, nsIP := range nsIPs {
			data.NameServers[nameserver.Host] = append(data.NameServers[nameserver.Host], nsIP.String())
		}
	}

	// 4. SUBDOMAIN SHADOW-SCAN (The Origin Hunter)
	huntSubdomains(input, &data)

	// 5. EXPLOIT & BYPASS FINGERPRINTING
	analyzeExploitSurface(input, &data)
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)

	return data, nil
}

func huntSubdomains(target string, data *models.IntelData) {
	subs := []string{"dev", "vpn", "mail", "api", "test", "staging", "internal", "webmail"}
	for _, s := range subs {
		host := s + "." + target
		ips, err := net.LookupIP(host)
		if err == nil {
			result := fmt.Sprintf("SUBDOMAIN: %s → %s", host, ips[0].String())
			// Check if subdomain IP is DIFFERENT from main IP (Potential Origin!)
			isWaf := false
			for _, mainIp := range data.TargetIPs {
				if ips[0].String() == mainIp { isWaf = true }
			}
			if !isWaf {
				data.ScanResults = append(data.ScanResults, "DEBUG: Potential Origin Found ["+ips[0].String()+"]")
			}
			data.ScanResults = append(data.ScanResults, result)
		}
	}
}

func analyzeExploitSurface(target string, data *models.IntelData) {
	client := GetClient()
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	srv := resp.Header.Get("Server")
	if srv != "" {
		data.ScanResults = append(data.ScanResults, "STACK: "+srv)
		checkCVE(srv, data)
	}

	if sid := resp.Header.Get("X-Sid"); sid != "" {
		data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Node-ID ["+sid+"]")
	}
	if rid := resp.Header.Get("X-Request-Id"); rid != "" {
		data.ScanResults = append(data.ScanResults, "DEBUG: Trace-ID Detected")
	}

	if resp.Header.Get("CF-RAY") != "" {
		data.IsWAF, data.WAFType = true, "Cloudflare (Global CDN)"
	} else if strings.Contains(strings.ToLower(srv), "arvancloud") || resp.Header.Get("ArvanCloud-Trace") != "" {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
	}
}

func checkCVE(srv string, data *models.IntelData) {
	s := strings.ToLower(srv)
	if strings.Contains(s, "nginx/1.10") {
		data.ScanResults = append(data.ScanResults, "VULN_WARN: CVE-2021-23017 (RCE)")
	}
	if strings.Contains(s, "apache/2.4.49") {
		data.ScanResults = append(data.ScanResults, "VULN_WARN: CVE-2021-41773 (Path Traversal)")
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
		if conn, err := net.DialTimeout("tcp", net.JoinHostPort(target, fmt.Sprintf("%d", p)), 1200*time.Millisecond); err == nil {
			results = append(results, fmt.Sprintf("PORT %d: OPEN", p))
			conn.Close()
		}
	}
	return results
}

func GetPhoneIntel(num string) (models.PhoneData, error) {
	return models.PhoneData{Number: num, Carrier: "MCI/Irancell", Risk: "LOW", SocialPresence: []string{"TG", "WA"}}, nil
}
