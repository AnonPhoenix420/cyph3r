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

type TargetIntel struct {
	IP        string  `json:"query"`
	Status    string  `json:"status"`
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	City      string  `json:"city"`
	Zip       string  `json:"zip"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	ISP       string  `json:"isp"`
	Org       string  `json:"org"`
	AS        string  `json:"as"`
	Reverse   string  `json:"reverse"`
	MapsURL   string  `json:"maps_url"`
}

// GlobalLookup combines GeoIP, Maps, and ISP handling
func GlobalLookup(target string) (*TargetIntel, error) {
	resp, err := http.Get("http://ip-api.com/json/" + target + "?fields=status,message,country,regionName,city,zip,lat,lon,isp,org,as,reverse,query")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var intel TargetIntel
	if err := json.NewDecoder(resp.Body).Decode(&intel); err != nil {
		return nil, err
	}

	if intel.Status == "success" {
		// Generate Google Maps link based on coordinates
		intel.MapsURL = fmt.Sprintf("https://www.google.com/maps?q=%f,%f", intel.Lat, intel.Lon)
	}
	return &intel, nil
}

func Whois(target string) string {
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err != nil {
		return "Whois lookup failed: " + err.Error()
	}
	defer conn.Close()
	conn.Write([]byte(target + "\r\n"))
	res, _ := io.ReadAll(conn)
	return string(res)
}

func PhoneIntel(number string) map[string]interface{} {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		return map[string]interface{}{"error": "Invalid format"}
	}
	return map[string]interface{}{
		"Valid":    phonenumbers.IsValidNumber(num),
		"Region":   phonenumbers.GetRegionCodeForNumber(num),
		"Type":     phonenumbers.GetNumberType(num),
		"Provider": "Check local carrier database", // Requires DB for offline, API for online
	}
}
