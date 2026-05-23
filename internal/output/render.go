package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func Render(payload *models.IntelPayload) {
	if strings.ToLower(payload.OutputFormat) == "json" {
		renderJSON(payload)
		return
	}
	renderTerminalHUD(payload)
}

func renderJSON(payload *models.IntelPayload) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(payload); err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error rendering output payload to JSON format: %v\n", err)
	}
}

func renderTerminalHUD(p *models.IntelPayload) {
	// Dynamically pull colors straight from your pre-compiled internal/output/colors.go file
	fmt.Printf("%s[+] CYPH3R GHOST ELITE INTEL REPORT FOR: %s%s\n", NeonPink, p.Target, Reset)
	fmt.Printf("%s[-] TARGET VECTOR MATRIX CLASSIFICATION: %s%s\n", Gray, p.Type, Reset)
	fmt.Println(strings.Repeat("-", 63))

	switch p.Type {
	case models.TypePhoneTarget:
		drawBoxLine(fmt.Sprintf("VALIDITY : %s", fallback(p.Phone.Valid, "UNKNOWN")))
		drawBoxLine(fmt.Sprintf("FORMAT   : %s", fallback(p.Phone.LocalFormat, "N/A")))
		drawBoxLine(fmt.Sprintf("COUNTRY  : %s", fallback(p.Phone.CountryCode, "N/A")))
		drawBoxLine(fmt.Sprintf("LOCATION : %s", fallback(p.Phone.Location, "N/A")))
		drawBoxLine(fmt.Sprintf("CARRIER  : %s", fallback(p.Phone.Carrier, "N/A")))
		drawBoxLine(fmt.Sprintf("LINE TYPE: %s", fallback(p.Phone.LineType, "N/A")))

	case models.TypeEmailTarget:
		drawBoxLine(fmt.Sprintf("STEALTH STATUS: %s", fallback(p.Email.Deliverable, "UNKNOWN")))
		drawBoxLine(fmt.Sprintf("USER VECTOR   : %s", fallback(p.Email.Username, "N/A")))
		drawBoxLine(fmt.Sprintf("HOST ROUTE    : %s", fallback(p.Email.Domain, "N/A")))
		drawBoxLine(fmt.Sprintf("DISPOSABLE    : %s", fallback(p.Email.Disposable, "NO")))
		if p.Email.ProfileLink != "" {
			drawBoxLine(fmt.Sprintf("AVATAR TRACE  : %s", p.Email.ProfileLink))
		}
		if len(p.Email.MXRecords) > 0 {
			fmt.Println(strings.Repeat("-", 63))
			drawBoxLine("RESOLVED MX STEALTH PATHS:")
			for _, mx := range p.Email.MXRecords {
				drawBoxLine(fmt.Sprintf("  ↳ %s", mx))
			}
		}

	case models.TypeGeoTarget:
		drawBoxLine(fmt.Sprintf("LATITUDE  : %s", fallback(p.Geo.Latitude, "N/A")))
		drawBoxLine(fmt.Sprintf("LONGITUDE : %s", fallback(p.Geo.Longitude, "N/A")))
		drawBoxLine(fmt.Sprintf("GRID POINT: %s", fallback(p.Geo.City, "N/A")))
		drawBoxLine(fmt.Sprintf("ZONE RANGE: %s", fallback(p.Geo.Country, "N/A")))
		drawBoxLine(fmt.Sprintf("TIME TRACE: %s", fallback(p.Geo.Timezone, "N/A")))
		if p.Geo.MapReference != "" {
			drawBoxLine(fmt.Sprintf("MAP LINK  : %s", p.Geo.MapReference))
		}

	case models.TypeNetworkTarget:
		if p.ASN != "" || p.ISP != "" {
			drawBoxLine(fmt.Sprintf("ASN: %s", fallback(p.ASN, "N/A")))
			drawBoxLine(fmt.Sprintf("ISP: %s", fallback(p.ISP, "N/A")))
		}
		if p.Geo.Country != "" || p.Geo.City != "" {
			drawBoxLine(fmt.Sprintf("LOC: %s, %s", fallback(p.Geo.City, "Unknown City"), fallback(p.Geo.Country, "Unknown Country")))
		}
		if len(p.Clusters) > 0 {
			fmt.Println(strings.Repeat("-", 63))
			drawBoxLine("AUTHORITATIVE CLUSTERS:")
			for _, cluster := range p.Clusters {
				if strings.TrimSpace(cluster.NameServer) == "" {
					continue
				}
				drawBoxLine(fmt.Sprintf("  [-] %-20s", cluster.NameServer))
				if p.Verbose {
					for _, ip := range cluster.IPs {
						if strings.TrimSpace(ip) != "" {
							drawBoxLine(fmt.Sprintf("    ↳ %-22s [ONLINE]", ip))
						}
					}
				}
			}
		}
	}

	fmt.Println(strings.Repeat("-", 63))
}

func drawBoxLine(content string) {
	cleanText := content
	replacements := []string{NeonPink, Cyan, NeonGreen, Reset, Gray}
	for _, r := range replacements {
		cleanText = strings.ReplaceAll(cleanText, r, "")
	}

	visibleLength := len(cleanText)
	targetWidth := 61

	if visibleLength >= targetWidth {
		fmt.Printf("| %s |\n", content)
	} else {
		padding := targetWidth - visibleLength
		fmt.Printf("| %s%s |\n", content, strings.Repeat(" ", padding))
	}
}

func fallback(val, def string) string {
	if strings.TrimSpace(val) == "" {
		return def
	}
	return val
}
