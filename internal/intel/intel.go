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

	// 1. Authoritative Cluster Recon
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		host := strings.TrimSuffix(s.Host, ".")
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[host] = ipStrings
	}

	// 2. IP Mapping
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { data.TargetIPs = append(data.TargetIPs, ip.String()) }
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
				address := fmt.Sprintf("%s:%d", targetIP, port)
				conn, err := net.DialTimeout("tcp", address, 800*time.Millisecond)
				if err == nil {
					data.ScanResults = append(data.ScanResults, fmt.Sprintf("PORT %d: OPEN", port))
					conn.Close()
				}
			}(p)
		}
		wg.Wait()

		// 4. Deep Geo & ISP Info
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

	// 5. Header Mining (Software, Node-ID, Trace-ID)
	mineHeaders(input, &data)
	return data, nil
}

func mineHeaders(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get("https://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	// Capture Server/Software
	if srv := resp.Header.Get("Server"); srv != "" {
		data.IsWAF, data.WAFType = true, srv
	}
	
	// ArvanCloud Detection & Leak Extraction
	if arvanID := resp.Header.Get("Ar-Request-Id"); arvanID != "" {
		data.ScanResults = append(data.ScanResults, "DEBUG: Arvan-Node-ID ["+arvanID+"]")
	}
	if trace := resp.Header.Get("X-Trace-Id"); trace != "" {
		data.ScanResults = append(data.ScanResults, "DEBUG: Trace-ID Detected ["+trace+"]")
	}
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return time.Since(start).String()
}
