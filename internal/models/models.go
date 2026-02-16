package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	City        string
	State       string
	Country     string
	Zip         string
	ISP         string
	Org         string
	NameServers map[string][]string
	Subdomains  []string
}

type PhoneData struct {
	Number   string
	Carrier  string
	Location string
	Country  string
	MapLink  string
}
