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
	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// 1. DNS Resolution
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil {
			data.TargetIPs = append(data.TargetIPs, ip.String())
		}
	}

	// 2. GEO & Latency (Fixed to populate HUD)
	if len(data.TargetIPs) > 0 {
		geo, _ := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	analyzeExploitSurface(input, &data)
	return data, nil
}

func fetchGeo(ip string) (models.GeoResponse, error) {
	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return models.GeoResponse{}, err
	}
	defer resp.Body.Close()
	
	var r models.GeoResponse
	json.NewDecoder(resp.Body).Decode(&r)
	return r, nil
}

func pingTarget(ip string) string {
	start := time.Now()
	// Dialing port 443 to measure real-world response time
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1200*time.Millisecond)
	if err != nil {
		return "TIMEOUT"
	}
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func analyzeExploitSurface(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://" + target)
	if err != nil { return }
	defer resp.Body.Close()

	srv := resp.Header.Get("Server")
	if strings.Contains(strings.ToLower(srv), "arvan") || resp.Header.Get("ArvanCloud-Trace") != "" {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
	}
	data.ScanResults = append(data.ScanResults, "STACK: "+srv)
}
