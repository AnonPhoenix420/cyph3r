package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/nyaruka/phonenumbers"
)

type FullIntel struct {
	IP          string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	ReverseDNS  string  `json:"reverse"`
}

// GetIntel pulls everything: Geo, ISP, Org, Zip, and Hostname
func GetIntel(target string) (*FullIntel, error) {
	// Querying ip-api with all necessary fields
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,regionName,city,zip,lat,lon,isp,org,as,reverse,query", target)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data FullIntel
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Whois fetches the registrar and handler info
func Whois(target string) string {
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err != nil {
		return "Whois unavailable"
	}
	defer conn.Close()
	conn.Write([]byte(target + "\r\n"))
	res, _ := io.ReadAll(conn)
	return string(res)
}

// PhoneLookup provides carrier and region info
func PhoneLookup(number string) string {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		return "Invalid number format"
	}
	region := phonenumbers.GetRegionCodeForNumber(num)
	return fmt.Sprintf("Region: %s | Valid: %v | Type: %s", region, phonenumbers.IsValidNumber(num), phonenumbers.GetNumberType(num))
}
