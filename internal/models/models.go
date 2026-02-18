package models

// IntelData aggregates all reconnaissance findings
type IntelData struct {
	TargetName  string
	TargetIPs   []string
	ReverseDNS  []string 
	Org         string
	City        string
	Region      string
	Country     string
	Lat         float64
	Lon         float64
	Latency     string 
	NameServers map[string][]string
	ScanResults []string
	RawGeo      string
	// Security Layers
	IsWAF       bool   
	WAFType     string 
}

// PhoneData stores OSINT results for mobile numbers
// (Restored to your full original structure)
type PhoneData struct {
	Number         string
	Carrier        string
	Country        string
	Risk           string
	HandleHint     string
	SocialPresence []string
	MapLink        string
}

// GeoResponse expanded to capture 100% of available API intel
type GeoResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Mobile      bool    `json:"mobile"`  
	Proxy       bool    `json:"proxy"`   
	Hosting     bool    `json:"hosting"` 
	Query       string  `json:"query"`
}
