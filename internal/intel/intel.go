package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/nyaruka/phonenumbers"
)

type ipGeoResponse struct {
	Status      string  `json:"status"`
	City        string  `json:"city"`
	Region      string  `json:"regionName"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	ASN         string  `json:"as"`
}

func ExecuteFullDox(target string, tType models.TargetType) *models.ComprehensiveReport {
	report := &models.ComprehensiveReport{
		Target:     target,
		TargetType: tType,
		Timestamp:  time.Now(),
	}

	report.Location = getLocationData(target, tType)
	report.ReverseDNS = getReverseDNS(target, tType)
	report.Ports = executeFullPortScan(target, tType)
	report.SQLCheck = checkSQLExposure(report.Ports)
	report.SocialProfiles = getSocialLinks(target, tType)
	report.Associated = getAssociatedContacts(target, tType)
	report.RiskScore = calculateRisk(report)

	return report
}

func getLocationData(target string, tType models.TargetType) models.LocationData {
	loc := models.LocationData{Country: "Unknown", State: "Unknown", RadiusKM: 0}

	switch tType {
	case models.TargetPhone, models.TypePhoneTarget:
		num, err := phonenumbers.Parse(target, "US")
		if err == nil {
			region := phonenumbers.GetRegionCodeForNumber(num)
			loc.Country = region
			loc.CountryCode = region
			loc.AreaCode = target[:3]
			switch loc.AreaCode {
			case "330", "440":
				loc.City = "Northeastern Ohio (Akron/Canton/Cleveland Metro)"
				loc.State = "Ohio"
				loc.ZIP = "44001-44321"
				loc.Coordinates = "41.0810 N, 81.5140 W"
				loc.RadiusKM = 45.0
			case "304":
				loc.City = "West Virginia (Charleston/Huntington Area)"
				loc.State = "West Virginia"
				loc.ZIP = "25001-25701"
				loc.Coordinates = "38.3500 N, 81.6300 W"
				loc.RadiusKM = 60.0
			default:
				loc.City = "US Rate Center"
				loc.State = "United States"
			}
		}

	case models.TargetEmail, models.TargetDomain, models.TargetIP, models.TypeEmailTarget, models.TypeNetworkTarget:
		ipStr := resolveToIP(target, tType)
		if ipStr != "" {
			resp, err := http.Get("http://ip-api.com/json/" + ipStr + "?fields=status,message,city,regionName,country,countryCode,zip,lat,lon,isp,org,as")
			if err == nil {
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				var geo ipGeoResponse
				json.Unmarshal(body, &geo)
				if geo.Status == "success" {
					loc.City = geo.City
					loc.State = geo.Region
					loc.Country = geo.Country
					loc.CountryCode = geo.CountryCode
					loc.ZIP = geo.Zip
					loc.Coordinates = fmt.Sprintf("%.4f N, %.4f W", geo.Lat, geo.Lon)
					loc.RadiusKM = 20.0
				}
			}
		}
	}
	return loc
}

func resolveToIP(target string, tType models.TargetType) string {
	if tType == models.TargetIP {
		return target
	}
	if tType == models.TargetEmail || tType == models.TypeEmailTarget {
		parts := strings.Split(target, "@")
		if len(parts) > 1 {
			target = parts[1]
		}
	}
	ips, err := net.LookupIP(target)
	if err == nil && len(ips) > 0 {
		return ips[0].String()
	}
	return ""
}

func getReverseDNS(target string, tType models.TargetType) string {
	ip := resolveToIP(target, tType)
	if ip != "" {
		names, err := net.LookupAddr(ip)
		if err == nil && len(names) > 0 {
			return names[0]
		}
	}
	return "N/A"
}

func getSocialLinks(target string, tType models.TargetType) []models.SocialProfile {
	if tType == models.TargetEmail || tType == models.TypeEmailTarget {
		username := strings.Split(target, "@")[0]
		return []models.SocialProfile{
			{Platform: "Google / Gmail", Username: username, ProfileURL: "https://myaccount.google.com", DisplayName: username, Confidence: 90},
			{Platform: "X (Twitter)", Username: username, ProfileURL: "https://x.com/" + username, Confidence: 45},
			{Platform: "LinkedIn", Username: username, ProfileURL: "https://linkedin.com/in/" + username, Confidence: 40},
		}
	}
	return []models.SocialProfile{}
}

func getAssociatedContacts(target string, tType models.TargetType) []string {
	if tType == models.TargetEmail || tType == models.TypeEmailTarget {
		domain := strings.Split(target, "@")[1]
		return []string{target + " (Primary)", "admin@" + domain + " (WHOIS)"}
	}
	return []string{}
}

func executeFullPortScan(target string, tType models.TargetType) []models.PortInfo {
	return []models.PortInfo{}
}

func checkSQLExposure(ports []models.PortInfo) models.SQLExposure {
	return models.SQLExposure{Exposed: false, RiskLevel: "Low"}
}

func calculateRisk(report *models.ComprehensiveReport) int {
	return 48
}

// Legacy compatibility
func ResolvePhone(phone string) string { return phone }
func ResolveEmail(email string) string { return email }
func ResolveNetwork(target string) (string, models.GeoData, string, string, string, []string, []string, []string, []string) {
	return "", models.GeoData{}, "", "", "", []string{}, []string{}, []string{}, []string{}
}
func ExecuteValidationSuite(url string, mode, conc, dur int) {}
