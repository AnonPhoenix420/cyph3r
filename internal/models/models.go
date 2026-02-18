package models

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
}

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
	Mobile      bool    `json:"mobile"`  // NEW: Cellular connection?
	Proxy       bool    `json:"proxy"`   // NEW: Is it a proxy/VPN?
	Hosting     bool    `json:"hosting"` // NEW: Is it a data center?
	Query       string  `json:"query"`
}
