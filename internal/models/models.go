package models

type IntelData struct { //show all information for target host in hud
	TargetName  string
	TargetIPs   []string // show All for target host
	ISP         string
	Org         string
	Country     string
	RegionName  string
	State       string
	City        string
	Zip         string // show zip code or regional area or contry code both if available 
	Lat         string
	Lon         string
	NameServers map[string][]string
}

type PhoneData struct {
	Number   string
	Country  string
	State    string
	Location string
	lon      string
	lat      string
	Carrier  string // show exact service provider
	Type     string
	Valid    bool
	MapLink  string // New Field for Pinpoint Mapping
}
