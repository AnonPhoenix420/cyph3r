package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	ISP         string
	Org         string
	Country     string
	RegionName  string
	State       string // Added
	City        string
	Zip         string
	Lat         string // String format
	Lon         string // String format
	MapLink     string
	NameServers map[string][]string
}

type PhoneData struct {
	Number   string
	Country  string
	State    string
	Location string
	Lat      string
	Lon      string
	Carrier  string
	Type     string
	Valid    bool
	MapLink  string
}
