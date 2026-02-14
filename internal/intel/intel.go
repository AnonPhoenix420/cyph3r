package intel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"cyph3r/internal/models"
	"github.com/likexian/whois"
)

func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData
	data.IPs, data.Nameservers = LookupNodes(target)

	rawWhois, err := whois.Whois(target)
	if err == nil {
		data.WhoisRaw = rawWhois
		data.Registrar = extractField(rawWhois, "Registrar:")
	}

	query := target
	if len(data.IPs) > 0 { query = data.IPs[0] }

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon,isp,org,as", query))
	if err == nil {
		defer resp.Body.Close()
		var res struct {
			Status string `json:"status"`
			Country string `json:"country"`
			City string `json:"city"`
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
			ISP string `json:"isp"`
			Org string `json:"org"`
			ASN string `json:"as"`
		}
		json.NewDecoder(resp.Body).Decode(&res)
		if res.Status == "success" {
			data.Country, data.City, data.Lat, data.Lon = res.Country, res.City, res.Lat, res.Lon
			data.ISP, data.Org, data.ASN = res.ISP, res.Org, res.ASN
		}
	}
	return data, nil
}

func extractField(raw, field string) string {
	for _, line := range strings.Split(raw, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(field)) {
			parts := strings.Split(line, ":")
			if len(parts) > 1 { return strings.TrimSpace(parts[1]) }
		}
	}
	return "UNKNOWN"
}
