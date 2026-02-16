package models

// IntelData handles Domain/IP intelligence
type IntelData struct {
	TargetName  string
	TargetIPs   []string
	City, State, Country, Zip, ISP, Org string
	NameServers map[string][]string
	Subdomains  []string
}

// PhoneData handles Cellular/Satellite intelligence
type PhoneData struct {
	Number   string
	Valid    bool
	Carrier  string
	Type     string // Mobile, Landline, VOIP
	Location string
	State    string
	Country  string
	Lat      string
	Lon      string
	Timezone string
	Risk     string // High/Low based on VOIP detection
	MapLink  string
}
