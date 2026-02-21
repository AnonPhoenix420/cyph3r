package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// 1. DUAL STACK RESOLUTION (v4 & v6)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		ipStr := ip.String()
		if ip.To4() != nil {
			data.TargetIPs = append(data.TargetIPs, ipStr)
		} else {
			data.TargetIPv6s = append(data.TargetIPv6s, ipStr)
		}
		
		// Map PTR for all discovered IPs
		ptrs, _ := net.LookupAddr(ipStr)
		for _, ptr := range ptrs {
			data.ReverseDNS = append(data.ReverseDNS, fmt.Sprintf("%-15s → %s", ipStr, strings.TrimSuffix(ptr, ".")))
		}
	}

	// 2. GEO/ORG TELEMETRY (via Primary v4)
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, _ := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			data.Org, data.ISP, data.AS = g.Org, g.Isp, g.As
			data.City, data.Region, data.RegionName = g.City, g.Region, g.RegionName
			data.Country, data.CountryCode, data.Zip = g.Country, g.CountryCode, g.Zip
			data.Timezone, data.Lat, data.Lon = g.Timezone, g.Lat, g.Lon
			data.IsHosting = g.Hosting
			resp.Body.Close()
		}

		// 3. MULTI-PORT PROBE
		ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087}
		var wg sync.WaitGroup
		for _, p := range ports {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", data.TargetIPs[0], port), 1*time.Second)
				if err == nil {
					data.ScanResults = append(data.ScanResults, fmt.Sprintf("PORT %-4d: OPEN", port))
					conn.Close()
				}
			}(p)
		}
		wg.Wait()
	}

	// 4. CLUSTER RECON
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		host := strings.TrimSuffix(s.Host, ".")
		nips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range nips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[host] = ipStrings
	}

	mineLeaks(input, &data)
	return data, nil
}

func mineLeaks(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 4 * time.Second,
	}
	req, _ := http.NewRequest("GET", "https://"+target, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		data.WAFType = resp.Header.Get("Server")
		if data.WAFType != "" { data.IsWAF = true }
		
		// LEAK SYNC
		if id := resp.Header.Get("Ar-Request-Id"); id != "" { data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Node-ID ["+id+"]") }
		if ray := resp.Header.Get("CF-RAY"); ray != "" { data.ScanResults = append(data.ScanResults, "DEBUG: Cloudflare-Ray ["+ray+"]") }
		if edge := resp.Header.Get("X-Ar-Edge-Id"); edge != "" { data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Edge-Loc ["+edge+"]") }
	}
}
