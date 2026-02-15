package intel

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var data models.PhoneData
	data.Number = number
	data.Valid = true

	clean := strings.TrimPrefix(number, "+")
	
	// Global OSINT Prefix Map (Expandable)
	globalMap := map[string][]string{
		"1330": {"USA", "Ohio", "Akron/Canton", "41.0814", "-81.5190", "Verizon/AT&T Cluster"},
		"1212": {"USA", "New York", "Manhattan", "40.7128", "-74.0060", "Verizon Cluster"},
		"44":   {"UK", "United Kingdom", "Global Hub", "55.3781", "-3.4360", "BT/Vodafone Architecture"},
		"91":   {"India", "India", "Global Hub", "20.5937", "78.9629", "Jio/Airtel Architecture"},
	}

	for prefix, info := range globalMap {
		if strings.HasPrefix(clean, prefix) {
			data.Country, data.State, data.Location, data.Lat, data.Lon, data.Carrier = info[0], info[1], info[2], info[3], info[4], info[5]
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
			break
		}
	}

	if data.Country == "" {
		data.Country, data.State, data.Location, data.Carrier = "International Node", "Scanning...", "Triangulating...", "Global Routing"
	}

	data.Type = "Mobile / VoIP High-Priority"
	return data, nil
}
