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

// GetFullIntel serves as the primary data aggregator for a target node.
func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Resolve Network Identity (DNS)
	// LookupNodes is defined in dns.go within the same package.
	data.IPs, data.Nameservers = LookupNodes(target)

	// 2. Registry Intelligence (WHOIS)
	// We perform a deep query to find ownership and registrar details.
	rawWhois, err := whois.Whois(target)
	if err == nil {
		data.WhoisRaw = rawWhois
		data.Registrar = extractField(rawWhois, "Registrar:")
	}

	// 3. Geographic Intelligence (Geo-IP)
	// We use an external API to map the target's physical location.
	if len(data.IPs) > 0 {
		// We use the first resolved IP for geographic mapping.
		fetchGeoData(data.IPs[0], &data)
	} else {
		// If DNS didn't resolve, we try to map the domain string directly.
		fetchGeoData(target, &data)
	}

	return data, nil
}

// fetchGeoData queries the IP-API to populate geographic and ISP metadata.
func fetchGeoData(query string, data *models.IntelData) {
	client := http.Client{Timeout: time.Second * 5}
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,city,lat,lon,isp,org,as", query)

	resp, err := client.Get(apiURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

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

// extractField is a helper function to parse specific lines from raw WHOIS text.
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
