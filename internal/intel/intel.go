package intel

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	data := models.IntelData{
		TargetName:  input,
		NameServers: make(map[string][]string),
	}

	// 1. DNS Cluster Recon & IP Lookup
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips { ipStrings = append(ipStrings, ip.String()) }
		data.NameServers[s.Host] = ipStrings
	}

	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil { 
			data.TargetIPs = append(data.TargetIPs, ip.String()) 
			names, _ := net.LookupAddr(ip.String())
			data.ReverseDNS = append(data.ReverseDNS, names...)
		}
	}

	// 2. Deep Geo & Phone Data Extraction
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 3 * time.Second}
		resp, _ := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=66846719")
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			var g models.GeoResponse
			json.Unmarshal(body, &g)
			
			// RESTORING DATA TO MODEL
			data.Org = g.Org
			data.City, data.Region, data.Country = g.City, g.RegionName, g.Country
			data.Lat, data.Lon = g.Lat, g.Lon
			data.IsMobile = g.Mobile
			data.IsProxy = g.Proxy
			data.IsHosting = g.Hosting
			data.RawGeo = string(body)
			
			resp.Body.Close()
		}
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	analyzeWAF(input, &data)
	return data, nil
}

// ... pingTarget and analyzeWAF functions remain the same ...
