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
	geoRegex   = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
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
	// Essential Flags
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

	// Custom Help
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
		default:
			fmt.Println("[-] Invalid protocol. Use tcp, udp, or http.")
		}
		return
	}

	// MONITOR
	if *monitorFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		interval, _ := time.ParseDuration(*intervalFlag)
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), interval)
		return
	}

	// INTEGRITY
	if *runTestFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		intel.ExecuteValidationSuite("http://"+cleanHost, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

	// RECON ENGINE
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
		target = cleanHost
		targetType = models.TypeNetworkTarget
	}

	// EXECUTE RECON
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
			if !isIP(target) {
				fmt.Println("[+] Domain detected. Performing passive DNS analysis...")
			} else {
				fmt.Println("[+] IP detected. Skipping DNS analysis...")
			}
		}
		// (Continue with your existing intel logic here)
	}
	
	// Final Output
	if *jsonFlag {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", " ")
		encoder.Encode(payload)
	} else {
		output.Render(&payload)
	}
}
