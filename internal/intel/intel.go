cat << 'EOF' > internal/intel/intel.go
package intel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/nyaruka/phonenumbers"
)

type CrtShEntry struct {
	NameValue string `json:"name_value"`
}

func DiscoverOriginAndOSINT(targetDomain string) models.ExtractedIntel {
	var intel models.ExtractedIntel
	subMap := make(map[string]bool)

	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", targetDomain)
	client := &http.Client{Timeout: 12 * time.Second}

	fmt.Printf("[*] Querying Certificate Transparency logs for %s...\n", targetDomain)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] Warning: CT log query failed: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[!] Warning: crt.sh returned HTTP status %d (Target may be rate-limiting or too large)\n", resp.StatusCode)
		} else {
			var entries []CrtShEntry
			if err := json.Unmarshal(body, &entries); err == nil {
				for _, entry := range entries {
					for _, sub := range strings.Split(entry.NameValue, "\n") {
						sub = strings.TrimSpace(sub)
						sub = strings.TrimPrefix(sub, "*.")
						if sub != "" && !subMap[sub] {
							subMap[sub] = true
							intel.Subdomains = append(intel.Subdomains, sub)
						}
					}
				}
				fmt.Printf("[+] Discovered %d unique subdomains from CT logs.\n", len(intel.Subdomains))
			} else {
				fmt.Printf("[!] Warning: Failed to parse CT JSON response: %v\n", err)
			}
		}
	}

	intel.FaviconHash = fetchFaviconHash(targetDomain, client)

	ipMap := make(map[string]bool)
	limit := len(intel.Subdomains)
	if limit > 100 {
		limit = 100
		fmt.Println("[*] Throttling DNS resolution to first 100 subdomains for speed.")
	}
	for i := 0; i < limit; i++ {
		sub := intel.Subdomains[i]
		ips, err := net.LookupIP(sub)
		if err == nil {
			for _, ip := range ips {
				ipStr := ip.String()
				if !isCDNEdge(ipStr) && !ipMap[ipStr] {
					ipMap[ipStr] = true
					intel.RealIPs = append(intel.RealIPs, fmt.Sprintf("%s (%s)", ipStr, sub))
				}
			}
		}
	}

	txtRecords, _ := net.LookupTXT(targetDomain)
	rawText := strings.Join(txtRecords, " ")

	intel.Emails = extractEmails(rawText)
	intel.PhoneNumbers = extractPhones(rawText)
	intel.SocialHandles = extractSocials(rawText)

	return intel
}

func fetchFaviconHash(domain string, client *http.Client) string {
	resp, err := client.Get(fmt.Sprintf("https://%s/favicon.ico", domain))
	if err != nil {
		resp, err = client.Get(fmt.Sprintf("http://%s/favicon.ico", domain))
		if err != nil {
			return "Unavailable"
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		return "Unavailable"
	}
	hash := md5.Sum(body)
	return hex.EncodeToString(hash[:])
}

func isCDNEdge(ip string) bool {
	return strings.HasPrefix(ip, "104.") ||
		strings.HasPrefix(ip, "172.64.") ||
		strings.HasPrefix(ip, "173.245.") ||
		strings.HasPrefix(ip, "192.254.") ||
		strings.HasPrefix(ip, "13.32.") ||
		strings.HasPrefix(ip, "151.101.")
}

func extractEmails(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
}

func extractPhones(text string) []string {
	re := regexp.MustCompile(`(?:\+\d{1,3}\s?)?(?:\(\d{3}\)|\d{3})[-.\s]?\d{3}[-.\s]?\d{4}`)
	matches := re.FindAllString(text, -1)
	var validatedPhones []string
	for _, m := range matches {
		num, err := phonenumbers.Parse(m, "US")
		if err == nil && phonenumbers.IsValidNumber(num) {
			validatedPhones = append(validatedPhones, phonenumbers.Format(num, phonenumbers.INTERNATIONAL))
		} else {
			validatedPhones = append(validatedPhones, m)
		}
	}
	return uniqueStrings(validatedPhones)
}

func extractSocials(text string) []string {
	re := regexp.MustCompile(`(?:twitter\.com|x\.com|linkedin\.com|github\.com|facebook\.com)/[a-zA-Z0-9_.-]+`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
}

func uniqueStrings(input []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range input {
		if !keys[entry] {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
EOF
