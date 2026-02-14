package models

type IntelData struct {
	IP          string
	TargetName  string 
	TargetIPs   []string  // All resolved IPs for the target
	Org         string
	Country     string
	CountryCode string
	RegionName  string 
	Zip         string
	City        string
	Lat         float64
	Lon         float64
	NameServers map[string]string // Hostname -> IP
	PhoneNumbers string // All resolved phone numbers for the target
}
