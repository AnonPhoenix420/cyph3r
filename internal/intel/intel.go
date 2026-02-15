package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
	"strings"
	"cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips { data.TargetIPs = append(data.TargetIPs, ip.String()) }
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org")
	if err == nil {
		defer resp.Body.Close()
		var t struct { Lat, Lon float64; Reg, ISP, Org, City, Zip, Ctry string }
		json.NewDecoder(resp.Body).Decode(&t)
		data.Lat, data.Lon = fmt.Sprintf("%.6f", t.Lat), fmt.Sprintf("%.6f", t.Lon)
		data.State, data.ISP, data.Org, data.City, data.Zip, data.Country = t.Reg, t.ISP, t.Org, t.City, t.Zip, t.Ctry
		data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
	}
	return data, nil
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number, d.Valid = number, true
	clean := strings.TrimPrefix(number, "+")
	globalMap := map[string][]string{
		"1330": {"USA", "Ohio", "Akron/Canton", "41.0814", "-81.5190", "Verizon/AT&T"},
		"44":   {"UK", "United Kingdom", "London Hub", "51.5074", "-0.1278", "BT/Vodafone"},
	}
	for pfx, info := range globalMap {
		if strings.HasPrefix(clean, pfx) {
			d.Country, d.State, d.Location, d.Lat, d.Lon, d.Carrier = info[0], info[1], info[2], info[3], info[4], info[5]
			d.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", d.Lat, d.Lon)
			break
		}
	}
	return d, nil
}
