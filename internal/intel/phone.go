package intel

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func GetPhoneIntel(number string) {
	num, err := phonenumbers.Parse(number, "US")
	if err != nil {
		fmt.Println("Invalid number format")
		return
	}

	// Get the Carrier/Organization for free
	carrier := phonenumbers.GetCarrierForNumber(num, "en")
	if carrier == "" {
		carrier = "Unknown/Landline"
	}

	fmt.Printf("%s[+] PHONE:   %s%s%s\n", output.White, output.NeonGreen, number, output.Reset)
	fmt.Printf("%s[+] CARRIER: %s%s%s\n", output.White, output.NeonBlue, carrier, output.Reset)
}
