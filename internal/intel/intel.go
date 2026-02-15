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
	data.TargetName = input // FIX: Ensures the input is explicitly stored
	data.NameServers = make(map[string][]string)

	// 1. Resolve ALL IPs (Dual Stack)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}
	
	// 2. Resolve Name Servers with BOTH IP types
	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		var nsAddrList []string
		nsIPs, _ := net.LookupIP(ns.Host)
		for _, ip := range nsIPs {
			nsAddrList = append(nsAddrList, ip.String())
		}
		data.NameServers[ns.Host] = nsAddrList
	}

	// 3. Geographic Deep-Dive
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org")
	if err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&data)
	}

	return data, nil
}
