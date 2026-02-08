package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// GetFullIntel serves as the primary data aggregator for the system.
func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Resolve Network Identity (Calls dns.go)
	// Fetches the Neon Blue IPs and Neon Yellow Nameservers
	data.IPs, data.Nameservers = LookupNodes(target)

	// 2. Fetch Geographic & ISP Intelligence
	// Using a 5-second timeout to prevent the system from hanging
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// Requesting specific fields to populate our Neon Yellow Geo-HUD
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,countryCode,regionName,city,zip,lat,lon,timezone,isp,org,as", target)
	
	resp, err := client.Get(apiURL)
	if err != nil {
		return data, fmt.Errorf("NETWORK_TIMEOUT: GEOLOCATION_UNREACHABLE")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	// Temporary struct to map API response to our models
	var apiRes struct {
		Status      string  `json:"status"`
		Message     string  `json:"message"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"regionName"`
		City        string  `json:"city"`
		Zip         string  `json:"zip"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Timezone    string  `json:"timezone"`
		ISP         string  `json:"isp"`
		Org         string  `json:"org"`
		ASN         string  `json:"as"`
	}

	if err := json.Unmarshal(body, &apiRes); err != nil {
		return data, err
	}

	// 3. Map Intelligence to Central Model
	if apiRes.Status == "success" {
		data.Country = apiRes.Country
		data.CountryCode = apiRes.CountryCode
		data.Region = apiRes.Region
		data.City = apiRes.City
		data.Zip = apiRes.Zip
		data.Lat = apiRes.Lat
		data.Lon = apiRes.Lon
		data.Timezone = apiRes.Timezone
		data.ISP = apiRes.ISP
		data.Org = apiRes.Org
		data.ASN = apiRes.ASN
	}

	return data, nil
}
