package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/nyaruka/phonenumbers"
)

type NodeIntel struct {
	IPs     []string
	ISP     string
	City    string
	Country string
}

// GetFullIntel gathers ISP and Geo-data using public APIs
func GetFullIntel(target string) (*NodeIntel, error) {
	data := &NodeIntel{}

	// 1. DNS Lookup
	ips, _ := net.LookupIP(target)
	for _, ip := range ips {
		data.IPs = append(data.IPs, ip.String())
	}

	// 2. IP-API Lookup (No Key Required)
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s", target))
	if err == nil && resp != nil {
		defer resp.Body.Close()
		var geo struct {
			Status  string \`json:"status"\`
			ISP     string \`json:"isp"\`
			City    string \`json:"city"\`
			Country string \`json:"country"\`
		}
		json.NewDecoder(resp.Body).Decode(&geo)
		if geo.Status == "success" {
			data.ISP = geo.ISP
			data.City = geo.City
			data.Country = geo.Country
		}
	}

	return data, nil
}

// PhoneLookup provides carrier/regional metadata
func PhoneLookup(num string) string {
	p, err := phonenumbers.Parse(num, "")
	if err != nil {
		return "Invalid Format"
	}
	isValid := phonenumbers.IsValidNumber(p)
	region := phonenumbers.GetRegionCodeForNumber(p)
	return fmt.Sprintf("VALID: %v | REGION: %s", isValid, region)
}
