package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type CrtShEntry struct {
	NameValue string `json:"name_value"`
}

type ExtractedIntel struct {
	RealIPs       []string
	Emails        []string
	PhoneNumbers  []string
	SocialHandles []string
	Subdomains    []string
}

// DiscoverOriginAndOSINT performs deep extraction across CT logs, DNS records, and asset fields
func DiscoverOriginAndOSINT(targetDomain string) ExtractedIntel {
	var intel ExtractedIntel
	subMap := make(map[string]bool)

	// 1. Query Certificate Transparency Logs
	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", targetDomain)
	client := &http.Client{Timeout: 8 * time.Second}
	resp, err := client.Get(url)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var entries []CrtShEntry
		if json.Unmarshal(body, &entries) == nil {
			for _, entry := range entries {
				for _, sub := range strings.Split(entry.NameValue, "\n") {
					sub = strings.TrimSpace(sub)
					if sub != "" && !strings.HasPrefix(sub, "*") && !subMap[sub] {
						subMap[sub] = true
						intel.Subdomains = append(intel.Subdomains, sub)
					}
				}
			}
		}
	}

	// 2. Resolve Subdomains and Filter out Cloudflare/CDN Edge Nodes to find Real IP
	ipMap := make(map[string]bool)
	for _, sub := range intel.Subdomains {
		ips, err := net.LookupIP(sub)
		if err == nil {
			for _, ip := range ips {
				ipStr := ip.String()
				// Filter out common Cloudflare/CDN blocks to isolate true backend origin
				if !strings.HasPrefix(ipStr, "104.") && 
				   !strings.HasPrefix(ipStr, "172.64.") && 
				   !strings.HasPrefix(ipStr, "173.245.") && 
				   !strings.HasPrefix(ipStr, "192.254.") && !ipMap[ipStr] {
					ipMap[ipStr] = true
					intel.RealIPs = append(intel.RealIPs, fmt.Sprintf("%s (%s)", ipStr, sub))
				}
			}
		}
	}

	// 3. Extract Deep OSINT Fields (Emails, Phones, Socials from TXT/MX/SPF records)
	txtRecords, _ := net.LookupTXT(targetDomain)
	rawText := strings.Join(txtRecordStrings(txtRecords), " ")
	
	// Also check SPF squash records if present
	spfSquash, err := net.LookupTXT("spfsquash." + targetDomain)
	if err == nil {
		rawText += " " + strings.Join(txtRecordStrings(spfSquash), " ")
	}

	intel.Emails = extractEmails(rawText)
	intel.PhoneNumbers = extractPhones(rawText)
	intel.SocialHandles = extractSocials(rawText)

	return intel
}

func txtRecordStrings(records []string) []string {
	return records
}

func extractEmails(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
}

func extractPhones(text string) []string {
	re := regexp.MustCompile(`(?:\+\d{1,3}\s?)?(?:\(\d{3}\)|\d{3})[-.\s]?\d{3}[-.\s]?\d{4}`)
	matches := re.FindAllString(text, -1)
	return uniqueStrings(matches)
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
