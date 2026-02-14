package probes

import (
	"fmt"
	"net"
	"time"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

func RunFullScan(target string) {
	output.Info("Initializing Tactical Scan: " + target)
	
	// Comprehensive port list for server testing
	ports := []struct {
		num   int
		proto string
	}{
		{21, "FTP"}, {22, "SSH"}, {23, "TELNET"}, {25, "SMTP"},
		{53, "DNS"}, {80, "HTTP"}, {110, "POP3"}, {143, "IMAP"},
		{443, "HTTPS"}, {3306, "MYSQL"}, {5432, "POSTGRES"}, {8080, "HTTP-ALT"},
	}

	for _, p := range ports {
		address := fmt.Sprintf("%s:%d", target, p.num)
		conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
		
		if err != nil {
			continue 
		}
		conn.Close()
		
		fmt.Printf("%s[+] PORT %d/%s: %sOPEN %s[ACK/SYN]%s\n", 
			output.White, p.num, p.proto, output.NeonGreen, output.NeonBlue, output.Reset)
	}
	output.Info("Tactical scan complete.")
}
