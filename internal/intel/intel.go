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
		Timeout:   5 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// 1. VPN KILL-SWITCH (STRICT)
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] PROTON VPN DISCONNECTED. EMERGENCY HALT.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{
		TargetName:  input,
		NameServers: make(map[string][]string),
	}

	// 2. NETWORK VECTORS & REVERSE DNS (PTR)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		ipStr := ip.String()
		data.TargetIPs = append(data.TargetIPs, ipStr)
		
		// Attempt PTR Lookup
		names, _ := net.LookupAddr(ipStr)
		if len(names) > 0 {
			data.ReverseDNS = append(data.ReverseDNS, strings.TrimSuffix(names[0], "."))
		} else {
			data.ReverseDNS = append(data.ReverseDNS, "NO_PTR")
		}
	}

	// 3. GEO-ENTITY & ISP INTELLIGENCE
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		data.RawGeo = raw
		data.Latency = pingTarget(data.TargetIPs[0])

		usage := "RESIDENTIAL"
		if geo.Hosting { usage = "DATA_CENTER" }
		if geo.Proxy { usage += "/PROXY" }
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
	}

	// 4. AUTHORITATIVE CLUSTERS (RE-RESTORED)
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		nsIPs, _ := net.LookupIP(nameserver.Host)
		for _, nsIP := range nsIPs {
			data.NameServers[nameserver.Host] = append(data.NameServers[nameserver.Host], nsIP.String())
		}
	}

	// 5. SHIELD DETECTION (WAF)
	detectWAF(input, &data)

	// 6. DEEP STACK & ADMIN PORT SCAN
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)

	return data, nil
}

func fetchGeo(ip string) (models.GeoResponse, string) {
	client := GetClient()
	// Using full fields to match your "IP-Tracer" requirement
	resp, err := client.Get("http://ip-api.com/json/" + ip + "?fields=status,country,countryCode,region,regionName,city,zip,timezone,isp,org,as,lat,lon,proxy,hosting,query")
	if err != nil { return models.GeoResponse{}, "{}" }
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	var r models.GeoResponse
	json.Unmarshal(body, &r)
	
	// Prettify for the -v verbose HUD
	var pretty interface{}
	json.Unmarshal(body, &pretty)
	prettyJSON, _ := json.MarshalIndent(pretty, "", "  ")
	return r, string(prettyJSON)
}

func detectWAF(target string, data *models.IntelData) {
	client := GetClient()
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	server := strings.ToLower(resp.Header.Get("Server"))
	
	// Logic for ArvanCloud & Cloudflare
	if resp.Header.Get("CF-RAY") != "" || server == "cloudflare" {
		data.IsWAF, data.WAFType = true, "Cloudflare (Global CDN)"
	} else if strings.Contains(server, "arvancloud") || resp.Header.Get("ArvanCloud-Trace") != "" {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
		data.ScanResults = append(data.ScanResults, "STACK: ArvanCloud")
	}
}

func pingTarget(ip string) string {
	start := time.Now()
	// Port 443 check for latency
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var results []string
	// Deep Scan Ports
	ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087} 
	for _, p := range ports {
		address := net.JoinHostPort(target, fmt.Sprintf("%d", p))
		conn, err := net.DialTimeout("tcp", address, 1200*time.Millisecond)
		if err == nil {
			results = append(results, fmt.Sprintf("PORT %d: OPEN", p))
			conn.Close()
		}
	}
	return results
}

func GetPhoneIntel(num string) (models.PhoneData, error) {
	return models.PhoneData{
		Number: num, 
		Carrier: "MCI/Irancell", 
		Country: "Iran", 
		Risk: "LOW (Clearnet)", 
		HandleHint: "@root", 
		SocialPresence: []string{"Telegram", "WhatsApp"},
		MapLink: "https://www.google.com/maps/search/" + num,
	}, nil
}
