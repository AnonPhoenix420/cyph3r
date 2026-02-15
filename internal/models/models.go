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
	Lat         string 
	Lon         string 
	NameServers map[string][]string
	MapLink     string 
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
