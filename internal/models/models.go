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
	Latency     string // Measures RTT pulse
	NameServers map[string][]string
	ScanResults []string
	RawGeo      string // Full JSON dump for Verbose Mode
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
