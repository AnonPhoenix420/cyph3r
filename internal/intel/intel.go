package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Uses system's native VPN tunnel
func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	ips, _ := net.LookupIP(input)
	for _, ip := range ips { data.TargetIPs = append(data.TargetIPs, ip.String()) }
	
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		data.RawGeo = raw
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		addrs, _ := net.LookupHost(ns.Host)
		data.NameServers[ns.Host] = addrs
	}

	data.ScanResults = performTacticalScan(input)
	return data, nil
}

func fetchGeo(ip string) (GeoResponse, string) {
	client := GetClient()
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil { return GeoResponse{Org: "SECURE_NODE"}, "{}" }
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var r GeoResponse
	json.Unmarshal(body, &r)
	return r, string(body)
}

type GeoResponse struct {
	Country, City, Isp, Org string
	Lat, Lon float64
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 1200*time.Millisecond)
	if err != nil { return "LOST" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var results []string
	var serverHeader string
	client := GetClient()
	ports := []int{80, 443, 8080}
	for _, p := range ports {
		addr := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
		if err == nil {
			conn.Close()
			if serverHeader == "" {
				proto := "http"; if p == 443 { proto = "https" }
				if hResp, hErr := client.Get(fmt.Sprintf("%s://%s", proto, target)); hErr == nil {
					serverHeader = hResp.Header.Get("Server")
					hResp.Body.Close()
				}
			}
			results = append(results, fmt.Sprintf("PORT %d: OPEN", p))
		}
	}
	if serverHeader == "" { serverHeader = "HIDDEN" }
	results = append(results, "STACK: "+serverHeader)
	return results
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	clean := strings.TrimPrefix(number, "+")
	d := models.PhoneData{Number: number, Risk: "LOW", SocialPresence: []string{"WhatsApp", "Telegram"}}
	if strings.HasPrefix(clean, "98") {
		d.Country, d.Carrier = "Iran", "MCI/Irancell"
	} else if strings.HasPrefix(clean, "1") {
		d.Country, d.Carrier = "USA/Canada", "North American Band"
	}
	d.HandleHint = "uid_" + clean[len(clean)-6:]
	d.MapLink = "http://googleusercontent.com/maps.google.com/search?q=" + d.Country
	return d, nil
}
