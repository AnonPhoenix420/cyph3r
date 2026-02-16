package intel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number = number

	// Using a multi-vector OSINT API for phone metadata
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=YOUR_KEY_OPTIONAL&number=%s", number)
	
	resp, err := client.Get(url)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	var res struct {
		Valid       bool   `json:"valid"`
		Carrier     string `json:"carrier"`
		Location    string `json:"location"`
		Type        string `json:"line_type"`
		CountryName string `json:"country_name"`
		Prefix      string `json:"country_prefix"`
	}
	json.NewDecoder(resp.Body).Decode(&res)

	// Mapping Intel to the Model
	d.Valid = res.Valid
	d.Carrier = res.Carrier
	d.Location = res.Location
	d.Country = res.CountryName
	d.Type = res.Type
	
	// Passive Risk Assessment: VOIP numbers are flagged as high risk
	if d.Type == "special_services" || d.Type == "toll_free" {
		d.Risk = "HIGH (Potential Burner)"
	} else {
		d.Risk = "LOW (Physical Asset)"
	}

	// Generating the Tactical Map Vector
	// In a real scenario, we'd pull Lat/Lon from a HLR lookup, 
	// here we simulate the vector link
	d.MapLink = fmt.Sprintf("https://www.google.com/maps/search/%s+%s", d.Location, d.Country)

	return d, nil
}
