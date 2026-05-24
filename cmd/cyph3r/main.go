package main

import (
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
		return models.TypeGeoTarget
	}
	return models.TargetDomain
}

func main() {
	// Core flags
	targetFlag := flag.String("target", "", "Target: IP, Domain, or Email")
	phoneFlag := flag.String("phone", "", "Phone number lookup")
	scanFlag := flag.Bool("scan", false, "Execute port scan")
	monitorFlag := flag.Bool("monitor", false, "Continuous HUD monitor")
	protoFlag := flag.String("proto", "tcp", "Protocol for monitor")
	intervalFlag := flag.String("interval", "2s", "Monitor interval")
	fullFlag := flag.Bool("full", false, "Enable elite full dox report")
	verboseFlag := flag.Bool("v", false, "Verbose mode")
	shieldFlag := flag.Bool("shield", false, "Check current connection shield")

	// Stress test flags
	runTestFlag := flag.Bool("test-integrity", false, "Run stress validation suite")

	flag.Parse()

	// Banner
	fmt.Print(output.ClearLine)
	output.Banner()

	// Shield Check
	if *shieldFlag {
		output.RenderShieldReport()
		return
	}

	// Direct Phone Lookup
	if *phoneFlag != "" {
		metrics := intel.GetPhoneMetrics(*phoneFlag)
		output.RenderPhoneReport(*phoneFlag, metrics.LineStatus, metrics.Carrier, metrics.Locale)

		if *fullFlag || *verboseFlag {
			report := intel.ExecuteFullDox(*phoneFlag, models.TargetPhone)
			output.RenderReport(report)
		}
		return
	}

	// Target required
	if *targetFlag == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal: --target or --phone is required.")
		flag.Usage()
		os.Exit(1)
	}

	rawInput := strings.TrimSpace(*targetFlag)
	targetType := detectTargetType(rawInput)
	target := sanitizeToDomain(rawInput)

	// Monitor Mode
	if *monitorFlag {
		interval, _ := time.ParseDuration(*intervalFlag)
		probes.ExecuteContinuousMonitor(rawInput, strings.ToLower(*protoFlag), interval)
		return
	}

	// Full Elite Dox Mode
	if *fullFlag || *verboseFlag {
		report := intel.ExecuteFullDox(target, targetType)

		// Add port scan if requested
		if *scanFlag && (targetType == models.TargetDomain || targetType == models.TargetIP) {
			openPorts := probes.ExecutePortScan(target)
			fmt.Printf("\n%s[ PORT SCAN RESULTS ]%s\n", output.Cyan, output.Reset)
			for _, p := range openPorts {
				fmt.Printf(" • %s\n", p)
			}
		}

		output.RenderReport(report)
		return
	}

	// Legacy Mode (original behavior)
	intelCache, _ := cache.NewResponseCache()
	var payload models.IntelPayload

	if intelCache != nil {
		if cached, found := intelCache.Get(target); found {
			json.Unmarshal(cached, &payload)
		}
	}

	if payload.Target == "" {
		payload = models.IntelPayload{
			Target:   target,
			Type:     targetType,
			ScanTime: time.Now(),
		}

		switch targetType {
		case models.TargetEmail, models.TypeEmailTarget:
			payload.OwnerName = intel.ResolveEmail(target)
		case models.TargetPhone:
			payload.Phone = intel.ResolvePhone(target)
		default:
			resIP, geo, asn, owner, date, _, _, _, _ := intel.ResolveNetwork(target)
			payload.ISP = fmt.Sprintf("Network Stack (%s)", resIP)
			payload.Geo = geo
			payload.ASN = asn
			payload.OwnerName = owner
			payload.CreatedDate = date
		}

		if intelCache != nil {
			_ = intelCache.Set(target, payload)
		}
	}

	if *scanFlag {
		openPorts := probes.ExecutePortScan(target)
		payload.OpenPorts = openPorts
	}

	output.Render(&payload)
}
