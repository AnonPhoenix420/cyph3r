package models

type TacticalConfig struct {
	Target string
	Vector string
	PPS    int
	Port   string
	Power  int
}

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	TargetIPv6s []string
	ReverseDNS  []string
	Org         string
	ISP         string 
	AS          string 
	City        string
	Region      string
	RegionName  string 
	Country     string
	CountryCode string 
	Zip         string 
	Timezone    string 
	Lat         float64
	Lon         float64
	NameServers map[string][]string
	ScanResults []string
	IsWAF       bool
	WAFType     string
	IsHosting   bool
}

type GeoResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Hosting     bool    `json:"hosting"`
}
