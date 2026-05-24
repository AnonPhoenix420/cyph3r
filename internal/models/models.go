package models

import "time"

type TargetType string

const (
	TargetPhone  TargetType = "phone"
	TargetEmail  TargetType = "email"
	TargetIP     TargetType = "ip"
	TargetDomain TargetType = "domain"
)

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
	Port        int    `json:"port"`
	Service     string `json:"service"`
	State       string `json:"state"`
	Banner      string `json:"banner"`
	Vulnerable  bool   `json:"vulnerable"`
	AdminPort   bool   `json:"admin_port"`
}

type SQLExposure struct {
	Exposed   bool   `json:"exposed"`
	Ports     []int  `json:"ports"`
	RiskLevel string `json:"risk_level"`
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
