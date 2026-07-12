package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/cache"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
	"github.com/AnonPhoenix420/cyph3r/internal/stress"
)

var (
	phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$|^7\d{9}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func sanitizeToDomain(input string) string {
	cleaned := strings.TrimSpace(input)
	if strings.Contains(cleaned, "://") {
		parts := strings.SplitN(cleaned, "://", 2)
		cleaned = parts[1]
	}
	if idx := strings.IndexAny(cleaned, "/?#"); idx != -1 {
		cleaned = cleaned[:idx]
	}
	return strings.TrimSpace(cleaned)
}

func isIP(input string) bool {
	return net.ParseIP(input) != nil
}

func main() {
	// Flags
	targetFlag := flag.String("target", "", "Target input node")
	phoneFlag := flag.String("phone", "", "Standalone telephony lookup")
	portFlag := flag.Int("p", 80, "Target port")
	hulkFlag := flag.Bool("hulk", false, "Engage resilience stress testing")
	protoFlag := flag.String("proto", "tcp", "Protocol (tcp/udp/http)")
	methodFlag := flag.String("method", "GET", "HTTP method (GET/POST)")
	concurrencyFlag := flag.Int("c", 50, "Concurrency streams")
	durationFlag := flag.Int("d", 10, "Duration in seconds")
	monitorFlag := flag.Bool("monitor", false, "HUD monitor loop")
	intervalFlag := flag.String("interval", "2s", "Interval")
	jsonFlag := flag.Bool("json", false, "Output as JSON")
	runTestFlag := flag.Bool("test-integrity", false, "Run integrity suite")
	testModeFlag := flag.Int("mode", 1, "Mode: 1=LOAD, 2=STRESS")

	flag.Usage = func() { output.DisplayHelp() }
	flag.Parse()

	rawInput := strings.TrimSpace(*targetFlag)
	if rawInput == "" && *phoneFlag != "" { rawInput = strings.TrimSpace(*phoneFlag) }
	if rawInput == "" {
		fmt.Fprintln(os.Stderr, "[-] Fatal Error: Target identifier required.")
		os.Exit(1)
	}

	cleanHost := sanitizeToDomain(rawInput)
	targetAddr := net.JoinHostPort(cleanHost, fmt.Sprintf("%d", *portFlag))

	// HULK RESILIENCE ENGINE
	if *hulkFlag {
		fmt.Printf("[!] ENGAGING HULK MODE: %s on %s\n", targetAddr, *protoFlag)
		targetURL := "http://" + targetAddr
		switch strings.ToLower(*protoFlag) {
		case "udp":
			stress.ExecuteUDPFlood(targetAddr, *concurrencyFlag, *durationFlag)
		case "tcp":
			stress.ExecuteTCPFlood(targetAddr, *concurrencyFlag, *durationFlag)
		case "http":
			stress.ExecuteHTTPCapacityTest(targetURL, strings.ToUpper(*methodFlag), *concurrencyFlag, *durationFlag)
		}
		return
	}

	// MONITOR & INTEGRITY
	if *monitorFlag {
		output.Banner()
		interval, _ := time.ParseDuration(*intervalFlag)
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), interval)
		return
	}
	if *runTestFlag {
		output.Banner()
		intel.ExecuteValidationSuite("http://"+cleanHost, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

	// RECON ENGINE
	var target string
	var targetType models.TargetType
	if emailRegex.MatchString(rawInput) {
		target = rawInput
		targetType = models.TypeEmailTarget
	} else if phoneRegex.MatchString(rawInput) {
		target = rawInput
		targetType = models.TypePhoneTarget
	} else {
		target = cleanHost
		targetType = models.TypeNetworkTarget
	}

	intelCache, _ := cache.NewResponseCache()
	var payload models.IntelPayload
	var cacheHit = false
	if intelCache != nil {
		if cachedData, found := intelCache.Get(target); found {
			json.Unmarshal(cachedData, &payload)
			cacheHit = true
		}
	}

	if !cacheHit {
		payload = models.IntelPayload{Target: target, Type: targetType, ScanTime: time.Now()}
		
		if targetType == models.TypeNetworkTarget {
			fmt.Printf("[+] Analyzing: %s\n", target)
			
			// Call the engine with the exact 10 return values
			ip, geo, asn, owner, created, ports, banners, vulns, leaks, sql := intel.ResolveNetworkElite(target, 0, "")
			
			// Map results to the updated struct
			payload.TargetIP = ip
			payload.Geo = geo
			payload.ASN = asn
			payload.OwnerName = owner
			payload.CreatedDate = created
			payload.OpenPorts = ports
			payload.Banners = banners
			payload.Vulnerabilities = vulns
			payload.ExposedLeaks = leaks
			payload.SQLMetrics = sql
		}
	}
	
	if *jsonFlag {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", " ")
		encoder.Encode(payload)
	} else {
		output.Render(&payload)
	}
}
