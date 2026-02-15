package intel

import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var d models.IntelData
	d.TargetName = input
	// ... logic ...
	return d, nil
}

func GetPhoneIntel(number string) (models.PhoneData, error) {
	var d models.PhoneData
	d.Number = number
	// ... logic ...
	return d, nil
}
