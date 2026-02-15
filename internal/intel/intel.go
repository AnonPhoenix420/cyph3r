package intel

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// GetTargetIntel handles the original network recon logic
func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.NameServers = make(map[string]string)

	parsedIP := net.ParseIP(input)
	if parsedIP != nil {
		data.IP = input
		data.TargetIPs = append(data.TargetIPs, input)
		names, _ := net.LookupAddr(input)
		if len(names) > 0 { data.TargetName = names[0] }
	} else {
		data.TargetName = input
		ips, _ := net.LookupIP(input)
		for _, ip := range ips {
			data.TargetIPs = append(data.TargetIPs, ip.String())
		}
		if len(data.TargetIPs) > 0 { data.IP = data.TargetIPs[0] }
	}

	nsRecords, _ := net.LookupNS(data.TargetName)
	for _, ns := range nsRecords {
		nsIPs, _ := net.LookupIP(ns.Host)
		if len(nsIPs) > 0 {
			data.NameServers[ns.Host] = nsIPs[0].String()
		}
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.IP + "?fields=status,country,countryCode,regionName,city,zip,lat,lon,isp,org")
	if err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&data)
	}

	return data, nil
}

// GetPhoneIntel handles the NEW phone metadata logic (REQUIRED for main.go)
func GetPhoneIntel(number string) (models.PhoneData, error) {
	var data models.PhoneData
	// Mock response for stability; you can connect a real API here later
	data.Number = number
	data.Valid = true
	data.LocalFormat = number
	data.Carrier = "Global Gateway"
	data.Location = "Detected"
	data.Type = "Mobile/VOIP"
	
	return data, nil
}
