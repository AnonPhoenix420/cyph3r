package models

type IntelData struct {
	IPs         []string
	Nameservers []string
	WhoisRaw    string
	Registrar   string
	CreatedDate string
	ExpiryDate  string
	Country     string
	CountryCode string
	Region      string
	City        string
	Zip         string
	Timezone    string
	Lat         float64
	Lon         float64
	ISP         string
	Org         string
	ASN         string
	IsAlive     bool
}

type ProbeResult struct {
	Port     int
	State    string
	Service  string
	Protocol string
}
