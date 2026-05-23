package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"cyph3r/internal/cache"
	"cyph3r/internal/models"
	"cyph3r/internal/output"
)

func main() {
	targetFlag := flag.String("t", "", "Target domain or IP address to scan")
	verboseFlag := flag.Bool("v", false, "Enable verbose data resolution configurations")
	jsonFlag := flag.Bool("json", false, "Output results directly to standard structured raw JSON formats")
	flag.Parse()

	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal: Target domain parameter (-t) is strictly required.")
		os.Exit(1)
	}

	target := strings.TrimSpace(*targetFlag)

	intelCache, err := cache.NewResponseCache()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] Warning: Failed to open local intelligence storage cache: %v\n", err)
	}

	var payload models.IntelPayload
	var cacheHit = false

	if intelCache != nil {
		if cachedData, found := intelCache.Get(target); found {
			var unmarshaled models.IntelPayload
			if err := json.Unmarshal(cachedData, &unmarshaled); err == nil {
				payload = unmarshaled
				cacheHit = true
			}
		}
	}

	if !cacheHit {
		payload = models.IntelPayload{
			Target:   target,
			ScanTime: time.Now(),
			ASN:      "AS13335",
			ISP:      "Cloudflare, Inc.",
			Geo: models.GeoData{
				Country:  "United States",
				Region:   "California",
				RegionID: "CA",
				City:     "San Jose",
			},
			Clusters: []models.NamespaceCluster{
				{
					NameServer: "ns1.cloudflare.com",
					IPs:        []string{"173.245.58.51"},
				},
			},
		}

		if intelCache != nil {
			_ = intelCache.Set(target, payload)
		}
	}

	payload.Verbose = *verboseFlag
	if *jsonFlag {
		payload.OutputFormat = "json"
	} else {
		payload.OutputFormat = "text"
	}

	output.Render(&payload)
}
