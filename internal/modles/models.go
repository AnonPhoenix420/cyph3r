package models

// IntelData is the shared structure for the entire project
type IntelData struct {
	IPs         []string
	Nameservers []string
	Country     string
	CountryCode string
	Region      string
	City        string
	Zip         string
	Timezone    string
	ISP         string
	Org         string
	ASN         string
	Lat         float64
	Lon         float64
}
