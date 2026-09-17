package intel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

type ShieldStatus struct {
	IsActive       bool   `json:"is_active"`
	Location       string `json:"location"`
	ISP            string `json:"isp"`
	VPNDetected    bool   `json:"vpn_detected"`
	Recommendation string `json:"recommendation"`
}

// CheckShield verifies VPN / Shield status via IP check (OPSEC Awareness)
func CheckShield() ShieldStatus {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return ShieldStatus{
			IsActive:       false,
			VPNDetected:    false,
			Recommendation: "Unable to verify connection shield - check network connectivity",
		}
	}
	defer resp.Body.Close()

	var r struct {
		Status      string `json:"status"`
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		Region      string `json:"regionName"`
		City        string `json:"city"`
		ISP         string `json:"isp"`
		Org         string `json:"org"`
		Query       string `json:"query"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil || r.Status != "success" {
		return ShieldStatus{
			IsActive:       false,
			VPNDetected:    false,
			Recommendation: "Failed to parse shield telemetry payload",
		}
	}

	// Detect common VPN / Datacenter / Privacy providers across ISP and Org strings
	combinedMetadata := strings.ToLower(r.ISP + " " + r.Org)
	isVPN := strings.Contains(combinedMetadata, "m247") ||
		strings.Contains(combinedMetadata, "proton") ||
		strings.Contains(combinedMetadata, "mullvad") ||
		strings.Contains(combinedMetadata, "datacentre") ||
		strings.Contains(combinedMetadata, "vpn") ||
		strings.Contains(combinedMetadata, "cloud") ||
		strings.Contains(combinedMetadata, "expressvpn") ||
		strings.Contains(combinedMetadata, "nordvpn") ||
		strings.Contains(combinedMetadata, "digitalocean") ||
		strings.Contains(combinedMetadata, "aws") ||
		strings.Contains(combinedMetadata, "ovh")

	recommendation := "Connection appears unprotected (Direct exposure risk)"
	if isVPN {
		recommendation = "Active shield / VPN detected - Good OPSEC posture"
	}

	locString := "Unknown Location"
	if r.City != "" && r.Country != "" {
		locString = fmt.Sprintf("%s, %s, %s", r.City, r.Region, r.Country)
	} else if r.Country != "" {
		locString = r.Country
	}

	return ShieldStatus{
		IsActive:       true,
		VPNDetected:    isVPN,
		Location:       locString,
		ISP:            r.ISP,
		Recommendation: recommendation,
	}
}

// GetShieldReport returns a formatted report compatible with ComprehensiveReport
func GetShieldReport() models.ComprehensiveReport {
	shield := CheckShield()

	// Safely parse location components without index-out-of-range panics
	parts := strings.Split(shield.Location, ", ")
	city := "Unknown"
	state := ""
	country := "Unknown"

	if len(parts) >= 3 {
		city = parts[0]
		state = parts[1]
		country = parts[2]
	} else if len(parts) == 2 {
		city = parts[0]
		country = parts[1]
	} else if len(parts) == 1 {
		country = parts[0]
	}

	risk := 75
	if shield.VPNDetected {
		risk = 25
	}

	return models.ComprehensiveReport{
		Target:     "Current Connection Shield",
		TargetType: models.TargetIP,
		Location: models.LocationData{
			City:    city,
			State:   state,
			Country: country,
		},
		Associated: []string{
			"ISP / Provider: " + shield.ISP,
			"Status: " + shield.Recommendation,
		},
		RiskScore: risk,
		Timestamp: time.Now(),
	}
}
