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

	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		data.TargetIPs = append(data.TargetIPs, ip.String())
	}
	
	nsRecords, _ := net.LookupNS(input)
	for _, ns := range nsRecords {
		nsIPs, _ := net.LookupIP(ns.Host)
		var addrList []string
		for _, ip := range nsIPs {
			addrList = append(addrList, ip.String())
		}
		data.NameServers[ns.Host] = addrList
	}

	if len(data.TargetIPs) > 0 {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get("http://ip-api.com/json/" + data.TargetIPs[0] + "?fields=status,country,regionName,city,zip,lat,lon,isp,org")
		if err == nil {
			defer resp.Body.Close()
			var temp struct {
				Lat    float64 `json:"lat"`
				Lon    float64 `json:"lon"`
				Region string  `json:"regionName"`
				ISP    string  `json:"isp"`
				Org    string  `json:"org"`
				City   string  `json:"city"`
				Zip    string  `json:"zip"`
				Ctry   string  `json:"country"`
			}
			json.NewDecoder(resp.Body).Decode(&temp)
			
			data.Lat = fmt.Sprintf("%.6f", temp.Lat)
			data.Lon = fmt.Sprintf("%.6f", temp.Lon)
			data.State = temp.Region
			data.ISP = temp.ISP
			data.Org = temp.Org
			data.City = temp.City
			data.Zip = temp.Zip
			data.Country = temp.Ctry
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
		}
	}
	return data, nil
}
