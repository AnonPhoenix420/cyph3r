package models

type IntelData struct {
	IP          string
	Registrar   string
	ISP         string
	Org         string
	Country     string
	CountryCode string
	RegionName  string // State
	Zip         string
	City        string
	Lat         float64
	Lon         float64
	NameServers []string
	LocalHost   string
	LocalIPs    []string
}
