package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	targetPtr := flag.String("t", "", "Target domain")
	phonePtr := flag.String("p", "", "Phone number")
	verbosePtr := flag.Bool("v", false, "Verbose output")
	flag.Parse()

	output.Banner()
	fmt.Printf("\033[37m[*] INFO: Initializing Satellite Uplink...\033[0m\n")

	if *targetPtr != "" {
		runTargetScan(*targetPtr, *verbosePtr)
		os.Exit(0)
	}
    // ... Phone logic ...
}

func runTargetScan(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)
	data, err := intel.GetTargetIntel(target)
	done <- true

	if err != nil {
		fmt.Printf("\n\033[31m[!] UPLINK FAILURE: %v\033[0m\n", err)
		return
	}

	// --- ENHANCED HUD DISPLAY ---
	fmt.Printf("\n╔═══════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║ [!] TARGET_NODE: %-46s ║\n", data.TargetName)
	if data.IsWAF {
		fmt.Printf("║ [!] SHIELD:      \033[33m%-46s\033[0m ║\n", data.WAFType)
	}
	fmt.Printf("╚═══════════════════════════════════════════════════════════════╝\n")

	fmt.Println("\n[ NETWORK_VECTORS ]")
	for i, ip := range data.TargetIPs {
		ptr := "---"
		if i < len(data.ReverseDNS) { ptr = data.ReverseDNS[i] }
		fmt.Printf(" ↳ [v4] %-16s → \033[35m%s\033[0m \033[32m[ACTIVE]\033[0m\n", ip, ptr)
	}

	fmt.Println("\n[ INFRASTRUCTURE_STACK ]")
	for _, res := range data.ScanResults {
		if strings.HasPrefix(res, "USAGE:") {
			fmt.Printf("[*] TYPE:   \033[36m%s\033[0m\n", strings.TrimPrefix(res, "USAGE: "))
		} else {
			fmt.Printf("[+] %s\n", res)
		}
	}
    
    if verbose {
        fmt.Println("\n[ RAW_GEO_DATA ]")
        fmt.Println(data.RawGeo)
    }
}
