#!/bin/bash

# 1. Clean Slate
echo "[*] Nuking old config..."
rm -rf go.mod go.sum internal/ cmd/
export GO111MODULE=on

# 2. Setup Directories
echo "[*] Creating folder structure..."
mkdir -p cmd/cyph3r internal/models internal/intel internal/probes internal/output

# 3. Initialize Module (This fixes the 'not in std' error)
echo "module cyph3r" > go.mod
echo "" >> go.mod
echo "go 1.23" >> go.mod

# 4. Drop-in File: models.go
cat <<EOF > internal/models/models.go
package models
type IntelData struct {
	IPs, Nameservers []string
	WhoisRaw, Registrar, Country, City, ISP, Org, ASN string
	Lat, Lon float64
}
EOF

# 5. Drop-in File: intel.go (Fixed Imports)
cat <<EOF > internal/intel/intel.go
package intel
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"cyph3r/internal/models"
	"github.com/likexian/whois"
)
func GetFullIntel(target string) (models.IntelData, error) {
	var d models.IntelData
	d.IPs, d.Nameservers = LookupNodes(target)
	raw, err := whois.Whois(target)
	if err == nil { d.WhoisRaw = raw; d.Registrar = extract(raw, "Registrar:") }
	q := target
	if len(d.IPs) > 0 { q = d.IPs[0] }
	c := http.Client{Timeout: 5 * time.Second}
	resp, err := c.Get("http://ip-api.com/json/" + q)
	if err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&d)
	}
	return d, nil
}
func extract(raw, field string) string {
	for _, line := range strings.Split(raw, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(field)) {
			parts := strings.Split(line, ":")
			if len(parts) > 1 { return strings.TrimSpace(parts[1]) }
		}
	}
	return "UNKNOWN"
}
EOF

# 6. Drop-in File: dns.go
cat <<EOF > internal/intel/dns.go
package intel
import "net"
func LookupNodes(target string) ([]string, []string) {
	var ips, nss []string
	addr, _ := net.LookupHost(target)
	ips = addr
	ns, _ := net.LookupNS(target)
	for _, r := range ns { nss = append(nss, r.Host) }
	return ips, nss
}
EOF

# 7. Drop-in File: main.go
cat <<EOF > cmd/cyph3r/main.go
package main
import (
	"flag"
	"fmt"
	"cyph3r/internal/intel"
)
func main() {
	target := flag.String("target", "", "Domain to analyze")
	flag.Parse()
	if *target == "" { fmt.Println("Usage: cyph3r -target <domain>"); return }
	fmt.Printf("[*] Scanning: %s\n", *target)
	data, _ := intel.GetFullIntel(*target)
	fmt.Printf("[+] Registrar: %s\n", data.Registrar)
}
EOF

# 8. Final Sync and Build
echo "[*] Syncing dependencies (go mod tidy)..."
go mod tidy
echo "[*] Compiling binary..."
go build -o cyph3r ./cmd/cyph3r/main.go

if [ -f "./cyph3r" ]; then
    echo "[!] SUCCESS: ./cyph3r is ready."
else
    echo "[!] ERROR: Build failed. Check the output above."
fi
