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
	// Using the color constants already declared inside your colors.go file
	fmt.Printf("%s[+] CYPH3R GHOST ELITE INTEL REPORT FOR: %s%s\n", NeonPink, p.Target, Reset)
	fmt.Println(strings.Repeat("-", 63))

	if p.ASN != "" || p.ISP != "" {
		drawBoxLine(fmt.Sprintf("ASN: %s", fallback(p.ASN, "N/A")))
		drawBoxLine(fmt.Sprintf("ISP: %s", fallback(p.ISP, "N/A")))
	}
	
	if p.Geo.Country != "" || p.Geo.City != "" {
		drawBoxLine(fmt.Sprintf("LOC: %s, %s", fallback(p.Geo.City, "Unknown City"), fallback(p.Geo.Country, "Unknown Country")))
	}
	
	if p.Geo.Timezone != "" {
		drawBoxLine(fmt.Sprintf("TZ : %s", p.Geo.Timezone))
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
