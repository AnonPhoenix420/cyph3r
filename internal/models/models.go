package models

type IntelData struct {
	IPs         []string
	Nameservers []string
	WhoisRaw    string
	Registrar   string
	Country     string
	City        string
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
}
