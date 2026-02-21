package intel

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	data := models.IntelData{
		TargetName:  input,
		NameServers: make(map[string][]string),
	}

	// 1. DNS & Cluster Recon
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[s.Host] = ipStrings
	}

	// 2. IP Mapping & Reverse DNS (The Fix)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { 
			data.TargetIPs = append(data.TargetIPs, ip.String()) 
			names, _ := net.LookupAddr(ip.String())
			data.ReverseDNS = append(data.ReverseDNS, names...)
		}
	}

	// 3. Deep Geo & Tactical Flags (The Fix)
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 3 * time.Second}
		// Fetching all fields including proxy/hosting/mobile
		resp, _ := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			data.Org = g.Org
			data.City, data.Region, data.Country = g.City, g.RegionName, g.Country
			data.Lat, data.Lon = g.Lat, g.Lon
			data.RawGeo = string(body)
			
			// Inject flags into ScanResults for the HUD
			if g.Proxy { data.ScanResults = append(data.ScanResults, "DETECTION: PROXY/VPN") }
			if g.Hosting { data.ScanResults = append(data.ScanResults, "DETECTION: DATACENTER") }
			if g.Mobile { data.ScanResults = append(data.ScanResults, "DETECTION: MOBILE_NET") }
			resp.Body.Close()
		}
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	analyzeWAF(input, &data)
	return data, nil
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1000*time.Millisecond)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return time.Since(start).String()
}

func analyzeWAF(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get("https://" + target)
	if err != nil { return }
	defer resp.Body.Close()
	srv := resp.Header.Get("Server")
	if srv != "" {
		data.IsWAF, data.WAFType = true, srv
	}
}
