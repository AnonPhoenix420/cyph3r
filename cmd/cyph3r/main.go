package main

import (
	"flag"
	"fmt"
	"os"
	"cyph3r/internal/intel"
)

func main() {
	// 1. Define CLI Flags
	target := flag.String("target", "", "Domain or IP address to analyze")
	scan := flag.Bool("scan", false, "Execute tactical Nmap port scan")
	flag.Parse()

	// 2. Validation
	if *target == "" {
		fmt.Println("\033[31m[!] Error: Target is required.\033[0m")
		fmt.Println("Usage: ./cyph3r -target <domain> [-scan]")
		os.Exit(1)
	}

	// 3. Execution Header
	fmt.Printf("\n\033[35m[⚡] CYPH3R PULSE: %s\033[0m\n", *target)
	fmt.Println("\033[90m------------------------------------------\033[0m")

	// 4. Intelligence Gathering
	data, err := intel.GetFullIntel(*target)
	if err != nil {
		fmt.Printf("\033[31m[-] Intel phase failed: %v\033[0m\n", err)
	} else {
		fmt.Printf("\033[97m[+] IP ADDRESS: \033[32m%s\033[0m\n", data.IP)
		fmt.Printf("\033[97m[+] REGISTRAR:  \033[32m%s\033[0m\n", data.Registrar)
	}

	// 5. Tactical Scanning (The -scan flag)
	if *scan {
		fmt.Println("\n\033[33m[#] INITIATING TACTICAL NMAP PROBE...\033[0m")
		results := intel.RunSystemScan(*target)
		fmt.Println(results)
	}

	fmt.Println("\033[90m------------------------------------------\033[0m")
}
