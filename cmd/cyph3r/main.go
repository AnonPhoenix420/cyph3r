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
	// Traditional CYPH3R v2.6 long-flag mappings
	targetFlag := flag.String("target", "", "Target routing domain, email, or infrastructure IP vector")
	phoneFlag := flag.String("phone", "", "Execute international telephone vector metadata lookup")
	scanActiveFlag := flag.Bool("scan", false, "Execute tactical concurrent port scan and banner analysis")
	monitorFlag := flag.Bool("monitor", false, "Engage continuous HUD telemetry monitoring mode loop")
	protoFlag := flag.String("proto", "tcp", "Protocol mode selector for tracing [tcp, udp, http, https, ack]")
	intervalFlag := flag.String("interval", "2s", "Telemetry delay frequency window interval")
	
	// Stress Validation Infrastructure Mappings
	runTestFlag := flag.Bool("test-integrity", false, "Execute integrated validation stress suite")
	testModeFlag := flag.Int("mode", 1, "Select test vector: 1=LOAD, 2=STRESS, 3=SOAK, 4=SPIKE")
	concurrencyFlag := flag.Int("c", 50, "Number of concurrent verification testing streams")
	durationFlag := flag.Int("d", 10, "Duration of verification stress trace parameters in seconds")
	
	verboseFlag := flag.Bool("v", false, "Enable full operational tracing logs")
	jsonFlag := flag.Bool("json", false, "Output data structure as raw JSON matrix")
	
	flag.Parse()

	// 1. Direct Telephony Flag Override Shortcut
	if *phoneFlag != "" {
		output.ClearLine()
		output.Banner()
		payload := models.IntelPayload{
			Target:   strings.ReplaceAll(*phoneFlag, " ", ""),
			Type:     models.TypePhoneTarget,
			ScanTime: time.Now(),
			Phone:    intel.ResolvePhone(*phoneFlag),
		}
		output.Render(&payload)
		return
	}

	// 2. Structural Guardrail Validation Check
	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal: Operational parameter target mapping (--target or --phone) strictly required.")
		os.Exit(1)
	}

	rawInput := strings.TrimSpace(*targetFlag)

	// 3. Persistent HUD Telemetry Monitoring Route
	if *monitorFlag {
		output.ClearLine()
		output.Banner()
		interval, err := time.ParseDuration(*intervalFlag)
		if err != nil {
			interval = 2 * time.Second
		}
		probes.ExecuteContinuousMonitor(rawInput, strings.ToLower(*protoFlag), interval)
		return
	}

	// 4. System Stress Validation Suite Intercept Route
	if *runTestFlag {
		targetURL := rawInput
		if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
			targetURL = "http://" + targetURL
		}
		output.ClearLine()
		output.Banner()
		intel.ExecuteValidationSuite(targetURL, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

	// 5. Input Target Evaluation Tree
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

	// 6. Execution Cache Layers
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

	// 7. Core Threat Processing Logic
	if !cacheHit {
		payload = models.IntelPayload{
			Target:   target,
			Type:     targetType,
			ScanTime: time.Now(),
		}

		switch targetType {
		case models.TypeEmailTarget:
			// Resolve the profile metadata string
			avatarPtr := intel.ResolveEmail(target)
			payload.OwnerName = avatarPtr // store fallback tracking signature
			payload.Clusters = []string{"IDENTITY_VERIFIED"}

		case models.TypePhoneTarget:
			payload.Phone = intel.ResolvePhone(target)

		case models.TypeGeoTarget:
			coords := strings.Split(target, ",")
			payload.Geo = models.GeoData{
				Latitude:     strings.TrimSpace(coords[0]),
				Longitude:    strings.TrimSpace(coords[1]),
				City:         "Precision Grid Intercept",
				Country:      "Localized Anchor Node",
				Timezone:     "UTC/GMT Z-Time",
				MapReference: fmt.Sprintf("https://maps.google.com/?q=%s,%s", strings.TrimSpace(coords[0]), strings.TrimSpace(coords[1])),
			}

		case models.TypeNetworkTarget:
			resIP, geo, asn, owner, date, ports, banners, vulns, leaks := intel.ResolveNetwork(target)
			payload.ASN = asn
			payload.ISP = fmt.Sprintf("Network Stack (%s)", resIP)
			payload.Geo = geo
			payload.OwnerName = owner
			payload.CreatedDate = date
			
			if *scanActiveFlag {
				payload.OpenPorts = ports
				payload.Banners = banners
				payload.Vulnerabilities = vulns
				payload.ExposedLeaks = leaks
			}
			payload.Clusters = []string{"LIVE_NODE_CONNECTED"}
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

	// Clear viewport, draw your native logo, and paint the matching layout cards
	output.ClearLine()
	output.Banner()
	output.Render(&payload)
}
