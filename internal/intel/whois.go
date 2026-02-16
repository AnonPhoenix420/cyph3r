package intel

import (
	"github.com/likexian/whois"
	"strings"
)

func GetWhoisData(domain string) string {
	result, err := whois.Whois(domain)
	if err != nil {
		return "WHOIS_UNAVAILABLE"
	}

	// We'll extract just the key info so we don't dump a wall of text
	lines := strings.Split(result, "\n")
	var summary []string
	for _, line := range lines {
		if strings.Contains(line, "registrar:") || strings.Contains(line, "created:") || strings.Contains(line, "expire:") {
			summary = append(summary, strings.TrimSpace(line))
		}
	}
	
	if len(summary) == 0 { return "DATA_PROTECTED" }
	return strings.Join(summary, " | ")
}
