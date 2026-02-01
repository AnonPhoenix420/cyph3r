package output

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ANSI Color Definitions
var (
	Cyan    = color.New(color.FgCyan).SprintFunc()
	White   = color.New(color.FgWhite, color.Bold).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
)

// PrintIntelDisplay handles the high-density reconnaissance HUD
func PrintIntelDisplay(target string, ips []string, ns []string, born string, reg string, isp string, loc string, coords string) {
	fmt.Println(Cyan("\n──[ NODE INTELLIGENCE ]──"))

	// Helper to handle long lists (IPs/Name Servers)
	formatList := func(list []string) string {
		if len(list) == 0 { return "None Detected" }
		res := strings.Join(list, ", ")
		if len(res) > 55 { return res[:52] + "..." }
		return res
	}

	// Layout grid with fixed 15-character padding for perfect alignment
	fmt.Printf(" %-15s %s\n", White("TARGET IP(s):"), Yellow(formatList(ips)))
	fmt.Printf(" %-15s %s\n", White("REGISTRAR:"), reg)
	fmt.Printf(" %-15s %s\n", White("BORN ON:"), Magenta(born))
	fmt.Printf(" %-15s %s\n", White("ISP/ORG:"), isp)
	fmt.Printf(" %-15s %s\n", White("LOCATION:"), loc)
	fmt.Printf(" %-15s %s\n", White("NAME SERVERS:"), Cyan(formatList(ns)))
	fmt.Printf(" %-15s %s\n", White("COORDINATES:"), coords)
	fmt.Println()
}

// Banner prints the main ASCII logo (Assumes your ASCII is stored here)
func Banner() {
	banner := `
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM`
	
	fmt.Println(Cyan(banner))
	fmt.Printf("  %s\n", White("v2.6 [STABLE] // Wireframe HUD Edition"))
	fmt.Println(Cyan("  ───────────────────────────────────────"))
}

// ScanAnimation provides the "Sensor Calibration" visual effect
func ScanAnimation() {
	frames := []string{"◒", "◐", "◓", "◑"}
	fmt.Print(White("[*] Calibrating HUD Sensors... "))
	for i := 0; i < 10; i++ {
		fmt.Printf("\r[*] Calibrating HUD Sensors... %s ", Cyan(frames[i%len(frames)]))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r[*] Calibrating HUD Sensors... %s\n", Green("[READY]"))
}

// Helper wrappers for Live Feed status
func Success(msg string) { fmt.Printf("%s %s\n", Green("[+]"), msg) }
func Info(msg string)    { fmt.Printf("%s %s\n", Blue("[*]"), msg) }
func Warn(msg string)    { fmt.Printf("%s %s\n", Yellow("[!]"), msg) }
func Down(msg string)    { fmt.Printf("%s %s\n", Red("[-]"), msg) }

func BlueText(text string) string { return Blue(text) }
