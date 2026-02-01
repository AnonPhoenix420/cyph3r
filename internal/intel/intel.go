package intel

import (
	"context"
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
	IPs       []string
	NS        []string
	Registrar string
	BornOn    string
	ISP       string
	Org       string
	Country   string
	City      string
	Zip       string
	Lat       float64
	Lon       float64
	Coords    string
	Maps      string
}

func GetFullIntel(target string) (*NodeIntel, error) {
	// Go 1.24 optimized: 5-second deadline for the entire recon phase
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := &NodeIntel{}
	resolver := &net.Resolver{}

	// 1. Context-Aware DNS Lookups
	ips, _ := resolver.LookupIP(ctx, "ip", target)
	for _, ip := range ips {
		data.IPs = append(data.IPs, ip.String())
	}

	nss, _ := resolver.LookupNS(ctx, target)
	for _, ns := range nss {
		data.NS = append(data.NS, strings.TrimSuffix(ns.Host, "."))
	}

	// 2. WHOIS Intelligence
	// Note: whois library is not natively context-aware yet, so we wrap it
	whoisChan := make(chan string, 1)
	go func() {
		w, _ := whois.Whois(target)
		whoisChan <- w
	}()

	select {
	case w := <-whoisChan:
		lines := strings.Split(w, "\n")
		for _, line := range lines {
			lower := strings.ToLower(line)
			if strings.Contains(lower, "registrar:") && data.Registrar == "" {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					data.Registrar = strings.TrimSpace(parts[1])
				}
			}
			if (strings.Contains(lower, "creation date:") || strings.Contains(lower, "created:")) && data.BornOn == "" {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					data.BornOn = strings.TrimSpace(strings.Join(parts[1:], ":"))
				}
			}
		}
	case <-ctx.Done():
		data.Registrar = "TIMEOUT"
	}

	// 3. Optimized Geo-IP & ISP Intelligence
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,city,zip,lat,lon,isp,org", target), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		var geo struct {
			Status  string  `json:"status"`
			Country string  `json:"country"`
			City    string  `json:"city"`
			Zip     string  `json:"zip"`
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			ISP     string  `json:"isp"`
			Org     string  `json:"org"`
		}
		if json.NewDecoder(resp.Body).Decode(&geo) == nil && geo.Status == "success" {
			data.Country = geo.Country
			data.City = geo.City
			data.Zip = geo.Zip
			data.Lat = geo.Lat
			data.Lon = geo.Lon
			data.ISP = geo.ISP
			data.Org = geo.Org
			data.Coords = fmt.Sprintf("%.4f, %.4f", geo.Lat, geo.Lon)
			data.Maps = fmt.Sprintf("https://www.google.com/maps?q=%.4f,%.4f", geo.Lat, geo.Lon)
		}
	}

	// Global Fallbacks
	if data.Registrar == "" { data.Registrar = "Private/Protected" }
	if data.BornOn == "" { data.BornOn = "Unknown" }
	if data.ISP == "" { data.ISP = "Unknown Provider" }

	return data, nil
}

func PhoneLookup(number string) string {
	parsed, err := phonenumbers.Parse(number, "")
	if err != nil {
		return fmt.Sprintf("Error: Invalid Format [%s]", number)
	}
	region := phonenumbers.GetRegionCodeForNumber(parsed)
	return fmt.Sprintf("STATUS: %v | REGION: %s | CARRIER: %v", phonenumbers.IsValidNumber(parsed), region, phonenumbers.GetNumberType(parsed))
}
