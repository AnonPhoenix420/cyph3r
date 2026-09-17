package intel

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

type PhoneMetrics struct {
	LineStatus     string `json:"line_status"`
	Carrier        string `json:"carrier"`
	Locale         string `json:"locale"`
	Risk           int    `json:"risk"`
	CountryCode    int    `json:"country_code"`
	NationalNumber uint64 `json:"national_number"`
	IsMobile       bool   `json:"is_mobile"`
}

func GetPhoneMetrics(phone string) PhoneMetrics {
	parsedNum, err := phonenumbers.Parse(phone, "US")
	if err != nil {
		return PhoneMetrics{
			LineStatus: "INVALID_OR_UNRESOLVED",
			Carrier:    "Unknown",
			Locale:     "Unknown",
			Risk:       85,
		}
	}

	isValid := phonenumbers.IsValidNumber(parsedNum)
	status := "ACTIVE_SUBSCRIBER_LINE"
	risk := 12
	if !isValid {
		status = "SUSPECT_OR_INACTIVE"
		risk = 65
	}

	region := phonenumbers.GetRegionCodeForNumber(parsedNum)
	numType := phonenumbers.GetNumberType(parsedNum)
	isMobile := numType == phonenumbers.MOBILE || numType == phonenumbers.FIXED_LINE_OR_MOBILE

	carrierName := "Global Mobile Network Routing"
	if region != "" {
		carrierName = fmt.Sprintf("Carrier Region: %s", region)
	}

	return PhoneMetrics{
		LineStatus:     status,
		Carrier:        carrierName,
		Locale:         region,
		Risk:           risk,
		CountryCode:    int(parsedNum.GetCountryCode()),
		NationalNumber: parsedNum.GetNationalNumber(),
		IsMobile:       isMobile,
	}
}
