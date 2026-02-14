package probes

import (
	"fmt"
	"net"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func RunFullScan(target string) {
	output.Info("Initializing Tactical Scan: " + target)
	ports := []int{21, 22, 23, 25, 53, 80, 443, 3306, 8080}

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		
		if err != nil {
			continue // Port Closed
		}
		conn.Close()
		
		// Signaling the specific protocol and ACK
		protocol := "TCP"
		if port == 53 { protocol = "DNS/UDP" }
		if port == 443 { protocol = "HTTPS" }
		
		fmt.Printf("%s[+] PORT %d/%s: %sOPEN %s[ACK/SYN]%s\n", 
			output.White, port, protocol, output.NeonGreen, output.NeonBlue, output.Reset)
	}
	output.Info("Tactical scan complete.")
}
