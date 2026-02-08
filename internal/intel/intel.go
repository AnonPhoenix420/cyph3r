package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/likexian/whois"
)

// GetFullIntel aggregates DNS, WHOIS, and Geo-IP data into a single model.
func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Resolve Network Identity (via internal/intel/dns.go)
	// Fetches Neon Blue IPs and Neon Yellow Nameservers
	data.IPs, data.Nameservers = LookupNodes(target)

	// 2. WHOIS Intelligence (Registry Data)
	// Fetches registrar, expiration, and ownership info
	rawWhois, err := whois.Whois(target)
	if err == nil {
		data.WhoisRaw = rawWhois
		// Simple extraction for the HUD
		data.Registrar = extractField(rawWhois, "Registrar:")
	}

	// 3. Geographic & ISP Intelligence
	// Using a dedicated client with a 5-second timeout to prevent hangs
	client := http.Client{Timeout: time.Second * 5}
	
	// Querying ip-api.com for the detailed Geo-HUD data
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,countryCode,regionName,city,lat,lon,isp,org,as", target)
	
	resp, err := client.Get(apiURL)
	if err != nil {
		return data, fmt.Errorf("GEO_API_UNREACHABLE")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	var apiRes struct {
		Status      string  `json:"status"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"regionName"`
		City        string  `json:"city"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		ISP         string  `json:"isp"`
		Org         string  `json:"org"`
		ASN         string  `json:"as"`
	}

	if err := json.Unmarshal(body, &apiRes); err == nil && apiRes.Status == "success" {
		data.Country = apiRes.Country
		data.CountryCode = apiRes.CountryCode
		data.Region = apiRes.Region
		data.City = apiRes.City
		data.Lat = apiRes.Lat
		data.Lon = apiRes.Lon
		data.ISP = apiRes.ISP
		data.Org = apiRes.Org
		data.ASN = apiRes.ASN
	}

	return data, nil
}

// Helper function to pull specific strings out of the messy WHOIS block
func extractField(raw string, field string) string {
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(field)) {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "UNKNOWN"
}
