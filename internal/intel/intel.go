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

	// 1. DNS & PTR Recovery
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		host := strings.TrimSuffix(s.Host, ".")
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[host] = ipStrings
	}

	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { 
			ipStr := ip.String()
			data.TargetIPs = append(data.TargetIPs, ipStr) 
			ptrs, _ := net.LookupAddr(ipStr)
			for _, ptr := range ptrs {
				data.ReverseDNS = append(data.ReverseDNS, fmt.Sprintf("%s → %s", ipStr, strings.TrimSuffix(ptr, ".")))
			}
		}
	}

	// 2. Deep Public Data Scrape (DOX)
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, _ := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			data.Org, data.ISP, data.AS = g.Org, g.Isp, g.As
			data.City, data.Region, data.Country, data.Zip = g.City, g.RegionName, g.Country, g.Zip
			data.Lat, data.Lon = g.Lat, g.Lon
			data.IsMobile, data.IsProxy, data.IsHosting = g.Mobile, g.Proxy, g.Hosting
			resp.Body.Close()
		}
		
		// Port Scan
		ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087}
		var wg sync.WaitGroup
		for _, p := range ports {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", data.TargetIPs[0], port), 1*time.Second)
				if err == nil {
					data.ScanResults = append(data.ScanResults, fmt.Sprintf("PORT %d: OPEN", port))
					conn.Close()
				}
			}(p)
		}
		wg.Wait()
	}

	mineLeaks(input, &data)
	return data, nil
}

func mineLeaks(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get("https://" + target)
	if err == nil {
		defer resp.Body.Close()
		data.WAFType = resp.Header.Get("Server")
		if data.WAFType != "" { data.IsWAF = true }
		
		// ID LEAKS
		if id := resp.Header.Get("Ar-Request-Id"); id != "" { data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Node-ID ["+id+"]") }
		if cf := resp.Header.Get("CF-RAY"); cf != "" { data.ScanResults = append(data.ScanResults, "DEBUG: Cloudflare-Ray ["+cf+"]") }
	}
}
