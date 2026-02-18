package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// GetClient routes through system's active VPN and ignores invalid certs
func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

// isVPNActive checks for ProtonVPN interfaces to prevent leaks
func isVPNActive() bool {
	data, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return false
	}
	content := string(data)
	return strings.Contains(content, "tun") || 
	       strings.Contains(content, "proton") || 
	       strings.Contains(content, "wg")
}

// scrub redacts local identity data from output strings
func scrub(input string) string {
	output := input
	hostname, _ := os.Hostname()
	user := os.Getenv("USER")
	output = strings.ReplaceAll(output, hostname, "TARGET_NODE")
	if user != "" {
		output = strings.ReplaceAll(output, user, "operator")
	}
	output = strings.ReplaceAll(output, "/data/data/com.termux/files/home", "~")
	return output
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// FAIL-SAFE: Check VPN status before any networking
	if !isVPNActive() {
		fmt.Println("\n\033[31m[!] PROTON VPN NOT DETECTED. KILLING PROCESS.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}
	
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		data.RawGeo = raw // This is already scrubbed in fetchGeo
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

func fetchGeo(ip string) (models.GeoResponse, string) {
	client := GetClient()
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return models.GeoResponse{Org: "UPLINK_ENCRYPTED"}, "{}"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r models.GeoResponse
	json.Unmarshal(body, &r)

	var anyData interface{}
	json.Unmarshal(body, &anyData)
	prettyJSON, err := json.MarshalIndent(anyData, "", "  ")
	if err != nil {
		return r, scrub(string(body))
	}

	return r, scrub(string(prettyJSON))
}

func pingTarget(ip string) string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "443"), 2*time.Second)
	if err != nil {
		return "TIMEOUT"
	}
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
		conn, err := net.DialTimeout("tcp", addr, 1500*time.Millisecond)
		if err == nil {
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
	d := models.PhoneData{
		Number:         number,
		Risk:           "LOW",
		SocialPresence: []string{"WhatsApp", "Telegram", "Signal"},
	}
	
	if strings.HasPrefix(clean, "98") {
		d.Country, d.Carrier = "Iran", "MCI/Irancell"
	} else if strings.HasPrefix(clean, "1") {
		d.Country, d.Carrier = "USA/Canada", "Global Mobile"
	} else {
		d.Country, d.Carrier = "Unknown", "Provider Restricted"
	}
	
	d.HandleHint = "uid_" + clean[len(clean)-6:]
	d.MapLink = "https://www.google.com/maps/search/" + d.Country
	return d, nil
}
