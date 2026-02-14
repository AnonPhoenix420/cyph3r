package intel

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

// GetPhoneIntel extracts carrier and region data
func GetPhoneIntel(number string) {
	// Defaulting to "US" if no country code provided, but it handles international prefixes
	num, err := phonenumbers.Parse(number, "US")
	if err != nil {
		output.Error("Invalid phone number format")
		return
	}

	carrier := phonenumbers.GetCarrierForNumber(num, "en")
	if carrier == "" {
		carrier = "Unknown / Landline"
	}

	region := phonenumbers.GetRegionCodeForNumber(num)
	formatted := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)

	fmt.Printf("\n%s--- [ PHONE_INTELLIGENCE ] ---%s\n", output.NeonPink, output.Reset)
	fmt.Printf("%s[*] Number:   %s%s\n", output.White, output.NeonGreen, formatted)
	fmt.Printf("%s[*] Region:   %s%s\n", output.White, output.NeonGreen, region)
	fmt.Printf("%s[*] Carrier:  %s%s%s\n", output.White, output.NeonBlue, carrier, output.Reset)
}
