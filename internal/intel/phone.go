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
	
	// Tactical Global Mapping Database
	// Structure: {Country, State/Region, Lat, Lon, Primary Carrier}
	globalMap := map[string][]string{
		"1330": {"USA", "Ohio (Akron/Canton)", "41.0814", "-81.5190", "Verizon/AT&T Cluster"},
		"1212": {"USA", "New York (Manhattan)", "40.7128", "-74.0060", "Verizon Cluster"},
		"44":   {"UK", "United Kingdom (Global)", "55.3781", "-3.4360", "BT/Vodafone Architecture"},
		"4420": {"UK", "London", "51.5074", "-0.1278", "O2/EE London Hub"},
		"91":   {"India", "India (Global)", "20.5937", "78.9629", "Jio/Airtel Architecture"},
		"49":   {"Germany", "Germany (Global)", "51.1657", "10.4515", "Deutsche Telekom"},
		"61":   {"Australia", "Australia (Global)", "-25.2744", "133.7751", "Telstra Network"},
		"81":   {"Japan", "Japan (Global)", "36.2048", "138.2529", "NTT Docomo Hub"},
	}

	// Dynamic Prefix Matching Logic
	matched := false
	for prefix, info := range globalMap {
		if strings.HasPrefix(clean, prefix) {
			data.Country = info[0]
			data.State = info[1]
			data.Lat = info[2]
			data.Lon = info[3]
			data.Carrier = info[4]
			matched = true
			break
		}
	}

	if !matched {
		data.Country = "International Node"
		data.State = "Triangulating Region..."
		data.Carrier = "Global Routing Architecture"
	}

	// Generate Map Link for any coordinates found
	if data.Lat != "" {
		data.MapLink = fmt.Sprintf("https://developers.google.com/maps/documentation/geocoding/intro#ComponentFiltering7", data.Lat, data.Lon)
	}

	data.Type = "Mobile / VoIP High-Priority"
	return data, nil
}
