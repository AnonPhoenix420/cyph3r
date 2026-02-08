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

// GetFullIntel handles the heavy lifting of gathering network identity and location data.
func GetFullIntel(target string) (models.IntelData, error) {
	var data models.IntelData

	// 1. Handle Localhost / Internal Loopback
	if target == "localhost" || target == "127.0.0.1" {
		data.IPs = []string{"127.0.0.1"}
		data.Nameservers = []string{"LOCAL_NODE_INTERNAL"}
		data.City = "Home"
		data.Country = "Loopback"
		return data, nil
	}

	// 2. Configure High-Performance DNS Resolver (8.8.8.8)
	// This ensures we get NS records even if the local system is missing tools.
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: time.Second * 5}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	// 3. Resolve IP Addresses (A Records)
	ips, err := resolver.LookupIP(context.Background(), "ip", target)
	if err == nil {
		for _, ip := range ips {
			data.IPs = append(data.IPs, ip.String())
		}
	}

	// 4. Resolve Nameservers (NS Records)
	nsRecords, err := resolver.LookupNS(context.Background(), target)
	if err == nil {
		for _, ns := range nsRecords {
			data.Nameservers = append(data.Nameservers, ns.Host)
		}
	}

	// 5. Geographic Intelligence (Geo-IP)
	// Querying the API for location metadata
	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon", target))
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

		if err := json.Unmarshal(body, &apiRes); err == nil && apiRes.Status == "success" {
			data.Country = apiRes.Country
			data.City = apiRes.City
			data.Lat = apiRes.Lat
			data.Lon = apiRes.Lon
		}
	}

	return data, nil
}
