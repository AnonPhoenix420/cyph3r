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
	phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}\( |^7\d{9} \)`)
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

func detectTargetType(rawInput string) models.TargetType {
	clean := strings.TrimSpace(strings.ToLower(rawInput))
	if emailRegex.MatchString(clean) {
		return models.TargetEmail
	}
	if phoneRegex.MatchString(strings.ReplaceAll(clean, " ", "")) {
		return models.TargetPhone
	}
	if geoRegex.MatchString(clean) {
		return models.TargetGeo // Add to models if missing
	}
	return models.TargetDomain // Default to network/domain
}

func main() {
	// Traditional CYPH3R v2.6 flags
	targetFlag := flag.String("target", "", "Target routing domain, email, or infrastructure IP vector")
	phoneFlag := flag.String("phone", "", "Execute international telephone vector metadata lookup")
	scanActiveFlag := flag.Bool("scan", false, "Execute tactical concurrent port scan and banner analysis")
	monitorFlag := flag.Bool("monitor", false, "Engage continuous HUD telemetry monitoring mode loop")
	protoFlag := flag.String("proto", "tcp", "Protocol mode selector for tracing [tcp, udp, http, https, ack]")
	intervalFlag := flag.String("interval", "2s", "Telemetry delay frequency window interval")
	
	// Stress Validation
	runTestFlag := flag.Bool("test-integrity", false, "Execute integrated validation stress suite")
	testModeFlag := flag.Int("mode", 1, "Select test vector")
	concurrencyFlag := flag.Int("c", 50, "Number of concurrent streams")
	durationFlag := flag.Int("d", 10, "Duration in seconds")
	
	verboseFlag := flag.Bool("v", false, "Enable full operational tracing logs")
	jsonFlag := flag.Bool("json", false, "Output data structure as raw JSON matrix")
	fullFlag := flag.Bool("full", false, "Enable elite comprehensive dox report")

	flag.Parse()

	fmt.Println(`
   ______      ____  __  __ _____ ____
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/
\____/\__, /_/   /_/ /_/ /____/_/ |_|
     /____/         NETWORK_INTEL_SYSTEM
`)

	// 1. Direct Phone Handling (Legacy)
	if *phoneFlag != "" {
		fmt.Print(output.ClearLine)
		output.Banner()
		metrics := intel.GetPhoneMetrics(*phoneFlag)
		output.RenderPhoneReport(*phoneFlag, metrics.LineStatus, metrics.Carrier, metrics.Locale)

		if *fullFlag || *verboseFlag {
			report := intel.ExecuteFullDox(*phoneFlag, models.TargetPhone)
			output.RenderReport(report)
		}
		return
	}

	// 2. Target Required Check
	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal: Operational parameter target mapping (--target or --phone) strictly required.")
		os.Exit(1)
	}

	rawInput := strings.TrimSpace(*targetFlag)
	targetType := detectTargetType(rawInput)
	target := sanitizeToDomain(rawInput)

	// 3. Monitoring Mode
	if *monitorFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		interval, _ := time.ParseDuration(*intervalFlag)
		probes.ExecuteContinuousMonitor(rawInput, strings.ToLower(*protoFlag), interval)
		return
	}

	// 4. Stress Test Mode
	if *runTestFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		targetURL := rawInput
		if !strings.HasPrefix(targetURL, "http") {
			targetURL = "http://" + targetURL
		}
		intel.ExecuteValidationSuite(targetURL, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

	// 5. Elite Full Dox Mode (New)
	useFullDox := *fullFlag || *verboseFlag
	if useFullDox {
		fmt.Print(output.ClearLine)
		output.Banner()
		report := intel.ExecuteFullDox(target, targetType)
		output.RenderReport(report)
		return
	}

	// 6. Legacy Processing Path (Preserved)
	intelCache, _ := cache.NewResponseCache()
	var payload models.IntelPayload
	var cacheHit = false

	// Cache logic (your original)
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
		case models.TargetEmail:
			payload.OwnerName = intel.ResolveEmail(target) // Keep your existing function
		case models.TargetPhone:
			payload.Phone = intel.ResolvePhone(target)
		case models.TypeNetworkTarget, models.TargetDomain, models.TargetIP:
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
