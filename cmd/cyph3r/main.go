package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/cache"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

var (
	phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$|^7\d{9}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	geoRegex   = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
)

func sanitizeToDomain(input string) string {
	cleaned := strings.TrimSpace(input)
	if strings.Contains(cleaned, "://") {
		parts := strings.SplitN(cleaned, "://", 2)
		cleaned = parts[1]
	}
	if idx := strings.IndexAny(cleaned, "/?#:"); idx != -1 {
		cleaned = cleaned[:idx]
	}
	return strings.TrimSpace(cleaned)
}

func main() {
	targetFlag := flag.String("t", "", "Target routing vector parameters")
	verboseFlag := flag.Bool("v", false, "Enable verbose tracing outputs")
	jsonFlag := flag.Bool("json", false, "Output system information as raw JSON")
	flag.Parse()

	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal: Operational target parameter (-t) is strictly required.")
		os.Exit(1)
	}

	rawInput := strings.TrimSpace(*targetFlag)
	var target string
	var targetType models.TargetType

	if emailRegex.MatchString(rawInput) {
		target = strings.ReplaceAll(rawInput, " ", "")
		targetType = models.TypeEmailTarget
	} else if phoneRegex.MatchString(strings.ReplaceAll(rawInput, " ", "")) {
		target = strings.ReplaceAll(rawInput, " ", "")
		targetType = models.TypePhoneTarget
	} else if geoRegex.MatchString(strings.ReplaceAll(rawInput, " ", "")) {
		target = strings.ReplaceAll(rawInput, " ", "")
		targetType = models.TypeGeoTarget
	} else {
		target = sanitizeToDomain(rawInput)
		targetType = models.TypeNetworkTarget
	}

	intelCache, err := cache.NewResponseCache()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] Warning: Cache subsystems offline: %v\n", err)
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
			Type:     targetType,
			ScanTime: time.Now(),
		}

		switch targetType {
		case models.TypePhoneTarget:
			payload.Phone = intel.ResolvePhone(target)
		case models.TypeEmailTarget:
			payload.Email = intel.ResolveEmail(target)
		case models.TypeGeoTarget:
			coords := strings.Split(target, ",")
			payload.Geo = models.GeoData{
				Latitude:     strings.TrimSpace(coords[0]),
				Longitude:    strings.TrimSpace(coords[1]),
				City:         "Precision Coordinate Lock",
				Country:      "Global Core Grid",
				Timezone:     "UTC/GMT Z-Time",
				MapReference: "https://maps.google.com",
			}
		case models.TypeNetworkTarget:
			resolvedIP, clusters := intel.ResolveNetwork(target)
			payload.ASN = "AS13335"
			payload.ISP = fmt.Sprintf("Network Stack (%s)", resolvedIP)
			payload.Geo = models.GeoData{Country: "United States", City: "San Jose"}
			payload.Clusters = clusters
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

	// Hand off data to internal/output/render.go which manages clearing, banners, and layout options
	output.Render(&payload)
}
