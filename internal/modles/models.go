package models

type IntelData struct {
	// ... existing fields ...
	IPs         []string
	Nameservers []string
	Country     string
	City        string
	Lat         float64
	Lon         float64
	ISP         string
	Org         string

	// New Field for WHOIS
	WhoisRaw    string 
}
