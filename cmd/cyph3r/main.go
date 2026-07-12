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
	// FIXED: Preserving ports (:)
	if idx := strings.IndexAny(cleaned, "/?#"); idx != -1 {
		cleaned = cleaned[:idx]
	}
	return strings.TrimSpace(cleaned)
}

func main() {
	targetFlag := flag.String("target", "", "Target input node routing configuration vector")
	phoneFlag := flag.String("phone", "", "Execute standalone telephony metadata lookup")
	portFlag := flag.Int("p", 80, "Target port for stress/recon")
	tcpFlag := flag.Bool("tcp", false, "Force TCP protocol usage")
	hulkFlag := flag.Bool("hulk", false, "Engage high-intensity stress mode")
	
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

	cleanHost := sanitizeToDomain(rawInput)
	targetAddr := net.JoinHostPort(cleanHost, fmt.Sprintf("%d", *portFlag))

	// HULK STRESS MODE
	if *hulkFlag {
		fmt.Printf("[!] ENGAGING HULK MODE: %s on %s\n", targetAddr, *protoFlag)
		stress.ExecuteHighIntensityStress(targetAddr, *concurrencyFlag, *durationFlag)
		return
	}

	if *monitorFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		interval, _ := time.ParseDuration(*intervalFlag)
		if interval == 0 {
			interval = 2 * time.Second
		}
		probes.ExecuteContinuousMonitor(targetAddr, strings.ToLower(*protoFlag), interval)
		return
	}

	if *runTestFlag {
		fmt.Print(output.ClearLine)
		output.Banner()
		intel.ExecuteValidationSuite(targetAddr, *testModeFlag, *concurrencyFlag, *durationFlag)
		return
	}

	// [RECON ENGINE REMAINS UNTOUCHED]
	// ... (Rest of your existing recon logic follows here) ...
	// (Ensure you use 'cleanHost' for your Intel lookups)
	var target = cleanHost 
	// ... 
}
