package models

type IntelData struct {
	TargetName  string
	TargetIPs   []string
	Org         string
	City        string
	Region      string
	Country     string
	Lat         float64
	Lon         float64
	NameServers map[string][]string
	ScanResults []string
}

type PhoneData struct {
	Number         string
	Carrier        string
	Country        string
	Risk           string
	HandleHint     string
	SocialPresence []string
	MapLink        string
	Valid          bool
	BreachAlert    bool
}
