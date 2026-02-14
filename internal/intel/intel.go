package intel

import (
	"os/exec"
	"strings"
	"cyph3r/internal/models"
)

func GetFullIntel(target string) (models.IntelData, error) {
	var d models.IntelData
	d.Target = target

	// 1. System WHOIS
	out, _ := exec.Command("whois", target).Output()
	raw := string(out)
	for _, line := range strings.Split(raw, "\n") {
		if strings.Contains(line, "Registrar:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				d.Registrar = strings.TrimSpace(parts[1])
				break
			}
		}
	}

	// 2. System DIG (Get IP)
	ipOut, _ := exec.Command("dig", "+short", target).Output()
	d.IP = strings.TrimSpace(string(ipOut))

	// 3. System CURL (Check Web Status)
	curlOut, _ := exec.Command("curl", "-Is", "--max-time", "2", target).Output()
	if len(curlOut) > 0 {
		d.HTTPStatus = strings.Split(string(curlOut), "\n")[0]
	}

	return d, nil
}

func RunSystemScan(target string) string {
	// Tactical Fast Scan: Top 100 ports, service detection, no ping
	out, err := exec.Command("nmap", "-F", "--open", target).Output()
	if err != nil {
		return "Error: nmap not found or failed. Install with: apt install nmap"
	}
	return string(out)
}
