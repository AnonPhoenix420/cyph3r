package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/nyaruka/phonenumbers"
)

type NodeIntel struct {
	IPs, NS, NSIPs                 []string
	Country, CountryCode, Region   string
	RegionCode, City, Zip, TZ      string
	ISP, Org, ASN                  string
	Lat, Lon                       float64
}

func GetFullIntel(target string) (*NodeIntel, error) {
	data := &NodeIntel{}

	// 1. Resolve Target IPs
	ips, _ := net.LookupIP(target)
	for _, ip := range ips { data.IPs = append(data.IPs, ip.String()) }

	// 2. Resolve Name Servers and their IPs
	nsRecords, _ := net.LookupNS(target)
	for _, ns := range nsRecords {
		data.NS = append(data.NS, ns.Host)
		nsIPs, _ := net.LookupIP(ns.Host)
		for _, nip := range nsIPs {
			data.NSIPs = append(data.NSIPs, fmt.Sprintf("%s (%s)", ns.Host, nip.String()))
		}
	}

	// 3. Deep Geo-IP Lookup
	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,countryCode,region,regionName,city,zip,lat,lon,timezone,isp,org,as", target))
	if err == nil && resp != nil {
		defer resp.Body.Close()
		var geo struct {
			Status, Country, CountryCode, Region, RegionName string
			City, Zip, Timezone, ISP, Org, AS                string
			Lat, Lon                                         float64
		}
		json.NewDecoder(resp.Body).Decode(&geo)
		if geo.Status == "success" {
			data.Country, data.CountryCode = geo.Country, geo.CountryCode
			data.Region, data.RegionCode = geo.RegionName, geo.Region
			data.City, data.Zip, data.TZ = geo.City, geo.Zip, geo.Timezone
			data.ISP, data.Org, data.ASN = geo.ISP, geo.Org, geo.AS
			data.Lat, data.Lon = geo.Lat, geo.Lon
		}
	}
	return data, nil
}

func PhoneLookup(num string) string {
	p, _ := phonenumbers.Parse(num, "")
	return fmt.Sprintf("VALID: %v | REGION: %s", phonenumbers.IsValidNumber(p), phonenumbers.GetRegionCodeForNumber(p))
}
