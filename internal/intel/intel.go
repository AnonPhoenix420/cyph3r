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

	// 1. IP RESOLUTION (A & AAAA Records)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// 2. DNS INFRASTRUCTURE (NS & MX & TXT)
	ns, _ := net.LookupNS(input)
	for _, nameserver := range ns {
		data.NameServers["NS"] = append(data.NameServers["NS"], nameserver.Host)
	}

	mx, _ := net.LookupMX(input)
	for _, mail := range mx {
		data.NameServers["MX"] = append(data.NameServers["MX"], mail.Host)
	}

	txt, _ := net.LookupTXT(input)
	for _, t := range txt {
		data.NameServers["TXT"] = append(data.NameServers["TXT"], t)
	}

	// 3. DEEP GEOLOCATION & ASN
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		// Querying IP-API for the deep stack (ASN, ISP, Proxy status)
		resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org,as,mobile,proxy,hosting")
		if err == nil {
			defer resp.Body.Close()
			var t struct {
				Lat, Lon float64
				Status, Country, RegionName, City, Zip, Isp, Org, As string
				Mobile, Proxy, Hosting bool
			}
			json.NewDecoder(resp.Body).Decode(&t)

			data.Lat, data.Lon = fmt.Sprintf("%.6f", t.Lat), fmt.Sprintf("%.6f", t.Lon)
			data.Country, data.State, data.City, data.Zip = t.Country, t.RegionName, t.City, t.Zip
			data.ISP, data.Org = t.Isp, t.Org
			
			// Re-adding the real Map link
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
		}
	}
	return data, nil
}
