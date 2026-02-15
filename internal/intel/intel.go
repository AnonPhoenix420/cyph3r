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

	// 1. IP RESOLUTION
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// 2. DNS & PORTS
	ns, _ := net.LookupNS(input)
	for _, n := range ns { data.NameServers["NS"] = append(data.NameServers["NS"], n.Host) }
	
	// Call the scanner we moved to probes
	data.NameServers["PORTS"] = probes.ScanPorts(input)

	// 3. FULL GEO-INTEL (Ensuring Zip is captured)
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		// Added zip to the fields list
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
