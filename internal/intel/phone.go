package intel

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number, d.Valid = number, true
	clean := strings.TrimPrefix(number, "+")
	
	// Mock DB for triangulation
	globalMap := map[string][]string{
		"1330": {"USA", "Ohio", "Akron/Canton", "41.0814", "-81.5190", "Verizon/AT&T"},
		"44":   {"UK", "United Kingdom", "London Hub", "51.5074", "-0.1278", "BT/Vodafone"},
	}

	for pfx, info := range globalMap {
		if strings.HasPrefix(clean, pfx) {
			d.Country, d.State, d.Location, d.Lat, d.Lon, d.Carrier = info[0], info[1], info[2], info[3], info[4], info[5]
			d.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", d.Lat, d.Lon)
			break
		}
	}
	return d, nil
}
