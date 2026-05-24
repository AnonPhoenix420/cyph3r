package intel

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

type ShieldStatus struct {
	IsActive     bool   `json:"is_active"`
	Location     string `json:"location"`
	ISP          string `json:"isp"`
	VPNDetected  bool   `json:"vpn_detected"`
	Recommendation string `json:"recommendation"`
}

// CheckShield verifies VPN / Shield status via IP check (OPSEC Awareness)
func CheckShield() ShieldStatus {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return ShieldStatus{
			IsActive:      false,
			VPNDetected:   false,
			Recommendation: "Unable to verify connection shield - check network",
		}
	}
	defer resp.Body.Close()

	var r struct {
		Status  string `json:"status"`
		Country string `json:"country"`
		Region  string `json:"regionName"`
		City    string `json:"city"`
		ISP     string `json:"isp"`
		Query   string `json:"query"`
	}

	json.NewDecoder(resp.Body).Decode(&r)

	// Detect common VPN / Datacenter providers
	ispLower := strings.ToLower(r.ISP)
	isVPN := strings.Contains(ispLower, "m247") ||
		strings.Contains(ispLower, "proton") ||
		strings.Contains(ispLower, "datacentre") ||
		strings.Contains(ispLower, "vpn") ||
		strings.Contains(ispLower, "cloud") ||
		strings.Contains(ispLower, "expressvpn") ||
		strings.Contains(ispLower, "nordvpn")

	recommendation := "Connection appears unprotected"
	if isVPN {
		recommendation = "Active shield / VPN detected - Good OPSEC"
	}

	return ShieldStatus{
		IsActive:      true,
		VPNDetected:   isVPN,
		Location:      r.City + ", " + r.Region + ", " + r.Country,
		ISP:           r.ISP,
		Recommendation: recommendation,
	}
}

// GetShieldReport returns a formatted report compatible with ComprehensiveReport
func GetShieldReport() models.ComprehensiveReport {
	shield := CheckShield()

	return models.ComprehensiveReport{
		Target:     "Current Connection Shield",
		TargetType: models.TargetIP,
		Location: models.LocationData{
			City:    shield.Location,
			Country: strings.Split(shield.Location, ", ")[2],
		},
		Associated: []string{
			"ISP: " + shield.ISP,
			shield.Recommendation,
		},
		RiskScore: func() int {
			if shield.VPNDetected {
				return 25
			}
			return 75
		}(),
		Timestamp: time.Now(),
	}
}
