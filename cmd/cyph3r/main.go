package main

import (
	"flag"
	"fmt"
	"time"

	"cyph3r/internal/geo"
	"cyph3r/internal/netcheck"
	"cyph3r/internal/output"
	"cyph3r/internal/phone"
	"cyph3r/internal/version"
)

func main() {
	target := flag.String("target", "localhost", "Target host or IP")
	port := flag.String("port", "80", "Port number")
	proto := flag.String("proto", "tcp", "Protocol: tcp | udp | http | https")
	phoneNum := flag.String("phone", "", "Phone number metadata lookup")
	jsonOut := flag.Bool("json", false, "JSON output")
	monitor := flag.Bool("monitor", false, "Continuously monitor target")
	interval := flag.Int("interval", 5, "Check interval in seconds")
	showVer := flag.Bool("version", false, "Show version info")
	flag.Parse()

	if *showVer {
		fmt.Printf("%s v%s by %s\n", version.Name, version.Version, version.Author)
		return
	}

	output.Banner()
	geo.Lookup(*target)

	var wasUp *bool
	var downSince time.Time

	for {
		up := false
		result := map[string]any{
			"target": *target,
			"proto":  *proto,
			"time":   time.Now().Format(time.RFC3339),
		}

		switch *proto {
		case "tcp":
			ok, latency := netcheck.TCP(*target, *port)
			up = ok
			result["up"] = ok
			result["latency_ms"] = latency

		case "udp":
			ok := netcheck.UDP(*target, *port)
			up = ok
			result["up"] = ok

		case "http", "https":
			code, latency := netcheck.HTTP(*proto+"://"+*target+":"+*port, true)
			up = code > 0
			result["status"] = code
			result["latency_ms"] = latency
			result["up"] = up

		default:
			fmt.Println("Unknown protocol:", *proto)
			return
		}

		// --- State tracking ---
		if wasUp == nil {
			wasUp = new(bool)
			*wasUp = up

			if up {
				output.Up("Target is UP")
			} else {
				downSince = time.Now()
				output.Down("Target is DOWN")
			}
		} else {
			// UP → DOWN
			if *wasUp && !up {
				downSince = time.Now()
				output.Down("Target went DOWN")
			}

			// DOWN → UP
			if !*wasUp && up {
				downtime := time.Since(downSince).Round(time.Second)
				output.Up(fmt.Sprintf("Target is UP again (downtime: %s)", downtime))
				result["downtime"] = downtime.String()
			}
		}

		*wasUp = up

		if *phoneNum != "" {
			phone.Lookup(*phoneNum)
		}

		if *jsonOut {
			output.PrintJSON(result)
		}

		if !*monitor {
			break
		}

		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
