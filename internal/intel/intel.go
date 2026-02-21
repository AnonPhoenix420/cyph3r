package intel

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	shield := CheckShield()
	if !shield.IsActive { return models.IntelData{}, fmt.Errorf("VPN_OFFLINE") }

	data := models.IntelData{TargetName: input, NameServers: make(map[string][]string)}

	// DNS Resolution (IPv4 Focus for Anonymity)
	ips, _ := net.LookupIP(input)
	for _, ip := range ips {
		if ip.To4() != nil {
			data.TargetIPs = append(data.TargetIPs, ip.String())
		}
	}

	// Origin Hunting & Subdomains
	huntSubdomains(input, &data)
	
	// WAF Fingerprinting
	analyzeExploitSurface(input, &data)

	return data, nil
}

func huntSubdomains(target string, data *models.IntelData) {
	subs := []string{"dev", "vpn", "mail", "api", "test", "staging"}
	for _, s := range subs {
		host := s + "." + target
		if ips, err := net.LookupIP(host); err == nil && len(ips) > 0 {
			ip := ips[0].String()
			data.ScanResults = append(data.ScanResults, fmt.Sprintf("SUBDOMAIN: %s → %s", host, ip))
		}
	}
}

func analyzeExploitSurface(target string, data *models.IntelData) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", "http://"+target, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	
	resp, err := client.Do(req)
	if err != nil { return }
	defer resp.Body.Close()

	srv := resp.Header.Get("Server")
	if strings.Contains(strings.ToLower(srv), "arvan") || resp.Header.Get("ArvanCloud-Trace") != "" {
		data.IsWAF, data.WAFType = true, "ArvanCloud (Regional WAF)"
	}
	data.ScanResults = append(data.ScanResults, "STACK: "+srv)
}

func GetPhoneIntel(n string) (models.PhoneData, error) {
	return models.PhoneData{Number: n, Carrier: "MCI/Irancell", Risk: "LOW"}, nil
}
