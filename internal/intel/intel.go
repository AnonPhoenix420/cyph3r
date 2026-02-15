package intel

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)

	// Resolve Target IPs
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// Recursive NS Resolution
	ns, _ := net.LookupNS(input)
	for _, n := range ns {
		data.NameServers["NS"] = append(data.NameServers["NS"], n.Host)
		nsIPs, _ := net.LookupIP(n.Host)
		for _, nip := range nsIPs {
			key := "IP_" + n.Host
			data.NameServers[key] = append(data.NameServers[key], nip.String())
		}
	}
	
	data.NameServers["PORTS"] = probes.ScanPorts(input)

	// Geo-Intel
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,isp,org")
		if err == nil {
			defer resp.Body.Close()
			var t struct {
				Country, RegionName, City, Zip, Isp, Org string
			}
			json.NewDecoder(resp.Body).Decode(&t)
			data.Country, data.State, data.City, data.Zip = t.Country, t.RegionName, t.City, t.Zip
			data.ISP, data.Org = t.Isp, t.Org
		}
	}
	return data, nil
}
