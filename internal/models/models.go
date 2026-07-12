package models

import "time"

type TargetType string

const (
	TargetPhone       TargetType = "phone"
	TargetEmail       TargetType = "email"
	TargetIP          TargetType = "ip"
	TargetDomain      TargetType = "domain"
	TypePhoneTarget   TargetType = "phone"
	TypeEmailTarget   TargetType = "email"
	TypeGeoTarget     TargetType = "geo"
	TypeNetworkTarget TargetType = "network"
)

type GeoData struct {
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Timezone     string `json:"timezone"`
	MapReference string `json:"map_reference"`
}

type SQLExposure struct {
	Exposed   bool  `json:"exposed"`
	Ports     []int `json:"ports"`
	RiskLevel string `json:"risk_level"`
}

type IntelPayload struct {
	Target          string              `json:"target"`
	Type            TargetType          `json:"type"`
	ScanTime        time.Time           `json:"scan_time"`
	Phone           string              `json:"phone"`
	OwnerName       string              `json:"owner_name"`
	ASN             string              `json:"asn"`
	ISP             string              `json:"isp"`
	Geo             GeoData             `json:"geo"`
	CreatedDate     string              `json:"created_date"`
	OpenPorts       []string            `json:"open_ports"`
	Banners         []string            `json:"banners"`
	Vulnerabilities []string            `json:"vulnerabilities"`
	ExposedLeaks    []string            `json:"exposed_leaks"`
	
	// Added fields for ResolveNetworkElite mapping
	TargetIP        string              `json:"target_ip"`
	SQLMetrics      SQLExposure         `json:"sql_metrics"`
	
	Verbose         bool                `json:"verbose"`
	OutputFormat    string              `json:"output_format"`
	Clusters        []string            `json:"clusters"`
	
	// HTTP Header Interception Telemetry Elements
	HTTPMethod      string              `json:"http_method"`
	CapturedHeaders map[string][]string `json:"captured_headers"`
}

type LocationData struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	ZIP         string  `json:"zip"`
	AreaCode    string  `json:"area_code"`
	Coordinates string  `json:"coordinates"`
	RadiusKM    float64 `json:"radius_km"`
}

type SocialProfile struct {
	Platform    string `json:"platform"`
	Username    string `json:"username"`
	ProfileURL  string `json:"profile_url"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	Confidence  int    `json:"confidence"`
}

type PortInfo struct {
	Port       int    `json:"port"`
	Service    string `json:"service"`
	State      string `json:"state"`
	Banner     string `json:"banner"`
	Vulnerable bool   `json:"vulnerable"`
	AdminPort  bool   `json:"admin_port"`
}

type ComprehensiveReport struct {
	Target         string          `json:"target"`
	TargetType     TargetType      `json:"target_type"`
	ReverseDNS     string          `json:"reverse_dns"`
	Location       LocationData    `json:"location"`
	Associated     []string        `json:"associated"`
	SocialProfiles []SocialProfile `json:"social_profiles"`
	Ports          []PortInfo      `json:"ports"`
	SQLCheck       SQLExposure     `json:"sql_check"`
	Timestamp      time.Time       `json:"timestamp"`
	RiskScore      int             `json:"risk_score"`
}
