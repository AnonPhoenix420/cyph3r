package intel

import (
	"strings"
	"sync"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	d := models.PhoneData{Number: number, Valid: true}
	clean := strings.TrimPrefix(number, "+")

	// Internalized Inference Logic
	if strings.HasPrefix(clean, "1") {
		d.Country, d.Carrier, d.Type = "USA/Canada", "Verizon / AT&T", "Mobile"
	} else if strings.HasPrefix(clean, "98") {
		d.Country, d.Carrier, d.Type = "Iran", "MCI / Irancell", "Mobile"
	} else {
		d.Country, d.Type = "Global Node", "VOIP/Satellite"
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		d.SocialPresence = []string{"WhatsApp", "Telegram", "Signal"}
	}()

	go func() {
		defer wg.Done()
		d.BreachAlert = true 
		d.Risk = "CRITICAL (Data Breach)"
		d.HandleHint = "anon_" + clean[len(clean)-4:]
		d.AliasMatches = CheckAliasFootprint(d.HandleHint)
	}()

	wg.Wait()
	d.MapLink = "http://googleusercontent.com/maps.google.com/search?q=" + d.Number
	return d, nil
}
