package intel

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// IsVPNActive is the centralized "Kill-Switch" check.
// It uses Triple-Check validation: Interfaces, Routing, and ISP Identity.
func IsVPNActive() bool {
	// 1. Hardware/Interface Check
	data, _ := os.ReadFile("/proc/net/dev")
	content := string(data)
	if strings.Contains(content, "tun") || strings.Contains(content, "wg") || strings.Contains(content, "proton") {
		return true
	}

	// 2. Routing Table Check (Bypasses container isolation in Parrot/Termux)
	out, _ := exec.Command("sh", "-c", "ip route | grep -E 'tun|wg|proton'").Output()
	if len(out) > 0 {
		return true
	}

	// 3. ISP Identity Check (External Truth - Detects Datacamp/Proton nodes)
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/?fields=isp")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		sBody := strings.ToLower(string(body))
		// Matches Datacamp (Proton infrastructure) and other common VPN providers
		if strings.Contains(sBody, "datacamp") || strings.Contains(sBody, "proton") || strings.Contains(sBody, "m247") {
			return true
		}
	}

	return false
}
