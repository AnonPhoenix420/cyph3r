package intel

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ShieldStatus struct {
	IsActive bool
	Location string
	ISP      string
}

// CheckShield verifies VPN status via IP check
func CheckShield() ShieldStatus {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return ShieldStatus{IsActive: false}
	}
	defer resp.Body.Close()

	var r struct {
		Status  string `json:"status"`
		Country string `json:"country"`
		Isp     string `json:"isp"`
		Query   string `json:"query"`
	}
	json.NewDecoder(resp.Body).Decode(&r)

	// OPSEC: Detecting VPN Providers (Proton, M247, etc)
	isVpn := strings.Contains(strings.ToLower(r.Isp), "m247") || 
		     strings.Contains(strings.ToLower(r.Isp), "proton") ||
			 strings.Contains(strings.ToLower(r.Isp), "datacentre")

	return ShieldStatus{
		IsActive: isVpn,
		Location: r.Country,
		ISP:      r.Isp,
	}
}
