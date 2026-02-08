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
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: time.Second * 3}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	ips, _ := r.LookupIP(context.Background(), "ip", target)
	for _, ip := range ips { data.IPs = append(data.IPs, ip.String()) }

	ns, _ := r.LookupNS(context.Background(), target)
	for _, n := range ns { data.Nameservers = append(data.Nameservers, n.Host) }

	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,lat,lon", target)
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var apiRes struct {
			Country string  `json:"country"`
			City    string  `json:"city"`
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
		}
		json.Unmarshal(body, &apiRes)
		data.Country, data.City, data.Lat, data.Lon = apiRes.Country, apiRes.City, apiRes.Lat, apiRes.Lon
	}
	return data, nil
}
