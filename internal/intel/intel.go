package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
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

	// 1. Cluster Lookup
	ns, _ := net.LookupNS(input)
	for _, s := range ns {
		ips, _ := net.LookupIP(s.Host)
		var ipStrings []string
		for _, ip := range ips {
			ipStrings = append(ipStrings, ip.String())
		}
		data.NameServers[s.Host] = ipStrings
	}

	// 2. IP Resolution
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil {
			data.TargetIPs = append(data.TargetIPs, ip.String())
		}
	}

	// 3. Geo & Ping
	if len(data.TargetIPs) > 0 {
		geo, _ := fetchGeo(data.TargetIPs[0])
		data.Org, data.Lat, data.Lon = geo.Org, geo.Lat, geo.Lon
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	// 4. Port Surface Recon
	ports := []string{"80", "443", "8080", "2082", "2083", "2086", "2087"}
	for _, p := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(input, p), 400*time.Millisecond)
		if err == nil {
			data.ScanResults = append(data.ScanResults, "PORT "+p+": OPEN")
			conn.Close()
		}
	}

	analyzeWAF(input, &data)
	return data, nil
}

func fetchGeo(ip string) (models.GeoResponse, error) {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil { return models.GeoResponse{}, err }
	defer resp.Body.Close()
	var r models.GeoResponse
	json.NewDecoder(resp.Body).Decode(&r)
	return r, nil
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func analyzeWAF(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()
	srv := resp.Header.Get("Server")
	if strings.Contains(strings.ToLower(srv), "arvan") {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
	}
	data.ScanResults = append(data.ScanResults, "STACK: "+srv)
}
