package intel

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ShieldInfo struct {
	IsActive bool
	IP       string
	Location string
	ISP      string
}

func CheckShield() ShieldInfo {
	info := ShieldInfo{IsActive: false}
	data, _ := os.ReadFile("/proc/net/dev")
	route, _ := exec.Command("sh", "-c", "ip route").Output()
	combined := string(data) + string(route)
	
	if strings.Contains(combined, "tun") || strings.Contains(combined, "proton") || strings.Contains(combined, "wg") {
		info.IsActive = true
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/?fields=status,country,city,isp,query")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var r struct {
			Status  string `json:"status"`
			Country string `json:"country"`
			City    string `json:"city"`
			Isp     string `json:"isp"`
			Query   string `json:"query"`
		}
		json.Unmarshal(body, &r)
		if r.Status == "success" {
			info.IP, info.Location, info.ISP = r.Query, r.City+", "+r.Country, r.Isp
			sIsp := strings.ToLower(r.Isp)
			if strings.Contains(sIsp, "datacamp") || strings.Contains(sIsp, "proton") || strings.Contains(sIsp, "m247") {
				info.IsActive = true
			}
		}
	}
	return info
}
