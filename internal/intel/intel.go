package intel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
}

func GetTargetIntel(input string) (models.IntelData, error) {
	// Re-verify shield immediately before scan
	shield := CheckShield()
	if !shield.IsActive {
		fmt.Println("\n\033[31m[!] PROTON VPN DISCONNECTED. EMERGENCY HALT.\033[0m")
		os.Exit(1)
	}

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		ipStr := ip.String()
		data.TargetIPs = append(data.TargetIPs, ipStr)
		names, _ := net.LookupAddr(ipStr)
		if len(names) > 0 {
			data.ReverseDNS = append(data.ReverseDNS, strings.TrimSuffix(names[0], "."))
		} else {
			data.ReverseDNS = append(data.ReverseDNS, "NO_PTR")
		}
	}
	
	if len(data.TargetIPs) > 0 {
		geo, raw := fetchGeo(data.TargetIPs[0])
		data.Org, data.City, data.Country = geo.Org, geo.City, geo.Country
		data.Lat, data.Lon = geo.Lat, geo.Lon
		
		usage := "RESIDENTIAL"
		if geo.Hosting { usage = "DATA_CENTER" }
		if geo.Proxy { usage += "/PROXY" }
		data.ScanResults = append(data.ScanResults, "USAGE: "+usage)
		data.RawGeo = raw
		data.Latency = pingTarget(data.TargetIPs[0])
	}

	detectWAF(input, &data)
	data.ScanResults = append(data.ScanResults, performTacticalScan(input)...)
	return data, nil
}

// ... [Keep detectWAF, fetchGeo, pingTarget, performTacticalScan from previous steps] ...
