package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/likexian/whois-go"
	"github.com/nyaruka/phonenumbers"
)

// NodeIntel represents the high-density target profile
type NodeIntel struct {
	IPs       []string
	NS        []string
	Registrar string
	BornOn    string
	ISP       string
	Org       string
	Country   string
	City      string
	Zip       string
	Lat       float64
	Lon       float64
	Coords    string
	Maps      string
}

// GetFullIntel is the primary engine for deep-dive reconnaissance
func GetFullIntel(target string) (*NodeIntel, error) {
	data := &NodeIntel{}

	// 1. DNS Resolution (IPs & Name Servers)
	ips, _ := net.LookupIP(target)
	for _, ip := range ips {
		data.IPs = append(data.IPs, ip.String())
	}

	nss, _ := net.LookupNS(target)
	for _, ns := range nss {
		data.NS = append(data.NS, strings.TrimSuffix(ns.Host, "."))
	}

	// 2. WHOIS Data Extraction
	rawWhois, err := whois.Whois(target)
	if err == nil {
		lines := strings.Split(rawWhois, "\n")
		for _, line := range lines {
			lower := strings.ToLower(line)
			if strings.Contains(lower, "registrar:") && data.Registrar == "" {
				parts := strings.Split(line, ":")
				if len(parts) > 1 { data.Registrar = strings.TrimSpace(parts[1]) }
			}
			if (strings.Contains(lower, "creation date:") || strings.Contains(lower, "created:")) && data.BornOn == "" {
				parts := strings.Split(line, ":")
				if len(parts) > 1 { data.BornOn = strings.TrimSpace(strings.Join(parts[1:], ":")) }
			}
		}
	}

	// 3. Geo-IP & ISP Intelligence (Using ip-api.com)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,zip,lat,lon,isp,org", target))
	if err == nil {
		defer resp.Body.Close()
		var geo struct {
			Status  string  `json:"status"`
			Country string  `json:"country"`
			City    string  `json:"city"`
			Zip     string  `json:"zip"`
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			ISP     string  `json:"isp"`
			Org     string  `json:"org"`
		}
		if json.NewDecoder(resp.Body).Decode(&geo) == nil && geo.Status == "success" {
			data.Country = geo.Country
			data.City = geo.City
			data.Zip = geo.Zip
			data.Lat = geo.Lat
			data.Lon = geo.Lon
			data.ISP = geo.ISP
			data.Org = geo.Org
			data.Coords = fmt.Sprintf("%.4f, %.4f", geo.Lat, geo.Lon)
			data.Maps = fmt.Sprintf("https://www.google.com/maps?q=%.4f,%.4f", geo.Lat, geo.Lon)
		}
	}

	// Fallbacks for missing data
	if data.Registrar == "" { data.Registrar = "Private/Hidden" }
	if data.BornOn == "" { data.BornOn = "Unknown" }
	if data.ISP == "" { data.ISP = "Unknown ISP" }

	return data, nil
}

// PhoneLookup decrypts metadata for international phone numbers
func PhoneLookup(number string) string {
	parsed, err := phonenumbers.Parse(number, "")
	if err != nil {
		return fmt.Sprintf("Error: Invalid Number Format [%s]", number)
	}

	valid := "INVALID"
	if phonenumbers.IsValidNumber(parsed) {
		valid = "VALID"
	}

	region := phonenumbers.GetRegionCodeForNumber(parsed)
	return fmt.Sprintf("STATUS: %s | REGION: %s | TYPE: %v", valid, region, phonenumbers.GetNumberType(parsed))
}
