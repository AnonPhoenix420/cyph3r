package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

func isVPNActive() bool {
	data, _ := os.ReadFile("/proc/net/dev")
	content := string(data)
	if strings.Contains(content, "tun") || strings.Contains(content, "wg") || strings.Contains(content, "proton") {
		return true
	}
	out, _ := exec.Command("sh", "-c", "ip route | grep -E 'tun|wg|proton'").Output()
	if len(out) > 0 { return true }
	
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/?fields=isp")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		sBody := strings.ToLower(string(body))
		if strings.Contains(sBody, "datacamp") || strings.Contains(sBody, "proton") || strings.Contains(sBody, "m247") {
			return true
		}
	}
	return false
}

func scrub(input string) string {
	output := input
	hostname, _ := os.Hostname()
	user := os.Getenv("USER")
	output = strings.ReplaceAll(output, hostname, "TARGET_NODE")
	if user != "" { output = strings.ReplaceAll(output, user, "operator") }
	output = strings.ReplaceAll(output, "/data/data/com.termux/files/home", "~")
	return output
}

func GetTargetIntel(input string) (models.IntelData, error) {
	if !isVPNActive() {
		fmt.Println("\n\033[31m[!] PROTON VPN NOT DETECTED. KILLING PROCESS.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
		// NEW: Reverse DNS (PTR) Lookup - Tells you the server's true name
		names, _ := net.LookupAddr(ip.String())
		for _, name := range names {
			data.ReverseDNS = append(data.ReverseDNS, strings.TrimSuffix(name, "."))
		}
	}
	
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		
		// NEW: Intelligence Labeling
		usage := "RESIDENTIAL"
		if geo.Hosting { usage = "DATA_CENTER" }
		if geo.Mobile { usage = "MOBILE_NET" }
		if geo.Proxy { usage += "/PROXY_DETECTED" }
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
		
		data.RawGeo = raw
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		addrs, _ := net.LookupHost(ns.Host)
		data.NameServers[ns.Host] = addrs
	}

	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)
	return data, nil
}

func fetchGeo(ip string) (models.GeoResponse, string) {
	client := GetClient()
	// Requesting full security fields: mobile, proxy, hosting
	resp, err := client.Get("http://ip-api.com/json/" + ip + "?fields=status,message,country,countryCode,regionName,city,zip,lat,lon,timezone,isp,org,as,mobile,proxy,hosting,query")
	if err != nil { return models.GeoResponse{Org: "OFFLINE"}, "{}" }
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var r models.GeoResponse
	json.Unmarshal(body, &r)

	var anyData interface{}
	json.Unmarshal(body, &anyData)
	prettyJSON, _ := json.MarshalIndent(anyData, "", "  ")

	return r, scrub(string(prettyJSON))
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil { return "TIMEOUT" }
	defer conn.Close()
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds())
}

func performTacticalScan(target string) []string {
	var results []string
	var serverHeader string
	client := GetClient()
	ports := []int{80, 443, 8080}
	for _, p := range ports {
		addr := net.JoinHostPort(target, fmt.Sprintf("%d", p))
		if conn, err := net.DialTimeout("tcp", addr, 1500*time.Millisecond); err == nil {
			conn.Close()
			results = append(results, fmt.Sprintf("PORT %d: OPEN", p))
			if (p == 80 || p == 443) && serverHeader == "" {
				proto := "http"
				if p == 443 { proto = "https" }
				if hResp, hErr := client.Get(fmt.Sprintf("%s://%s", proto, target)); hErr == nil {
					serverHeader = hResp.Header.Get("Server")
					hResp.Body.Close()
				}
			}
		}
	}
	if serverHeader == "" { serverHeader = "F_STACK_HIDDEN" }
	results = append(results, "STACK: "+serverHeader)
	return results
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	clean := strings.TrimPrefix(number, "+")
	d := models.PhoneData{Number: number, Risk: "LOW", SocialPresence: []string{"WhatsApp", "Telegram"}}
	if strings.HasPrefix(clean, "98") { d.Country, d.Carrier = "Iran", "MCI/Irancell" }
	d.HandleHint = "uid_" + clean[len(clean)-6:]
	d.MapLink = "http://google.com/maps?q=" + d.Country
	return d, nil
}
