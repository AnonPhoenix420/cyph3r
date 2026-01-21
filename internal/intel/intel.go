package intel

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/nyaruka/phonenumbers"
)

// ================= WHOIS =================

func WhoisLookup(target string) (string, error) {
	conn, err := net.DialTimeout("tcp", "whois.iana.org:43", 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, _ = conn.Write([]byte(target + "\r\n"))
	buf, _ := io.ReadAll(conn)
	return string(buf), nil
}

// ================= PHONE =================

func LookupPhone(number string) (map[string]interface{}, error) {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"raw":    number,
		"valid":  phonenumbers.IsValidNumber(num),
		"region": phonenumbers.GetRegionCodeForNumber(num),
		"type":   phonenumbers.GetNumberType(num),
	}, nil
}

// ================= ICMP =================

func CheckICMP(target string) bool {
	conn, err := net.DialTimeout("ip4:icmp", target, time.Second*2)
	if err != nil {
		return false
	}
	defer conn.Close()
	_, err = conn.Write([]byte("ping"))
	return err == nil
}

// ================= GEOIP =================

type GeoResult struct {
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	ASN      string  `json:"as"`
	Org      string  `json:"org"`
	Hostname string  `json:"reverse"`
}

func GeoIPLookup(target string) (*GeoResult, error) {
	resp, err := http.Get("http://ip-api.com/json/" + target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Country string `json:"country"`
		Region  string `json:"regionName"`
		City    string `json:"city"`
		Lat     float64
		Lon     float64
		ASN     string `json:"as"`
		Org     string `json:"org"`
		Host    string `json:"reverse"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if data.Status != "success" {
		return nil, fmt.Errorf(data.Message)
	}

	return &GeoResult{
		City:     data.City,
		Region:   data.Region,
		Country:  data.Country,
		Lat:      data.Lat,
		Lon:      data.Lon,
		ASN:      data.ASN,
		Org:      data.Org,
		Hostname: data.Host,
	}, nil
}
