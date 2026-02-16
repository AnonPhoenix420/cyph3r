package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	City, State, Country, Org string
	NameServers map[string][]string // NS Name -> [IPs]
	Subdomains  []string
	ScanResults []string
	MapLink string
}

type PhoneData struct {
	Number         string
	Valid          bool
	Carrier        string
	Type           string
	Location       string
	Country        string
	Risk           string
	MapLink        string
	SocialPresence []string
	BreachAlert    bool
	HandleHint     string
	AliasMatches   []string
}
