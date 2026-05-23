package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"cyph3r/internal/models"
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
	fmt.Printf("%s[+] CYPH3R INTELLIGENCE REPORT FOR: %s%s\n", NeonPink, p.Target, Reset)
	fmt.Println(strings.Repeat("-", 63))

	if p.ASN != "" || p.ISP != "" {
		drawBoxLine(fmt.Sprintf("ASN: %s", fallback(p.ASN, "N/A")))
		drawBoxLine(fmt.Sprintf("ISP: %s", fallback(p.ISP, "N/A")))
	}

	if p.Geo.Country != "" || p.Geo.City != "" || p.Geo.RegionID != "" {
		geoString := fmt.Sprintf("LOC: %s, %s", fallback(p.Geo.City, "Unknown City"), fallback(p.Geo.Country, "Unknown Country"))
		if p.Geo.RegionID != "" {
			geoString += fmt.Sprintf(" (Region ID: %s)", p.Geo.RegionID)
		}
		drawBoxLine(geoString)
	}

	if len(p.Clusters) > 0 {
		fmt.Println(strings.Repeat("-", 63))
		drawBoxLine("AUTHORITATIVE CLUSTERS:")
		
		for _, cluster := range p.Clusters {
			if strings.TrimSpace(cluster.NameServer) == "" {
				continue
			}
			
			nsLine := fmt.Sprintf("  [-] %-20s", cluster.NameServer)
			drawBoxLine(nsLine)

			if p.Verbose {
				for _, ip := range cluster.IPs {
					if strings.TrimSpace(ip) == "" {
						continue
					}
					ipLine := fmt.Sprintf("    ↳ %-22s [ONLINE]", ip)
					drawBoxLine(ipLine)
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
