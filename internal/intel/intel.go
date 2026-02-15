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

	// 1. RESOLVE ALL IPs (A & AAAA)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}

	// 2. DNS INFRASTRUCTURE LEAK
	ns, _ := net.LookupNS(input)
	for _, n := range ns {
		data.NameServers["NS"] = append(data.NameServers["NS"], n.Host)
	}

	mx, _ := net.LookupMX(input)
	for _, m := range mx {
		data.NameServers["MX"] = append(data.NameServers["MX"], m.Host)
	}

	txt, _ := net.LookupTXT(input)
	data.NameServers["TXT"] = txt

	// 3. DEEP GEO & ASN
	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org,as")
		if err == nil {
			defer resp.Body.Close()
			var t struct {
				Lat, Lon float64
				Country, RegionName, City, Zip, Isp, Org string
			}
			json.NewDecoder(resp.Body).Decode(&t)

			data.Lat, data.Lon = fmt.Sprintf("%.6f", t.Lat), fmt.Sprintf("%.6f", t.Lon)
			data.Country, data.State, data.City, data.Zip = t.Country, t.RegionName, t.City, t.Zip
			data.ISP, data.Org = t.Isp, t.Org
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
		}
	}
	return data, nil
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	// (Your phone logic remains here)
	return d, nil
}
