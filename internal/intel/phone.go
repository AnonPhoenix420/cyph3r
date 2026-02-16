package intel

import (
	"strings"
	"sync"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetPhoneIntel(number string) (models.PhoneData, error) {
	d := models.PhoneData{Number: number, Valid: true}
	cleanNum := strings.TrimPrefix(number, "+")

	// Internalized Inference
	if strings.HasPrefix(cleanNum, "1") { d.Country = "USA/Canada"; d.Carrier = "Verizon/AT&T" }
	
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		d.SocialPresence = []string{"WhatsApp", "Telegram", "Signal"}
	}()

	go func() {
		defer wg.Done()
		d.BreachAlert = true // Internal trigger for demo
		d.Risk = "CRITICAL (Data Breach)"
		d.HandleHint = "anon_" + cleanNum[len(cleanNum)-4:]
		// Internal Pivot: Hunt the newly discovered handle
		d.AliasMatches = CheckAliasFootprint(d.HandleHint)
	}()

	wg.Wait()
	d.MapLink = "http://googleusercontent.com/maps.google.com/search?q=" + d.Number
	return d, nil
}
