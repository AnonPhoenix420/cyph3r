package intel

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Resolve Target IP
	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		return data, err
	}
	data.IP = ips[0].String()

	// 2. Localhost Identity (Notation)
	hostname, _ := os.Hostname()
	data.LocalHost = hostname
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					data.LocalIPs = append(data.LocalIPs, ipnet.IP.String())
				}
			}
		}
	}

	// 3. DNS Name Servers
	ns, _ := net.LookupNS(target)
	for _, nameserver := range ns {
		data.NameServers = append(data.NameServers, nameserver.Host)
	}

	// 4. Deep Recon API (Geo + Org + Postal)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/" + data.IP + "?fields=status,message,country,countryCode,regionName,city,zip,lat,lon,isp,org,as,reverse")
	if err == nil {
		defer resp.Body.Close()
		var apiResult struct {
			Country     string  `json:"country"`
			CountryCode string  `json:"countryCode"`
			RegionName  string  `json:"regionName"`
			City        string  `json:"city"`
			Zip         string  `json:"zip"`
			Lat         float64 `json:"lat"`
			Lon         float64 `json:"lon"`
			ISP         string  `json:"isp"`
			Org         string  `json:"org"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResult)

		data.Country = apiResult.Country
		data.CountryCode = apiResult.CountryCode
		data.RegionName = apiResult.RegionName
		data.Zip = apiResult.Zip
		data.City = apiResult.City
		data.Lat = apiResult.Lat
		data.Lon = apiResult.Lon
		data.ISP = apiResult.ISP
		data.Org = apiResult.Org
	}

	return data, nil
}
