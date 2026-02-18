package intel

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// Global Redaction Map
var redactionList = []string{
	"/data/data/com.termux/files/home", // Termux Home
	"parrot",                            // OS Name
	os.Getenv("USER"),                   // Current User
}

// FailSafeCheck confirms VPN and scrubs the environment
func FailSafeCheck() {
	if !isVPNActive() {
		fmt.Println("\n\033[31m[!] CRITICAL: PROTON VPN NOT DETECTED.\033[0m")
		fmt.Println("[*] OPSEC LOCK: Terminating process to prevent IP leak.")
		os.Exit(1)
	}
}

func isVPNActive() bool {
	ifaces, _ := net.Interfaces()
	// Proton typically uses tun0 (OpenVPN) or wg0/proton0 (WireGuard)
	for _, i := range ifaces {
		if (i.Flags&net.FlagUp != 0) && (strings.HasPrefix(i.Name, "tun") || strings.HasPrefix(i.Name, "wg") || strings.HasPrefix(i.Name, "proton")) {
			return true
		}
	}
	return false
}

func scrub(input string) string {
	output := input
	hostname, _ := os.Hostname()
	output = strings.ReplaceAll(output, hostname, "TARGET_NODE")
	
	for _, item := range redactionList {
		if item != "" {
			output = strings.ReplaceAll(output, item, "[REDACTED]")
		}
	}
	return output
}

// Full Drop-in Replacement for fetchGeo
func fetchGeo(ip string) (models.GeoResponse, string) {
	// Step 0: Kill process if no VPN
	FailSafeCheck()

	client := GetClient()
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return models.GeoResponse{Org: "SECURE_UPLINK"}, "{}"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r models.GeoResponse
	json.Unmarshal(body, &r)

	// Pretty Print & Scrub
	var anyData interface{}
	json.Unmarshal(body, &anyData)
	prettyJSON, _ := json.MarshalIndent(anyData, "", "  ")

	// Apply 99.9% Clean Sanity
	cleanRaw := scrub(string(prettyJSON))

	return r, cleanRaw
}
