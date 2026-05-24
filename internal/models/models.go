package models

import "time"

type TargetType int

const (
	TypeNetworkTarget TargetType = iota
	TypeEmailTarget
	TypePhoneTarget
	TypeGeoTarget
)

type GeoData struct {
	Latitude     string `json:"lat"`
	Longitude    string `json:"lon"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Timezone     string `json:"timezone"`
	MapReference string `json:"map_ref"`
}

type PhoneMetrics struct {
	Carrier   string `json:"carrier"`
	LineType  string `json:"line_type"`
	Location  string `json:"location"`
	IsActive  bool   `json:"is_active"`
	RiskScore int    `json:"risk_score"`
}

type IntelPayload struct {
	Target          string
	Type            TargetType
	ScanTime        time.Time
	Verbose         bool
	OutputFormat    string
	ASN             string
	ISP             string
	Geo             GeoData
	Clusters        []string
	OwnerName       string
	OwnerCountry    string
	CreatedDate     string
	OpenPorts       []string
	Banners         []string
	Vulnerabilities []string
	ExposedLeaks    []string
	Phone           PhoneMetrics
}
