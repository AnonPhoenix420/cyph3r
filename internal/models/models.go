package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	ISP         string
	Org         string
	Country     string
	RegionName  string
	State       string
	City        string
	Zip         string
	Lat         float64
	Lon         float64
	NameServers map[string][]string // Changed to []string
}

type PhoneData struct {
	Number   string
	Country  string
	Location string
	Carrier  string
	Type     string
	Valid    bool
}

