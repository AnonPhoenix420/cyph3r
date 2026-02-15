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

	// 1. Resolve ALL Target IPs (IPv4 + IPv6)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}
	if len(data.TargetIPs) > 0 {
		data.IP = data.TargetIPs[0]
	}

	// 2. Resolve Name Servers with IPv4 priority for "Local" style display
	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		nsIPs, _ := net.LookupIP(ns.Host)
		if len(nsIPs) > 0 {
			// Save the first resolved IP (usually IPv4)
			data.NameServers[ns.Host] = nsIPs[0].String()
		}
	}

	// 3. Deep Geo-Intelligence (Restoring the lost data)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.IP + "?fields=status,message,country,countryCode,regionName,city,zip,lat,lon,isp,org,as")
	if err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&data)
	}

	return data, nil
}
