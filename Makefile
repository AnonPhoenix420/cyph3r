mkdir -p internal/models internal/intel internal/output internal/probes internal/stress cmd/cyph3r

# 1. Models
cat << 'EOF' > internal/models/models.go
package models

type ExtractedIntel struct {
	Subdomains    []string
	RealIPs       []string
	FaviconHash   string
	Emails        []string
	PhoneNumbers  []string
	SocialHandles []string
}
EOF

# 2. Output & Banner
cat << 'EOF' > internal/output/output.go
package output

import "fmt"

func PrintBanner() {
	fmt.Println(`
╔═══════════════════════════════════════════════════════════════╗
║                   CYPH3R v2.6 OSINT SUITE                     ║
╚═══════════════════════════════════════════════════════════════╝
    `)
}
EOF

# 3. Probes
cat << 'EOF' > internal/probes/probes.go
package probes

import (
	"fmt"
	"net"
	"time"
)

func ExecuteContinuousMonitor(targetAddr, proto string, interval time.Duration) {
	fmt.Printf("[*] Engaging live HUD connection monitor for %s (%s) [Interval: %v]\n", targetAddr, proto, interval)
	for {
		conn, err := net.DialTimeout(proto, targetAddr, 3*time.Second)
		timestamp := time.Now().Format("15:04:05")
		if err != nil {
			fmt.Printf("[%s] [DOWN] Connection failed: %v\n", timestamp, err)
		} else {
			fmt.Printf("[%s] [UP] Target responsive and accepting connections\n", timestamp)
			conn.Close()
		}
		time.Sleep(interval)
	}
}
EOF

# 4. Stress Engine
cat << 'EOF' > internal/stress/stress.go
package stress

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ExecuteContinuousStress(targetAddr string, concurrency int, duration int) {
	fmt.Printf("[*] Launching stress engine against %s with pool size %d\n", targetAddr, concurrency)
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	if duration > 0 {
		time.AfterFunc(time.Duration(duration)*time.Second, func() {
			close(stopChan)
		})
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					conn, err := net.DialTimeout("tcp", targetAddr, 2*time.Second)
					if err == nil {
						conn.Close()
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("[+] Stress test cycle completed.")
}
EOF

# 5. Intel Engine
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

# 6. Main Entry Point
cat << 'EOF' > cmd/cyph3r/main.go
package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
	"github.com/AnonPhoenix420/cyph3r/internal/stress"
)

func main() {
	targetFlag := flag.String("target", "", "Target domain or IP address")
	portFlag := flag.Int("p", 443, "Target port")
	protoFlag := flag.String("proto", "tcp", "Wire protocol (tcp/udp)")
	monitorFlag := flag.Bool("monitor", false, "Engage live HUD connection monitor loop")
	hulkFlag := flag.Bool("hulk", false, "Engage continuous resilience stress engine")
	osintFlag := flag.Bool("osint", false, "Extract unmasked origin IPs, Favicon hashes, emails, phones, and social footprints")
	concurrencyFlag := flag.Int("c", 1000, "Concurrency pool size for stress tests")
	durationFlag := flag.Int("d", 0, "Test duration in seconds (0 for infinite loop)")
	intervalFlag := flag.Duration("interval", 2*time.Second, "HUD monitor ping interval")

	flag.Parse()

	if *targetFlag == "" {
		output.PrintBanner()
		fmt.Println("[!] Error: --target parameter is required.")
		return
	}

	cleanHost := strings.TrimPrefix(strings.TrimPrefix(*targetFlag, "https://"), "http://")
	if idx := strings.Index(cleanHost, "/"); idx != -1 {
		cleanHost = cleanHost[:idx]
	}
	if idx := strings.Index(cleanHost, ":"); idx != -1 {
		cleanHost = cleanHost[:idx]
	}

	targetAddr := fmt.Sprintf("%s:%d", cleanHost, *portFlag)

	if *osintFlag {
		output.PrintBanner()
		fmt.Printf("[+] LAUNCHING DEEP OSINT & PROXY-BYPASS INTEL SCAN: %s\n", cleanHost)
		results := intel.DiscoverOriginAndOSINT(cleanHost)

		fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║               CYPH3R DEEP OSINT FIELD INTELLIGENCE          ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

		fmt.Println("\n[ UNMASKED REAL ORIGIN NODES ]")
		if len(results.RealIPs) > 0 {
			for _, ip := range results.RealIPs {
				fmt.Printf("  ↳ %s\n", ip)
			}
		} else {
			fmt.Println("  ↳ No unmasked origin IPs detected (Strict CDN edge encapsulation active).")
		}

		fmt.Printf("\n[ FAVICON FINGERPRINT HASH (MD5) ]\n  ↳ %s\n", results.FaviconHash)

		fmt.Println("\n[ HARVESTED EMAILS ]")
		if len(results.Emails) > 0 {
			for _, email := range results.Emails {
				fmt.Printf("  ↳ %s\n", email)
			}
		} else {
			fmt.Println("  ↳ None exposed in public records.")
		}

		fmt.Println("\n[ EXTRACTED PHONE VECTORS ]")
		if len(results.PhoneNumbers) > 0 {
			for _, phone := range results.PhoneNumbers {
				fmt.Printf("  ↳ %s\n", phone)
			}
		} else {
			fmt.Println("  ↳ None detected.")
		}

		fmt.Println("\n[ LEAKED SOCIAL MEDIA REFERENCES ]")
		if len(results.SocialHandles) > 0 {
			for _, soc := range results.SocialHandles {
				fmt.Printf("  ↳ https://%s\n", soc)
			}
		} else {
			fmt.Println("  ↳ None mapped.")
		}
		return
	}

	if *hulkFlag {
		stress.ExecuteContinuousStress(targetAddr, *concurrencyFlag, *durationFlag)
		return
	}

	if *monitorFlag {
		output.PrintBanner()
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), *intervalFlag)
		return
	}

	output.PrintBanner()
	fmt.Printf("[+] Target specified: %s. Use --osint, --monitor, or --hulk to engage modules.\n", targetAddr)
}
EOF

# 7. Makefile with literal tabs guaranteed via printf
printf '# ─── CYPH3R v2.6 SYSTEM MAINTENANCE MAKEFILE ──────────────────────────\n\nBINARY_NAME=cyph3r\n\nall: build\n\nfix-imports:\n\t@echo "[*] Normalizing module import paths..."\n\t@find . -type f -name '\''*.go'\'' -exec sed -i '\''s|"cyph3r/internal|"github.com/AnonPhoenix420/cyph3r/internal|g'\'' {} +\n\nbuild: fix-imports\n\t@echo "[*] Syncing dependencies..."\n\t@go mod tidy\n\t@echo "[*] Building CYPH3R v2.6 production binary..."\n\tgo build -o $(BINARY_NAME) ./cmd/cyph3r\n\t@echo "[✓] Build complete."\n\nrepair: fix-imports\n\t@echo "[*] Initializing CYPH3R Self-Repair Routine..."\n\tgo clean -modcache\n\tgo mod tidy\n\tgo build -o $(BINARY_NAME) ./cmd/cyph3r\n\t@echo "[✓] Environment successfully repaired and resynced."\n\nclean:\n\t@echo "[*] Removing build artifacts and binaries..."\n\trm -f $(BINARY_NAME)\n\tgo clean\n\t@echo "[✓] Project environment cleaned."\n\n.PHONY: all build repair clean fix-imports\n' > Makefile
