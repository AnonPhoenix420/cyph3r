package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)

	// 1. RESOLVE IPs
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// 2. RESOLVE NAMESERVERS (DNS Intel)
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		data.NameServers["DNS"] = append(data.NameServers["DNS"], nameserver.Host)
	}

	// 3. GEOLOCATION (The "Where" in the world)
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		// Querying IP-API for the full stack
		resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org,as")
		if err == nil {
			defer resp.Body.Close()
			var t struct {
				Lat, Lon float64
				Status, Country, RegionName, City, Zip, Isp, Org, As string
			}
			json.NewDecoder(resp.Body).Decode(&t)

			data.Lat, data.Lon = fmt.Sprintf("%.6f", t.Lat), fmt.Sprintf("%.6f", t.Lon)
			data.Country = t.Country
			data.State = t.RegionName
			data.City = t.City
			data.Zip = t.Zip
			data.ISP = t.Isp
			data.Org = t.Org
			// Generate a REAL clickable Google Maps link
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
		}
	}
	return data, nil
}
