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

	// 1. DNS Cluster Recon
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		host := strings.TrimSuffix(s.Host, ".")
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[host] = ipStrings
	}

	// 2. IP Mapping & REVERSE DNS (PTR) Recovery
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { 
			ipStr := ip.String()
			data.TargetIPs = append(data.TargetIPs, ipStr) 
			
			// Recover Reverse DNS Strings
			ptrs, _ := net.LookupAddr(ipStr)
			for _, ptr := range ptrs {
				data.ReverseDNS = append(data.ReverseDNS, fmt.Sprintf("%s → %s", ipStr, strings.TrimSuffix(ptr, ".")))
			}
		}
	}

	if len(data.TargetIPs) > 0 {
		targetIP := data.TargetIPs[0]
		
		// 3. Port Scan Logic
		ports := []int{80, 443, 8080, 2082, 2083, 2086, 2087}
		var wg sync.WaitGroup
		for _, p := range ports {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", targetIP, port), 1*time.Second)
				if err == nil {
					data.ScanResults = append(data.ScanResults, fmt.Sprintf("PORT %d: OPEN", port))
					conn.Close()
				}
			}(p)
		}
		wg.Wait()

		// 4. Deep Telemetry Fetch
		client := &http.Client{Timeout: 5 * time.Second}
		resp, _ := client.Get("http://ip-api.com/json/" + targetIP + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			data.Org, data.City, data.Region, data.Country = g.Org, g.City, g.RegionName, g.Country
			data.Lat, data.Lon = g.Lat, g.Lon
			data.IsMobile, data.IsProxy, data.IsHosting = g.Mobile, g.Proxy, g.Hosting
			resp.Body.Close()
		}
		data.Latency = pingTarget(targetIP)
	}

	mineLeaks(input, &data)
	return data, nil
}

// ... mineLeaks and pingTarget remain the same ...
