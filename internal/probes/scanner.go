package probes

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// Vulnerability signatures (Simplified example)
var highRiskSignatures = []string{"OpenSSH_7.2", "Apache/2.2", "php/5.", "IIS/6.0", "Expired"}

func ScanPorts(target string) []string {
	var results []string
	ninjaPorts := []int{80, 443, 8080, 8443, 2082, 2083, 2086, 2087, 21, 22, 23, 25, 3306, 6443}

	for _, port := range ninjaPorts {
		address := fmt.Sprintf("%s:%d", target, port)
		time.Sleep(150 * time.Millisecond) 

		conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond)
		if err != nil { continue }

		info := ""
		if isSSL(port) {
			info = getSSLInfo(target, port)
		} else {
			info = grabBanner(conn)
			if info == "" { info = getHTTPInfo(target, port) }
		}
		conn.Close()

		// Ninja Upgrade: Vulnerability flagging
		riskTag := ""
		for _, sig := range highRiskSignatures {
			if strings.Contains(strings.ToLower(info), strings.ToLower(sig)) {
				riskTag = " [!] VULN_POTENTIAL"
			}
		}

		if info != "" {
			results = append(results, fmt.Sprintf("%d (%s)%s", port, info, riskTag))
		}
	}
	return results
}

// ... (keep your getHTTPInfo, getSSLInfo, and grabBanner from the last build)
