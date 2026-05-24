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
	Latitude     string
	Longitude    string
	City         string
	Country      string
	Timezone     string
	MapReference string
}

type IntelPayload struct {
	Target         string
	Type           TargetType
	ScanTime       time.Time
	Phone          string
	OwnerName      string
	ASN            string
	ISP            string
	Geo            GeoData
	CreatedDate    string
	OpenPorts      []string
	Banners        []string
	Vulnerabilities []string
	ExposedLeaks   []string
	Verbose        bool
	OutputFormat   string
	Clusters       []string
}

type LocationData struct {
	Country     string
	CountryCode string
	State       string
	City        string
	ZIP         string
	AreaCode    string
	Coordinates string
	RadiusKM    float64
}

type SocialProfile struct {
	Platform    string
	Username    string
	ProfileURL  string
	DisplayName string
	Bio         string
	Confidence  int
}

type PortInfo struct {
	Port       int
	Service    string
	State      string
	Banner     string
	Vulnerable bool
	AdminPort  bool
}

type SQLExposure struct {
	Exposed   bool
	Ports     []int
	RiskLevel string
}

type ComprehensiveReport struct {
	Target         string
	TargetType     TargetType
	ReverseDNS     string
	Location       LocationData
	Associated     []string
	SocialProfiles []SocialProfile
	Ports          []PortInfo
	SQLCheck       SQLExposure
	Timestamp      time.Time
	RiskScore      int
}
