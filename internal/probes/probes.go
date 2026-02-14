package probes

import (
	"fmt"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func RunFullScan(target string) {
	ports := []int{21, 22, 25, 53, 80, 443, 3306, 8080}
	output.PrintScanHeader()
	for _, port := range ports {
		alive, status := DialTarget(target, port)
		if alive {
			fmt.Printf("%s[+] PORT %d: %s%s\n", output.NeonGreen, port, status, output.Reset)
		}
	}
}
