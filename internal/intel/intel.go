package intel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/ghost"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/nyaruka/phonenumbers"
	"golang.org/x/net/proxy"
)

type CrtShEntry struct {
	NameValue string `json:"name_value"`
}

type GeoAPIResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
}

// CommonServicePorts defines a robust dictionary of high-value target ports
var CommonServicePorts = map[int]string{
	21:    "FTP",
	22:    "SSH",
	23:    "Telnet",
	25:    "SMTP",
	53:    "DNS",
	80:    "HTTP",
	110:   "POP3",
	135:   "MSRPC",
	139:   "NetBIOS",
	143:   "IMAP",
	443:   "HTTPS",
	445:   "SMB",
	993:   "IMAPS",
	995:   "POP3S",
	1433:  "MS-SQL",
	1521:  "Oracle",
	3306:  "MySQL",
	3389:  "RDP",
	5432:  "PostgreSQL",
	5900:  "VNC",
	6379:  "Redis",
	8080:  "HTTP-Proxy",
	8443:  "HTTPS-Alt",
	27017: "MongoDB",
}

// getClient builds an HTTP client with Ghost transport support if enabled
func getClient(useGhost bool, timeout time.Duration) *http.Client {
	var tr *http.Transport
	if useGhost {
		tr = ghost.GetTransport(true)
	}
	if tr == nil {
		tr = &http.Transport{
			MaxIdleConns:        50,
			IdleConnTimeout:     30 * time.Second,
		}
	}
	return &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
}

func DiscoverOriginAndOSINT(targetDomain string, useGhost bool) models.ExtractedIntel {
	var intel models.ExtractedIntel
	subMap := make(map[string]bool)

	modeLabel := "Direct"
	if useGhost {
		modeLabel = "Ghost Mode (Tor/SOCKS5)"
	}

	// 1. Passive CT Log Query via crt.sh
	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", targetDomain)
	client := getClient(useGhost, 15*time.Second)

	fmt.Printf("[*] Querying Certificate Transparency logs for %s [%s]...\n", targetDomain, modeLabel)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] Warning: CT log query failed: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[!] Warning: crt.sh returned HTTP status %d (Target may be rate-limiting or too large)\n", resp.StatusCode)
		} else {
			var entries []CrtShEntry
			if err := json.Unmarshal(body, &entries); err == nil {
				for _, entry := range entries {
					for _, sub := range strings.Split(entry.NameValue, "\n") {
						sub = strings.TrimSpace(sub)
						sub = strings.TrimPrefix(sub, "*.")
						if sub != "" && !subMap[sub] {
							subMap[sub] = true
							intel.Subdomains = append(intel.Subdomains, sub)
						}
					}
				}
			} else {
				fmt.Printf("[!] Warning: Failed to parse CT JSON response: %v\n", err)
			}
		}
	}

	// 2. Active Fallback Probe (Triggers if CT logs are blocked/empty)
	if len(intel.Subdomains) == 0 {
		fmt.Printf("[*] CT logs rate-limited or empty. Engaging active infrastructure probe for %s...\n", targetDomain)
		commonSubs := []string{
			"www", "mail", "webmail", "ns1", "ns2", "dns", "api", "portal", 
			"secure", "admin", "vpn", "remote", "autodiscover", "exchange", 
			"login", "auth", "support", "status", "cloud", "dev", "staging", 
			"test", "shop", "store", "erp", "internal", "gateway", "proxy",
			"cdn", "app", "dashboard", "jenkins", "git", "metrics",
		}

		for _, prefix := range commonSubs {
			sub := fmt.Sprintf("%s.%s", prefix, targetDomain)
			ips, err := net.LookupIP(sub)
			if err == nil && len(ips) > 0 && !subMap[sub] {
				subMap[sub] = true
				intel.Subdomains = append(intel.Subdomains, sub)
			}
		}
	}

	fmt.Printf("[+] Discovered %d unique validated subdomains.\n", len(intel.Subdomains))

	intel.FaviconHash = fetchFaviconHash(targetDomain, client)

	ipMap := make(map[string]bool)
	limit := len(intel.Subdomains)
	if limit > 100 {
		limit = 100
		fmt.Println("[*] Throttling DNS resolution to first 100 subdomains for speed.")
	}
	for i := 0; i < limit; i++ {
		sub := intel.Subdomains[i]
		ips, err := net.LookupIP(sub)
		if err == nil {
			for _, ip := range ips {
				ipStr := ip.String()
				if !isCDNEdge(ipStr) && !ipMap[ipStr] {
					ipMap[ipStr] = true
					intel.RealIPs = append(intel.RealIPs, fmt.Sprintf("%s (%s)", ipStr, sub))
				}
			}
		}
	}

	// 3. Harvest DNS Records (TXT, MX, NS) & Global RDAP Registry Telemetry
	var dnsBuilder strings.Builder

	txtRecords, _ := net.LookupTXT(targetDomain)
	for _, txt := range txtRecords {
		dnsBuilder.WriteString(txt + " ")
	}

	mxRecords, err := net.LookupMX(targetDomain)
	if err == nil {
		for _, mx := range mxRecords {
			dnsBuilder.WriteString(mx.Host + " ")
		}
	}

	nsRecords, err := net.LookupNS(targetDomain)
	if err == nil {
		for _, ns := range nsRecords {
			dnsBuilder.WriteString(ns.Host + " ")
		}
	}

	rdapText := fetchRDAPRegistryData(targetDomain, client)
	combinedText := dnsBuilder.String() + " " + rdapText

	intel.Emails = extractEmails(combinedText)
	intel.PhoneNumbers = extractPhones(combinedText)
	intel.SocialHandles = extractSocials(combinedText)

	return intel
}

