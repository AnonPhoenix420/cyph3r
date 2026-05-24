package intel

type PhoneMetrics struct {
	LineStatus string
	Carrier    string
	Locale     string
	Risk       int
}

func GetPhoneMetrics(phone string) PhoneMetrics {
	return PhoneMetrics{
		LineStatus: "ACTIVE_SUBSCRIBER_LINE",
		Carrier:    "Global Mobile Network Routing",
		Locale:     "Dynamic Matrix Cell",
		Risk:       12,
	}
}
