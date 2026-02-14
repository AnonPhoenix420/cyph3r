package intel

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

// GetPhoneIntel extracts carrier and region data
func GetPhoneIntel(number string) {
	// Parse the number
	num, err := phonenumbers.Parse(number, "US")
	if err != nil {
		output.Error("Invalid phone number format")
		return
	}

	// FIX: Added ', _' to handle the second return value (usually an error/status)
	carrier, _ := phonenumbers.GetCarrierForNumber(num, "en")

	region := phonenumbers.GetRegionCodeForNumber(num)
	formatted := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)

	fmt.Printf("\n%s--- [ PHONE_INTELLIGENCE ] ---%s\n", output.NeonPink, output.Reset)
	fmt.Printf("%s[*] Number:   %s%s\n", output.White, output.NeonGreen, formatted)
	fmt.Printf("%s[*] Region:   %s%s\n", output.White, output.NeonGreen, region)
	
	if carrier == "" {
		carrier = "Unknown / Landline"
	}
	fmt.Printf("%s[*] Carrier:  %s%s%s\n", output.White, output.NeonBlue, carrier, output.Reset)
}
