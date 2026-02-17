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
	Latency     string
	NameServers map[string][]string
	ScanResults []string
	RawGeo      string // <--- Added for Verbose Mode (-v)
}

type PhoneData struct {
	Number         string
	Carrier        string
	Country        string
	Risk           string
	HandleHint     string
	SocialPresence []string
	MapLink        string
}
