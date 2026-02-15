package intel

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number = number
	d.Valid = true
	
	// Strip symbols for prefix checking
	clean := strings.TrimPrefix(number, "+")
	
	// Tactical Prefix Mapping
	// In a production environment, you'd call a dedicated API here
	globalMap := map[string][]string{
		"1":    {"USA/Canada", "North America", "Global", "37.0902", "-95.7129", "Multi-Carrier"},
		"44":   {"UK", "United Kingdom", "London Hub", "51.5074", "-0.1278", "BT/Vodafone"},
		"1330": {"USA", "Ohio", "Akron/Canton", "41.0814", "-81.5190", "Verizon/AT&T"},
	}

	for pfx, info := range globalMap {
		if strings.HasPrefix(clean, pfx) {
			d.Country = info[0]
			d.State = info[1]
			d.Location = info[2]
			d.Lat = info[3]
			d.Lon = info[4]
			d.Carrier = info[5]
			// Real coordinate link
			d.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", d.Lat, d.Lon)
			break
		}
	}

	// Default if no prefix matched
	if d.Country == "" {
		d.Country = "Unknown"
		d.Carrier = "Unknown/VOIP"
	}

	return d, nil
}
