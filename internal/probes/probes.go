package probes

import (
	"fmt"
	"net"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

// RunFullScan executes a multi-port check on the target
func RunFullScan(target string) {
	output.Info("Initializing Tactical Scan: " + target)
	
	// Common ports to check
	ports := []int{21, 22, 23, 25, 53, 80, 443, 8080, 3306}

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		
		if err != nil {
			// Port is closed or filtered
			continue
		}
		conn.Close()
		
		fmt.Printf("%s[+] PORT %d: %sOPEN%s\n", output.White, port, output.NeonGreen, output.Reset)
	}
	output.Info("Tactical scan complete.")
}
