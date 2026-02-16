package intel

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number = number

	// Tactical Logic: Inference based on E.164 formatting
	if strings.HasPrefix(number, "+1") {
		d.Country = "USA/Canada"
		d.Valid = true
	} else if strings.HasPrefix(number, "+98") {
		d.Country = "Iran"
		d.Valid = true
	}

	// Risk Assessment Logic
	if !d.Valid {
		d.Risk = "HIGH (Unallocated/Virtual)"
		d.Type = "VOIP/Burner"
	} else {
		d.Risk = "LOW (Physical Asset)"
		d.Type = "Mobile"
	}

	d.MapLink = "https://www.google.com/maps/search/" + number
	return d, nil
}
