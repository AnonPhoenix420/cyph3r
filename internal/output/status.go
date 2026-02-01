package output

import (
	"fmt"
	"strings"
)

// PrintIntelHUD handles the high-density reconnaissance display
func PrintIntelHUD(target string, ips []string, ns []string, born string, reg string, isp string, loc string, coords string) {
	fmt.Println(CyanText("\n──[ NODE INTELLIGENCE ]──"))

	format := func(label, value string) {
		fmt.Printf(" %-15s %s\n", WhiteText(label+":"), value)
	}

	ipList := strings.Join(ips, ", ")
	if len(ipList) > 50 { ipList = ipList[:47] + "..." }

	format("TARGET IP(s)", YellowText(ipList))
	format("REGISTRAR", reg)
	format("BORN ON", MagentaText(born))
	format("ISP/ORG", isp)
	format("LOCATION", loc)
	format("NAME SERVERS", CyanText(strings.Join(ns, ", ")))
	format("COORDINATES", coords)
	fmt.Println()
}

// HUD Status Helpers
func Success(msg string) { fmt.Printf("%s %s\n", GreenText("[+]"), msg) }
func Info(msg string)    { fmt.Printf("%s %s\n", BlueText("[*]"), msg) }
func Warn(msg string)    { fmt.Printf("%s %s\n", YellowText("[!]"), msg) }
func Down(msg string)    { fmt.Printf("%s %s\n", RedText("[-]"), msg) }
