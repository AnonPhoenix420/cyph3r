package output
import (
	"fmt"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/probes" // Import results struct
)

func PrintIntelHUD(t string, ips, ns []string, b, r, isp, loc, c string) {
	fmt.Println(CyanText("──[ NODE INTELLIGENCE ]──"))
	fmt.Printf(" %-15s %s\n", WhiteText("TARGET:"), YellowText(t))
	fmt.Printf(" %-15s %s\n", WhiteText("IPS:"), strings.Join(ips, ", "))
	fmt.Printf(" %-15s %s\n", WhiteText("ISP:"), isp)
	fmt.Printf(" %-15s %s\n", WhiteText("LOC:"), loc)
	fmt.Println()
}

func PrintPortScan(res []probes.ScanResult) {
	fmt.Println(CyanText("──[ OPEN SERVICES ]──"))
	for _, r := range res {
		fmt.Printf(" %-15s %-10d %s\n", WhiteText("PORT:"), r.Port, GreenText("[OPEN]"))
	}
}

func Success(m string) { fmt.Printf("%s %s\n", GreenText("[+]"), m) }
func Warn(m string)    { fmt.Printf("%s %s\n", YellowText("[!]"), m) }
func Info(m string)    { fmt.Printf("%s %s\n", BlueText("[*]"), m) }
func Down(m string)    { fmt.Printf("%s %s\n", RedText("[-]"), m) }
