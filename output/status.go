package output

import (
	"encoding/json"
	"fmt"
	"time"
)

// ================= JSON OUTPUT HELPERS =================

// PrintJSON pretty prints any Go structure in JSON format
func PrintJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(Red + "JSON Marshal Error:" + Reset, err)
		return
	}
	fmt.Println(string(b))
}

// PrintStatusJSON formats a network check result for JSON
func PrintStatusJSON(proto, target string, port int, up bool, details string) {
	out := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"proto":   proto,
		"target":  target,
		"port":    port,
		"up":      up,
		"details": details,
	}
	PrintJSON(out)
}

// ================= PROMETHEUS LABEL HELPERS =================

// BuildLabels generates labels for Prometheus metrics
func BuildLabels(proto, scenario, target string) map[string]string {
	return map[string]string{
		"proto":    proto,
		"scenario": scenario,
		"target":   target,
	}
}

// ================= DASHBOARD FORMAT =================

// PrintDashboard prints live metrics to terminal
func PrintDashboard(sent, success, failure uint64, avg, p50, p95, p99 time.Duration, failRate float64, latencyThreshold time.Duration, failThreshold float64) {
	fmt.Printf("%s[Dashboard]%s Sent=%d Success=%d Fail=%d Avg=%s P50=%s P95=%s P99=%s\n",
		BoldMagenta, Reset, sent, success, failure, avg, p50, p95, p99)

	if failThreshold > 0 && failRate > failThreshold {
		Down("FAILURE THRESHOLD BREACHED")
	}
	if latencyThreshold > 0 && p95 > latencyThreshold {
		Down("LATENCY THRESHOLD BREACHED")
	}
}

// ================= SIMPLE STATUS HELPERS =================

// PrintCheck prints a single check result
func PrintCheck(proto, target string, port int, up bool, details string, jsonOut bool) {
	if jsonOut {
		PrintStatusJSON(proto, target, port, up, details)
		return
	}

	label := fmt.Sprintf("%s %s:%d %s", proto, target, port, details)
	if up {
		Up(label)
	} else {
		Down(label)
	}
}
