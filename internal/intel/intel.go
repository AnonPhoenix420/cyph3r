package intel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
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

	queryIP := target
	if len(data.IPs) > 0 {
		queryIP = data.IPs[0]
	}

	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon,isp,org,as", queryIP))
	if err == nil {
		defer resp.Body.Close()
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
		json.NewDecoder(resp.Body).Decode(&apiRes)
		if apiRes.Status == "success" {
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

func extractField(raw, field string) string {
	for _, line := range strings.Split(raw, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(field)) {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "UNKNOWN"
}
