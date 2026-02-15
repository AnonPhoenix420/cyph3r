package models

// IntelData handles remote target network recon
type IntelData struct {
	IP          string
	TargetName  string
	TargetIPs   []string
	ISP         string   // This was the missing field
	Org         string
	Country     string
	CountryCode string
	RegionName  string
	Zip         string
	City        string
	Lat         float64
	Lon         float64
	NameServers map[string]string
}

// PhoneData handles international phone metadata
type PhoneData struct {
	Number      string `json:"number"`
	Valid       bool   `json:"valid"`
	LocalFormat string `json:"local_format"`
	Carrier     string `json:"carrier"`
	Location    string `json:"location"`
	Type        string `json:"line_type"`
}

