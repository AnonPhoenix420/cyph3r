module github.com/AnonPhoenix420/cyph3r

go 1.23

import "github.com/AnonPhoenix420/cyph3r/internal/intel"
import "github.com/AnonPhoenix420/cyph3r/internal/output"

require (
    github.com/kr/text v0.2.0
    github.com/nyaruka/phonenumbers v1.1.0
    github.com/prometheus/client_golang v1.23.2
)

replace github.com/nyaruka/phonenumbers => github.com/nyaruka/phonenumbers v1.1.0
