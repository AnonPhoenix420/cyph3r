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

// GHOST_MODE Safety Constants
const (
	RequestTimeout = 5 * time.Second
	UserAgent      = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			// Prevent connection reuse to avoid TCP fingerprinting during Ghost Mode
			DisableKeepAlives: true, 
		},
		Timeout: RequestTimeout,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// 1. GHOST_MODE KILL-SWITCH: Pre-flight check
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] GHOST_MODE FAILURE: VPN NOT DETECTED. EMERGENCY HALT.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// 2. NETWORK VECTORS (IPv4 Only to prevent IPv6 Leaks)
	ips, err := net.LookupIP(input)
	if err != nil {
		return data, fmt.Errorf("DNS_RESOLUTION_FAILED")
	}

	for _, ip := range ips {
		// Strict IPv4 filter for Ghost Anonymity
		if ip.To4() != nil {
			ipStr := ip.String()
			data.TargetIPs = append(data.TargetIPs, ipStr)
			
			// PTR Reverse lookup
			names, _ := net.LookupAddr(ipStr)
			if len(names) > 0 {
				data.ReverseDNS = append(data.ReverseDNS, strings.TrimSuffix(names[0], "."))
			} else {
				data.ReverseDNS = append(data.ReverseDNS, "NO_PTR")
			}
		}
	}

	// 3. GEO & INFRA ANALYSIS
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

	// 4. AUTHORITATIVE CLUSTERS
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		nsIPs, _ := net.LookupIP(nameserver.Host)
		for _, nsIP := range nsIPs {
			if nsIP.To4() != nil {
				data.NameServers[nameserver.Host] = append(data.NameServers[nameserver.Host], nsIP.String())
			}
		}
	}

	// 5. SHADOW RECON (Subdomain Hunter)
	huntSubdomains(input, &data)

	// 6. EXPLOIT SURFACE & WAF DETECTION
	analyzeExploitSurface(input, &data)
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)

	return data, nil
}

func huntSubdomains(target string, data *models.IntelData) {
	subs := []string{"dev", "vpn", "mail", "api", "test", "staging", "internal", "webmail", "cloud"}
	for _, s := range subs {
		host := s + "." + target
		ips, err := net.LookupIP(host)
		if err == nil && len(ips) > 0 {
			ipStr := ips[0].String()
			isWaf := false
			for _, mainIp := range data.TargetIPs {
				if ipStr == mainIp { isWaf = true }
			}
			if !isWaf {
				data.ScanResults = append(data.ScanResults, "DEBUG: Potential Origin Found ["+ipStr+"]")
			}
			data.ScanResults = append(data.ScanResults, fmt.Sprintf("SUBDOMAIN: %s → %s", host, ipStr))
		}
	}
}

func analyzeExploitSurface(target string, data *models.IntelData) {
	client := GetClient()
	
	// SCRUBBED REQUEST: Prevents metadata leakage
	req, _ := http.NewRequest("GET", "http://"+target, nil)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "*/*")

	resp, err := client.Do(req)
	if err != nil { return }
	defer resp.Body.Close()

	srv := resp.Header.Get("Server")
	if srv != "" {
		data.ScanResults = append(data.ScanResults, "STACK: "+srv)
		checkCVE(srv, data)
	}

	// WAF & LEAK DETECTION
	if sid := resp.Header.Get("X-Sid"); sid != "" {
		data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Node-ID ["+sid+"]")
	}
	
	if resp.Header.Get("CF-RAY") != "" {
		data.IsWAF, data.WAFType = true, "Cloudflare (Global CDN)"
	} else if strings.Contains(strings.ToLower(srv), "arvan") || resp.Header.Get("ArvanCloud-Trace") != "" {
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
	// Dialing through TCP 443 to measure signal latency
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var results []string
	ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087}
	
	for _, p := range ports {
		addr := net.JoinHostPort(target, fmt.Sprintf("%d", p))
		conn, err := net.DialTimeout("tcp", addr, 1500*time.Millisecond)
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