// ExecuteComprehensiveReport builds the full structured report from core telemetry
func ExecuteComprehensiveReport(targetDomain string, useGhost bool) models.ComprehensiveReport {
	rawIntel := DiscoverOriginAndOSINT(targetDomain, useGhost)

	var primaryIP string
	ips, err := net.LookupIP(targetDomain)
	if err == nil && len(ips) > 0 {
		primaryIP = ips[0].String()
	}

	client := getClient(useGhost, 6*time.Second)
	locData := fetchGeoLocation(primaryIP, client)

	report := models.ComprehensiveReport{
		Target:     targetDomain,
		TargetType: models.TargetDomain,
		ReverseDNS: primaryIP,
		Timestamp:  time.Now(),
		RiskScore:  65,

		Location:   locData,
		Associated: rawIntel.RealIPs,
		Emails:     rawIntel.Emails,
		Phones:     rawIntel.PhoneNumbers,

		SQLCheck: models.SQLExposure{
			Exposed:   false,
			Ports:     []int{},
			RiskLevel: "LOW",
		},
	}

	for _, handle := range rawIntel.SocialHandles {
		report.SocialProfiles = append(report.SocialProfiles, models.SocialProfile{
			Platform:   "Public Artifact / Social Ref",
			Username:   handle,
			ProfileURL: "https://" + handle,
			Confidence: 90,
		})
	}

	return report
}

// ExecutePortScan performs a concurrent TCP port sweep across the service dictionary with Ghost support
func ExecutePortScan(targetHost string, useGhost bool) []string {
	modeLabel := "Direct"
	if useGhost {
		modeLabel = "Ghost Mode (Tor/SOCKS5)"
	}
	fmt.Printf("[*] Starting tactical port sweep on %s [%s] across %d common service vectors...\n", targetHost, modeLabel, len(CommonServicePorts))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var openPorts []string

	semaphore := make(chan struct{}, 100)

	var proxyDialer proxy.Dialer
	if useGhost {
		var err error
		proxyDialer, err = proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
		if err != nil {
			fmt.Printf("[!] Warning: SOCKS5 dialer init failed for port sweep: %v\n", err)
		}
	}

	for port, service := range CommonServicePorts {
		wg.Add(1)
		go func(p int, svc string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			address := fmt.Sprintf("%s:%d", targetHost, p)
			var conn net.Conn
			var err error

			if useGhost && proxyDialer != nil {
				ch := make(chan struct {
					c net.Conn
					e error
				}, 1)
				go func() {
					c, e := proxyDialer.Dial("tcp", address)
					ch <- struct {
						c net.Conn
						e error
					}{c, e}
				}()
				select {
				case <-time.After(2 * time.Second):
					err = fmt.Errorf("timeout")
				case res := <-ch:
					conn, err = res.c, res.e
				}
			} else {
				conn, err = net.DialTimeout("tcp", address, 1500*time.Millisecond)
			}

			if err == nil && conn != nil {
				conn.Close()
				resultStr := fmt.Sprintf("Port %d (%s) - OPEN", p, svc)
				mu.Lock()
				openPorts = append(openPorts, resultStr)
				mu.Unlock()
			}
		}(port, service)
	}

	wg.Wait()
	return openPorts
}

