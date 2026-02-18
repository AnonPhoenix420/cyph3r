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

// GetClient routes through system's active VPN and ignores invalid certs
func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

// isVPNActive uses Triple-Check validation for Termux/Parrot environments
func isVPNActive() bool {
	// Check 1: Virtual Interface Check
	data, _ := os.ReadFile("/proc/net/dev")
	content := string(data)
	if strings.Contains(content, "tun") || strings.Contains(content, "wg") || strings.Contains(content, "proton") {
		return true
	}

	// Check 2: Routing Table Check (Bypasses PRoot limitations)
	out, _ := exec.Command("sh", "-c", "ip route | grep -E 'tun|wg|proton'").Output()
	if len(out) > 0 {
		return true
	}

	// Check 3: ISP/Organization Identity Check (Targets Datacamp/Proton)
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		sBody := strings.ToLower(string(body))
		// Validates against your specific connection (Datacamp Limited)
		if strings.Contains(sBody, "datacamp") || strings.Contains(sBody, "proton") || strings.Contains(sBody, "m247") {
			return true
		}
	}
	return false
}

// scrub redacts local identity data from output
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
	d.MapLink = "http://maps.google.com/?q=" + d.Country
	return d, nil
}
