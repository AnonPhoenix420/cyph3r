package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	ISP         string
	Org         string
	Country     string
	State       string
	City        string
	Zip         string
	Lat         string 
	Lon         string 
	NameServers map[string][]string
	MapLink     string 
}

type PhoneData struct {
	Number, Country, State, Location, Lat, Lon, Carrier, Type string
	Valid bool
	MapLink string
}
