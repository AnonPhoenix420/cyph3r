package intel

import (
	"fmt"
	"strings"
	"cyph3r/internal/models" // FIXED: Removed the github URL prefix
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var data models.PhoneData
	data.Number = number
	data.Valid = true

	clean := strings.TrimPrefix(number, "+")
	
	globalMap := map[string][]string{
		"1330": {"USA", "Ohio", "Akron/Canton", "41.0814", "-81.5190", "Verizon/AT&T"},
		"44":   {"UK", "United Kingdom", "London Hub", "51.5074", "-0.1278", "BT/Vodafone"},
	}

	for prefix, info := range globalMap {
		if strings.HasPrefix(clean, prefix) {
			data.Country, data.State, data.Location, data.Lat, data.Lon, data.Carrier = info[0], info[1], info[2], info[3], info[4], info[5]
			data.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", data.Lat, data.Lon)
			break
		}
	}

	if data.Country == "" {
		data.Country = "International"
	}

	return data, nil
}
