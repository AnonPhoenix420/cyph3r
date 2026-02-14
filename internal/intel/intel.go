package intel

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.NameServers = make(map[string]string)

	// 1. REVERSIBLE TARGETING & IP ENUMERATION
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

	// 2. NAME SERVER SPIDERING (Name -> IP Mapping)
	nsRecords, _ := net.LookupNS(data.TargetName)
	for _, ns := range nsRecords {
		nsIPs, _ := net.LookupIP(ns.Host)
		if len(nsIPs) > 0 {
			data.NameServers[ns.Host] = nsIPs[0].String()
		} else {
			data.NameServers[ns.Host] = "Not Resolved"
		}
	}

	// 3. REMOTE GEO RECON (Using target IP only)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.IP + "?fields=status,country,countryCode,regionName,city,zip,lat,lon,isp,org")
	if err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&data)
	}

	return data, nil
}
