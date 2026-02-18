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
	// 1. Tactical Argument Mapping
	targetPtr := flag.String("t", "", "Target domain (e.g., google.com)")
	phonePtr := flag.String("p", "", "Phone number (e.g., +13302354552)")
	verbosePtr := flag.Bool("v", false, "Enable verbose output (Raw JSON)")
	flag.Parse()

	// 2. Initialize Visual Uplink
	output.Banner()
	fmt.Printf("\033[37m[*] INFO: Initializing Satellite Uplink...\033[0m\n")

	// 3. Vector Execution Logic
	if *targetPtr != "" {
		runTargetScan(*targetPtr, *verbosePtr)
		os.Exit(0)
	}

	if *phonePtr != "" {
		runPhoneTrace(*phonePtr)
		os.Exit(0)
	}

	// 4. Default Failover
	fmt.Printf("\n\033[31m[!] ERROR: No search vector provided.\033[0m\n")
	fmt.Println("Usage:")
	fmt.Println(" ./cyph3r -t <domain> (Network Intelligence)")
	fmt.Println(" ./cyph3r -p <phone> (Mobile Intelligence)")
	os.Exit(1)
}

func runTargetScan(target string, verbose bool) {
	done := make(chan bool)
	go output.LoadingAnimation(done, target)
	
	// We use the 'models' package here to type-check the data
	var data models.IntelData
	var err error
	
	data, err = intel.GetTargetIntel(target)
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
		if i < len(data.ReverseDNS) {
			ptr = data.ReverseDNS[i]
		}
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

func runPhoneTrace(number string) {
	done := make(chan bool)
	go output.LoadingAnimation(done, number)
	
	data, err := intel.GetPhoneIntel(number)
	done <- true

	if err != nil {
		fmt.Printf("\n\033[31m[!] TRACE FAILURE: %v\033[0m\n", err)
		return
	}
	
	// Assuming your output package has this, otherwise print manually
	fmt.Printf("\n[+] PHONE INTEL RECOVERED: %s (%s)\n", data.Number, data.Carrier)
}
