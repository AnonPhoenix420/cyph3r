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
	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// DNS Cluster Recon
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		host := strings.TrimSuffix(s.Host, ".")
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[host] = ipStrings
	}

	// IP Mapping
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { data.TargetIPs = append(data.TargetIPs, ip.String()) }
	}

	// Deep Geo Fetch
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, _ := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			
			data.Org = g.Org
			data.City, data.Region, data.Country = g.City, g.RegionName, g.Country
			data.Lat, data.Lon = g.Lat, g.Lon
			data.IsMobile, data.IsProxy, data.IsHosting = g.Mobile, g.Proxy, g.Hosting
			data.RawGeo = string(body)
			resp.Body.Close()
		}
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	analyzeWAF(input, &data)
	return data, nil
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1*time.Second)
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
	if err == nil {
		defer resp.Body.Close()
		if srv := resp.Header.Get("Server"); srv != "" { 
			data.IsWAF, data.WAFType = true, srv 
		}
	}
}
