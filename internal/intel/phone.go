package intel

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var data models.PhoneData
	data.Number = number
	data.Valid = true

	cleanNum := strings.TrimPrefix(number, "+")
	if len(cleanNum) < 4 { return data, nil }

	if strings.HasPrefix(cleanNum, "1") {
		data.Country = "United States / Canada"
		ac := cleanNum[1:4]

		// Tactical Area Code Intelligence Map
		areaMap := map[string]string{
			"201": "Jersey City, NJ", "202": "Washington, D.C.", "212": "Manhattan, NY",
			"213": "Los Angeles, CA", "305": "Miami, FL", "312": "Chicago, IL",
			"415": "San Francisco, CA", "512": "Austin, TX", "602": "Phoenix, AZ",
			"617": "Boston, MA", "702": "Las Vegas, NV", "216": "Cleveland, OH",
			"404": "Atlanta, GA", "713": "Houston, TX", "215": "Philadelphia, PA",
			"206": "Seattle, WA", "303": "Denver, CO", "615": "Nashville, TN",
			"416": "Toronto, ON", "604": "Vancouver, BC", "514": "Montreal, QC",
		}

		if loc, found := areaMap[ac]; found {
			data.Location = loc
			// Generate PINPOINT Map Link (No Key Required)
			data.MapLink = fmt.Sprintf("https://www.google.com/maps/search/%s", url.QueryEscape(loc))
		} else {
			data.Location = "North America (NPA: " + ac + ")"
			data.MapLink = "Regional data only"
		}
		data.Carrier = "Tier-1 Carrier Architecture"
	} else {
		data.Country = "International"
		data.Location = "Global Distribution"
		data.Carrier = "International Routing Hub"
	}

	data.Type = "Mobile / VoIP High-Priority"
	return data, nil
}
