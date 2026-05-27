package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/cache"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
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
	targetFlag := flag.String("target", "", "Target input node routing configuration vector")
	phoneFlag := flag.String("phone", "", "Execute standalone telephony metadata lookup")
	
	delayFlag := flag.String("delay", "0s", "Introduce spacing delays between validation packets")
	agentFlag := flag.String("agent", "", "Override network footprint with a custom client signature")
	methodFlag := flag.String("method", "GET", "HTTP verb operation configuration parameter (GET/POST)")

	runTestFlag := flag.Bool("test-integrity", false, "Engage Elite Network Systems Testing suite")
	testModeFlag := flag.Int("mode", 1, "Select verification model: 1=LOAD, 2=STRESS, 3=SOAK, 4=SPIKE")
	concurrencyFlag := flag.Int("c", 50, "Simultaneous validation connection streams")
	durationFlag := flag.Int("d", 10, "Testing matrix window duration parameter in seconds")

	monitorFlag := flag.Bool("monitor", false, "Engage continuous HUD monitor loop execution")
	protoFlag := flag.String("proto", "tcp", "Protocol mode selector for telemetry checking loops")
	intervalFlag := flag.String("interval", "2s", "Telemetry tracking update frequency window interval")
	
	jsonFlag := flag.Bool("json", false, "Format final target layout output structure as raw JSON matrix")
	verboseFlag := flag.Bool("v", false, "Enable full logging debug tracing variables")
	
	flag.Parse()

	rawInput := strings.TrimSpace(*targetFlag)
	if rawInput == "" && *phoneFlag != "" {
		rawInput = strings.TrimSpace(*phoneFlag)
	}

	if rawInput == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal Parameter Error: An operational target identifier mapping (--target) is strictly required.")
		os.Exit(1)
	}

	if *monitorFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		interval, _ := time.ParseDuration(*intervalFlag)
		if interval == 0 {
			interval = 2 * time.Second
		}
		probes.ExecuteContinuousMonitor(rawInput, strings.ToLower(*protoFlag), interval)
		return
	}

	if *runTestFlag {
		targetURL := rawInput
		if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
			targetURL = "http://" + targetURL
		}
		fmt.Print(output.ClearLine)
		output.Banner()
		intel.ExecuteValidationSuite(targetURL, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

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

	intelCache, _ := cache.NewResponseCache()
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

		var socialTracks []string
		if targetType != models.TypeGeoTarget {
			foundProfiles := intel.ResolveSocialFootprint(target)
			for _, profile := range foundProfiles {
				socialTracks = append(socialTracks, fmt.Sprintf("[%s Nodes] ➔ %s", profile.Platform, profile.ProfileURL))
			}
		}

		switch targetType {
		case models.TypeEmailTarget:
			payload.ISP = "Enterprise Mail MX Architecture"
			payload.ExposedLeaks = append(intel.CheckThreatFeeds(target), intel.ResolveEmail(target))
			payload.Vulnerabilities = socialTracks
			payload.Clusters = []string{"IDENTITY_VERIFIED", "GLOBAL_CROSS_REFERENCE_ENGAGED"}

		case models.TypePhoneTarget:
			alloc, provider, zone := intel.ResolvePhone(target)
			payload.Phone = target
			payload.ISP = provider
			payload.OwnerName = alloc
			payload.CreatedDate = "TELEPHONY_RECORD_LIVE"
			payload.ExposedLeaks = append(intel.CheckThreatFeeds(target), fmt.Sprintf("Zone: %s", zone))
			payload.Vulnerabilities = socialTracks
			payload.Clusters = []string{"TELEPHONY_INTELLIGENCE_NODE", "E164_ALIGNED"}

		case models.TypeGeoTarget:
			coords := strings.Split(target, ",")
			payload.ISP = "Satellite Mapping Coordinate Alignment"
			payload.Geo = models.GeoData{
				Latitude:  strings.TrimSpace(coords[0]),
				Longitude: strings.TrimSpace(coords[1]),
				City:      "Precision Grid Intercept",
				Country:   "Geocentric Anchor Point Cluster",
			}
			payload.Clusters = []string{"GEO_ANCHOR_VALIDATED"}

		case models.TypeNetworkTarget:
			parsedDelay, _ := time.ParseDuration(*delayFlag)
			resIP, geo, asn, owner, date, ports, banners, vulns, leaks, sqlCheck := intel.ResolveNetworkElite(target, parsedDelay, *agentFlag)
			
			payload.ASN = asn
			payload.ISP = fmt.Sprintf("Network Interface Stack (%s)", resIP)
			payload.Geo = geo
			payload.OwnerName = owner
			payload.CreatedDate = date
			payload.ExposedLeaks = append(leaks, intel.CheckThreatFeeds(target)...)
			payload.OpenPorts = ports
			payload.Banners = banners
			payload.Vulnerabilities = append(vulns, socialTracks...)
			
			if sqlCheck.Exposed {
				payload.Clusters = append(payload.Clusters, fmt.Sprintf("SQL_EXPOSED_RISK_%s", sqlCheck.RiskLevel))
			}
			payload.Clusters = append(payload.Clusters, "LIVE_NODE_CONNECTED")

			payload.HTTPMethod = strings.ToUpper(*methodFlag)
			urlStr := "http://" + target
			if req, err := http.NewRequest(payload.HTTPMethod, urlStr, nil); err == nil {
				if *agentFlag != "" {
					req.Header.Set("User-Agent", *agentFlag)
				} else {
					req.Header.Set("User-Agent", "CYPH3R/Master-Engine-2026")
				}
				payload.CapturedHeaders = req.Header
			}
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

	fmt.Print(output.ClearLine)
	output.Banner()
	output.Render(&payload)
}
