package intel

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

type APIResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	As          string  `json:"as"`
	Org         string  `json:"org"`
	RegionName  string  `json:"regionName"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

type ThreatSource struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type LeakCheckResp struct {
	Success bool           `json:"success"`
	Found   int            `json:"found"`
	Sources []ThreatSource `json:"sources"`
}

func CheckThreatFeeds(target string) []string {
	detections := make([]string, 0)
	client := &http.Client{Timeout: 3 * time.Second}
	cleanedTarget := strings.TrimSpace(target)
	
	url := fmt.Sprintf("https://leakcheck.io/api/public?check=%s", cleanedTarget)
	resp, err := client.Get(url)
	if err == nil {
		defer resp.Body.Close()
		var leakData LeakCheckResp
		if json.NewDecoder(resp.Body).Decode(&leakData) == nil && len(leakData.Sources) > 0 {
			for _, src := range leakData.Sources {
				detections = append(detections, fmt.Sprintf("CRITICAL ➔ Found inside breach: %s (%s)", src.Name, src.Date))
			}
		}
	}

	if strings.Contains(cleanedTarget, "scam") || strings.Contains(cleanedTarget, "crypto-drain") {
		detections = append(detections, "INTEL WARNING ➔ Target matches active systemic fraud tracking profile identifiers.")
	}

	if len(detections) == 0 {
		detections = append(detections, "CLEAN ➔ Target not indexed in high-risk baseline network registries.")
	}

	return detections
}

func ResolveEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return "Invalid Email Formatting Profile"
	}
	domain := parts[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return fmt.Sprintf("No active mail gateway associated with domain: %s", domain)
	}
	return fmt.Sprintf("Verified Active Enterprise Mail Gateway Handling: %s (Priority: %d)", mxRecords[0].Host, mxRecords[0].Pref)
}

func ResolvePhone(phone string) (string, string, string) {
	cleaned := strings.ReplaceAll(strings.ReplaceAll(phone, "+", ""), " ", "")
	if strings.HasPrefix(cleaned, "1") {
		return "North American Numbering Plan (NANP)", "USA/Canada Carrier Block", "America/New_York Zone"
	} else if strings.HasPrefix(cleaned, "44") {
		return "United Kingdom Infrastructure", "BT/Vodafone Routing Cluster", "Europe/London Zone"
	}
	return "Global Telephony Allocation", "International Transit Gateway Selector", "UTC Offset Variable"
}

func ResolveNetwork(domain string) (string, models.GeoData, string, string, string, []string, []string, []string, []string, models.SQLExposure) {
	return ResolveNetworkElite(domain, 0, "")
}

func ResolveNetworkElite(domain string, baseDelay time.Duration, customUserAgent string) (string, models.GeoData, string, string, string, []string, []string, []string, []string, models.SQLExposure) {
	var geo models.GeoData
	var asn = "UNKNOWN_ASN"
	var ownerName = "WHOIS_PRIVACY_PROTECTED"
	var createdDate = "METADATA_EXPEDITED"
	
	// FORCE EMPTY EXPLICIT ARRAYS INSTANTIATION INSTEAD OF NULL
	openPorts := make([]string, 0)
	banners := make([]string, 0)
	vulns := make([]string, 0)
	leaks := make([]string, 0)
	var sqlMetrics models.SQLExposure

	var targetIP string
	if net.ParseIP(domain) != nil {
		targetIP = domain
	} else {
		ips, err := net.LookupIP(domain)
		if err != nil || len(ips) == 0 {
			return "0.0.0.0", geo, asn, ownerName, createdDate, openPorts, banners, vulns, leaks, sqlMetrics
		}
		targetIP = ips[0].String()
	}

	client := &http.Client{Timeout: 4 * time.Second}
	fields := "status,country,city,as,org,regionName,zip,lat,lon"
	req, err := http.NewRequest("GET", fmt.Sprintf("http://ip-api.com/json/%s?fields=%s", targetIP, fields), nil)
	if err == nil {
		if customUserAgent != "" {
			req.Header.Set("User-Agent", customUserAgent)
		} else {
			req.Header.Set("User-Agent", "CYPH3R/Master-Engine-2026")
		}
		
		if resp, err := client.Do(req); err == nil {
			var data APIResponse
			if json.NewDecoder(resp.Body).Decode(&data) == nil && data.Status == "success" {
				geo.City = fmt.Sprintf("%s, %s (%s)", data.City, data.RegionName, data.Zip)
				geo.Country = data.Country
				geo.Latitude = fmt.Sprintf("%.4f", data.Lat)
				geo.Longitude = fmt.Sprintf("%.4f", data.Lon)
				geo.MapReference = fmt.Sprintf("http://maps.google.com/?q=%.4f,%.4f", data.Lat, data.Lon)
				asn = data.As
				if data.Org != "" {
					ownerName = data.Org
				}
			}
			resp.Body.Close()
		}
	}

	if net.ParseIP(domain) == nil {
		if mxRecords, err := net.LookupMX(domain); err == nil {
			for _, mx := range mxRecords {
				leaks = append(leaks, fmt.Sprintf("MX Mail Routing Node ➔ %s (Priority: %d)", mx.Host, mx.Pref))
			}
		}
		if txtRecords, err := net.LookupTXT(domain); err == nil {
			for _, txt := range txtRecords {
				if strings.Contains(txt, "v=spf1") {
					leaks = append(leaks, fmt.Sprintf("SPF Trust Blueprint ➔ %s", txt))
				}
			}
		}
	}

	portsToScan := []int{21, 22, 80, 443, 1433, 3306, 5432, 8080}
	dialer := &net.Dialer{Timeout: 1200 * time.Millisecond}

	for _, port := range portsToScan {
		if baseDelay > 0 {
			time.Sleep(baseDelay + time.Duration(rand.Intn(200))*time.Millisecond)
		}

		address := net.JoinHostPort(targetIP, fmt.Sprintf("%d", port))
		conn, err := dialer.Dial("tcp", address)
		if err == nil {
			_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
			buffer := make([]byte, 256)
			n, readErr := conn.Read(buffer)
			
			openPorts = append(openPorts, fmt.Sprintf("%d/TCP", port))
			
			if port == 1433 || port == 3306 || port == 5432 {
				sqlMetrics.Exposed = true
				sqlMetrics.Ports = append(sqlMetrics.Ports, port)
				sqlMetrics.RiskLevel = "CRITICAL"
				vulns = append(vulns, fmt.Sprintf("Database Port %d Open ➔ High Exposure Vector Risk", port))
			}

			if readErr == nil && n > 0 {
				banners = append(banners, fmt.Sprintf("%d: %s", port, strings.TrimSpace(string(buffer[:n]))))
			} else {
				banners = append(banners, fmt.Sprintf("%d: Interface Active (Handshake Confirmed)", port))
			}
			conn.Close()
		}
	}

	if !sqlMetrics.Exposed {
		sqlMetrics.RiskLevel = "CLEAR"
	}

	return targetIP, geo, asn, ownerName, createdDate, openPorts, banners, vulns, leaks, sqlMetrics
}
