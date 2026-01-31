package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/nyaruka/phonenumbers" // Ensure this matches go.mod
)

type FullIntel struct {
	IP         string  `json:"query"`
	Status     string  `json:"status"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	ISP        string  `json:"isp"`
	Org        string  `json:"org"`
	AS         string  `json:"as"`
	ReverseDNS string  `json:"reverse"`
}

func GetIntel(target string) (*FullIntel, string, error) {
	// IP-API for Geo + ISP + Zip + Org
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=66846719", target)
	resp, err := http.Get(url)
	if err != nil { return nil, "", err }
	defer resp.Body.Close()

	var data FullIntel
	json.NewDecoder(resp.Body).Decode(&data)

	// WHOIS for Handler details
	whois := "Whois data unavailable"
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err == nil {
		defer conn.Close()
		fmt.Fprintf(conn, "%s\r\n", target)
		res, _ := io.ReadAll(conn)
		whois = string(res)
	}

	return &data, whois, nil
}

func PhoneLookup(number string) string {
	num, err := phonenumbers.Parse(number, "")
	if err != nil { return "Invalid Format" }
	region := phonenumbers.GetRegionCodeForNumber(num)
	return fmt.Sprintf("Region: %s | Valid: %v | Type: %s", region, phonenumbers.IsValidNumber(num), phonenumbers.GetNumberType(num))
}
