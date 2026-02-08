package intel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// Handle Localhost/Loopback detection
	if target == "localhost" || target == "127.0.0.1" {
		data.IPs = []string{"127.0.0.1"}
		data.Nameservers = []string{"LOCAL_NODE_INTERNAL"}
		data.City = "Localhost"
		data.Country = "Loopback"
		return data, nil
	}

	// 1. Setup Custom Resolver (Bypasses system tool issues)
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: time.Second * 3}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	// 2. Resolve IP Addresses
	ips, _ := r.LookupIP(context.Background(), "ip", target)
	for _, ip := range ips {
		data.IPs = append(data.IPs, ip.String())
	}

	// 3. Resolve Nameservers (NS Records)
	ns, _ := r.LookupNS(context.Background(), target)
	for _, n := range ns {
		data.Nameservers = append(data.Nameservers, n.Host)
	}

	// 4. Geographic Data Fetch
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon", target)
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var apiRes struct {
			Status  string  `json:"status"`
			Country string  `json:"country"`
			City    string  `json:"city"`
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
		}
		json.Unmarshal(body, &apiRes)
		data.Country = apiRes.Country
		data.City = apiRes.City
		data.Lat = apiRes.Lat
		data.Lon = apiRes.Lon
	}

	return data, nil
}
