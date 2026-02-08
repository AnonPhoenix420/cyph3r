package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	// This points to your LOCAL folder internal/models
	"github.com/AnonPhoenix420/cyph3r/internal/models" 
	"github.com/likexian/whois"
)

// GetFullIntel aggregates DNS, WHOIS, and Geo-IP data.
func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Resolve DNS (via dns.go in this same package)
	data.IPs, data.Nameservers = LookupNodes(target)

	// 2. WHOIS Lookup
	rawWhois, err := whois.Whois(target)
	if err == nil {
		data.WhoisRaw = rawWhois
		data.Registrar = extractField(rawWhois, "Registrar:")
	}

	// 3. Geo-IP Intelligence
	client := http.Client{Timeout: time.Second * 5}
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon,isp,org,as", target)
	
	resp, err := client.Get(apiURL)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var apiRes struct {
			Status  string  `json:"status"`
			Country string  `json:"country"`
			City    string  `json:"city"`
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			ISP     string  `json:"isp"`
			Org     string  `json:"org"`
			ASN     string  `json:"as"`
		}
		if err := json.Unmarshal(body, &apiRes); err == nil && apiRes.Status == "success" {
			data.Country = apiRes.Country
			data.City = apiRes.City
			data.Lat = apiRes.Lat
			data.Lon = apiRes.Lon
			data.ISP = apiRes.ISP
			data.Org = apiRes.Org
			data.ASN = apiRes.ASN
		}
	}

	return data, nil
}

// extractField helps pull registrar info from WHOIS
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
