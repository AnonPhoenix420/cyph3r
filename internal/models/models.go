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
	IsWAF       bool
	WAFType     string
	IsMobile    bool 
	IsProxy     bool
	IsHosting   bool
}

type GeoResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	Mobile      bool    `json:"mobile"`
	Proxy       bool    `json:"proxy"`
	Hosting     bool    `json:"hosting"`
}
