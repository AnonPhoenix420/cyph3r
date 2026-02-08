package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/likexian/whois" // The new WHOIS dependency
)

func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. DNS & Identity (from dns.go)
	data.IPs, data.Nameservers = LookupNodes(target)

	// 2. WHOIS Intelligence (New Section)
	// We fetch the raw record; in a more advanced version, you could use a parser
	// for now, we'll grab the raw block to extract key strings.
	rawWhois, err := whois.Whois(target)
	if err == nil {
		data.WhoisRaw = rawWhois // Store it in your model for the HUD
	}

	// 3. Geo-IP & ISP Intelligence
	client := http.Client{Timeout: time.Second * 5}
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,countryCode,regionName,city,lat,lon,isp,org,as", target)
	
	resp, err := client.Get(apiURL)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

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
		}
	}

	return data, nil
}