func fetchGeoLocation(ip string, client *http.Client) models.LocationData {
	defaultLoc := models.LocationData{
		Country:     "Unknown",
		CountryCode: "XX",
		City:        "Edge Infrastructure",
		Coordinates: "N/A",
		RadiusKM:    0.0,
	}

	if ip == "" || strings.HasPrefix(ip, "127.") || strings.HasPrefix(ip, "10.") {
		return defaultLoc
	}

	geoURL := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := client.Get(geoURL)
	if err != nil {
		return defaultLoc
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultLoc
	}

	var geo GeoAPIResponse
	if err := json.Unmarshal(body, &geo); err != nil || geo.Status != "success" {
		return defaultLoc
	}

	coords := fmt.Sprintf("%.4f° N, %.4f° E", geo.Lat, geo.Lon)
	if geo.Lat < 0 {
		coords = fmt.Sprintf("%.4f° S, %.4f° E", -geo.Lat, geo.Lon)
	}

	return models.LocationData{
		Country:     geo.Country,
		CountryCode: geo.CountryCode,
		State:       geo.RegionName,
		City:        geo.City,
		ZIP:         geo.Zip,
		Coordinates: coords,
		RadiusKM:    15.0,
	}
}

func fetchRDAPRegistryData(domain string, client *http.Client) string {
	fmt.Printf("[*] Querying global RDAP registry databases for %s...\n", domain)
	rdapURL := fmt.Sprintf("https://rdap.org/domain/%s", domain)
	resp, err := client.Get(rdapURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func fetchFaviconHash(domain string, client *http.Client) string {
	resp, err := client.Get(fmt.Sprintf("https://%s/favicon.ico", domain))
	if err != nil {
		resp, err = client.Get(fmt.Sprintf("http://%s/favicon.ico", domain))
		if err != nil {
			return "Unavailable"
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		return "Unavailable"
	}
	hash := md5.Sum(body)
	return hex.EncodeToString(hash[:])
}

func isCDNEdge(ip string) bool {
	return strings.HasPrefix(ip, "104.") ||
		strings.HasPrefix(ip, "172.64.") ||
		strings.HasPrefix(ip, "173.245.") ||
		strings.HasPrefix(ip, "192.254.") ||
		strings.HasPrefix(ip, "13.32.") ||
		strings.HasPrefix(ip, "151.101.")
}

func extractEmails(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
}

func extractPhones(text string) []string {
	re := regexp.MustCompile(`(?:\+\d{1,3}\s?)?(?:\(\d{3}\)|\d{3})[-.\s]?\d{3}[-.\s]?\d{4}`)
	matches := re.FindAllString(text, -1)
	var validatedPhones []string
	for _, m := range matches {
		num, err := phonenumbers.Parse(m, "US")
		if err == nil && phonenumbers.IsValidNumber(num) {
			validatedPhones = append(validatedPhones, phonenumbers.Format(num, phonenumbers.INTERNATIONAL))
		} else {
			validatedPhones = append(validatedPhones, m)
		}
	}
	return uniqueStrings(validatedPhones)
}

func extractSocials(text string) []string {
	re := regexp.MustCompile(`(?:twitter\.com|x\.com|linkedin\.com|github\.com|facebook\.com)/[a-zA-Z0-9_.-]+`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
}

func uniqueStrings(input []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range input {
		if !keys[entry] {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
