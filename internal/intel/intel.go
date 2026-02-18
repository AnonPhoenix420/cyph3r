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
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// Security Handshake
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] PROTON VPN NOT DETECTED. KILLING PROCESS.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	
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
	
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		
		usage := "RESIDENTIAL"
		if geo.Hosting { usage = "DATA_CENTER" }
		if geo.Proxy { usage += "/PROXY" }
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
		data.RawGeo = raw
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	detectWAF(input, &data)
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)
	return data, nil
}

// --- CORE HELPERS ---

func fetchGeo(ip string) (models.GeoResponse, string) {
	client := GetClient()
	resp, err := client.Get("http://ip-api.com/json/" + ip + "?fields=status,country,regionName,city,isp,org,as,lat,lon,proxy,hosting,query")
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

func detectWAF(target string, data *models.IntelData) {
	client := GetClient()
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	server := strings.ToLower(resp.Header.Get("Server"))
	cfRay := resp.Header.Get("CF-RAY")

	if cfRay != "" || server == "cloudflare" {
		data.IsWAF = true
		data.WAFType = "Cloudflare (Global CDN)"
	} else if strings.Contains(server, "arvancloud") {
		data.IsWAF = true
		data.WAFType = "ArvanCloud (Regional WAF)"
	}
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var res []string
	ports := []int{80, 443, 8080}
	for _, p := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(target, fmt.Sprintf("%d", p)), 1500*time.Millisecond)
		if err == nil {
			conn.Close()
			res = append(res, fmt.Sprintf("PORT %d: OPEN", p))
		}
	}
	return res
}

func GetPhoneIntel(num string) (models.PhoneData, error) {
	return models.PhoneData{Number: num, Country: "US", Carrier: "Global"}, nil
}
