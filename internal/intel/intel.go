package intel

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
	"github.com/likexian/whois"
	"github.com/nyaruka/phonenumbers"
)

type NodeIntel struct {
	IPs, NS                       []string
	Registrar, BornOn, ISP, Coords string
	City, Country                 string
}

func GetFullIntel(target string) (*NodeIntel, error) {
	data := &NodeIntel{}
	ips, _ := net.LookupIP(target)
	for _, ip := range ips { data.IPs = append(data.IPs, ip.String()) }
	
	w, _ := whois.Whois(target)
	lines := strings.Split(w, "\n")
	for _, l := range lines {
		if strings.Contains(strings.ToLower(l), "registrar:") { data.Registrar = strings.TrimSpace(strings.Split(l, ":")[1]) }
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, _ := client.Get(fmt.Sprintf("http://ip-api.com/json/%s", target))
	if resp != nil {
		defer resp.Body.Close()
		var geo struct { ISP, City, Country string; Lat, Lon float64 }
		json.NewDecoder(resp.Body).Decode(&geo)
		data.ISP, data.City, data.Country = geo.ISP, geo.City, geo.Country
		data.Coords = fmt.Sprintf("%.4f, %.4f", geo.Lat, geo.Lon)
	}
	return data, nil
}

func PhoneLookup(num string) string {
	p, _ := phonenumbers.Parse(num, "")
	return fmt.Sprintf("VALID: %v | REGION: %s", phonenumbers.IsValidNumber(p), phonenumbers.GetRegionCodeForNumber(p))
}
