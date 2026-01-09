package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"cyph3r/output"
	"github.com/nyaruka/phonenumbers"
)

/*
====================
ENTRYPOINT
====================
*/
func main() {
	// Banner
	output.Banner()

	// Start application
	run()
}

/*
====================
APPLICATION CORE
====================
*/
func run() {
	target := flag.String("target", "localhost", "Target host or IP")
	port := flag.Int("port", 80, "Port number")
	proto := flag.String("proto", "tcp", "Protocol: tcp|udp|http|https|dns")
	geoip := flag.Bool("geoip", false, "GeoIP lookup")
	phone := flag.String("phone", "", "Phone number info")
	jsonOut := flag.Bool("json", false, "JSON output")
	monitor := flag.Bool("monitor", false, "Continuous monitoring")
	interval := flag.Int("interval", 5, "Monitor interval (seconds)")
	portscan := flag.Bool("portscan", false, "Perform port scan")
	scanstart := flag.Int("scanstart", 1, "Port scan start")
	scanend := flag.Int("scanend", 1024, "Port scan end")
	flag.Parse()

	// Resolve localhost → all interface IPs
	targets, err := resolveTargets(*target)
	if err != nil {
		fmt.Println(output.RedText("Target resolution failed:"), err)
		os.Exit(1)
	}

	if len(targets) > 1 {
		fmt.Println(output.BlueText("Localhost detected — scanning all interfaces:"))
		for _, t := range targets {
			fmt.Println(" •", t)
		}
		fmt.Println()
	}

	// GeoIP
	if *geoip {
		for _, t := range targets {
			res, err := geoIPLookup(t)
			if err != nil {
				fmt.Println(output.RedText("GeoIP failed for"), t, ":", err)
				continue
			}

			if *jsonOut {
				printJSON(res)
			} else {
				fmt.Println(output.YellowBoldText("GeoIP:"), t)
				fmt.Println(output.GeoLabel("Country:"), res.Country)
				fmt.Println(output.GeoLabel("Region:"), res.Region)
				fmt.Println(output.GeoLabel("City:"), res.City)
				fmt.Println(output.GeoLabel("ASN:"), res.ASN)
				fmt.Println(output.GeoLabel("Org:"), res.Org)
				fmt.Println()
			}
		}
		return
	}

	// Phone lookup
	if *phone != "" {
		lookupPhone(*phone)
		return
	}

	// Port scan
	if *portscan {
		for _, t := range targets {
			ports := scanPorts(t, *scanstart, *scanend)
			if *jsonOut {
				printJSON(map[string]interface{}{
					"target":     t,
					"open_ports": ports,
				})
			} else {
				fmt.Println(output.PinkBoldText("Port scan:"), t)
				for _, p := range ports {
					fmt.Println(output.PortLabel(fmt.Sprintf("Port %d/tcp OPEN", p)))
				}
				fmt.Println()
			}
		}
		return
	}

	// Diagnostics loop
	for {
		for _, t := range targets {
			up := false
			details := ""

			switch strings.ToLower(*proto) {
			case "tcp":
				addr := fmt.Sprintf("%s:%d", t, *port)
				conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
				up = err == nil
				if up {
					conn.Close()
				}
				details = fmt.Sprintf("err=%v", err)

			case "udp":
				up = checkUDP(t, *port)
				details = "udp probe"

			case "http":
				code, err := checkHTTP(t, false)
				up = err == nil && code >= 200 && code < 400
				details = fmt.Sprintf("HTTP %d err=%v", code, err)

			case "https":
				code, err := checkHTTP(t, true)
				up = err == nil && code >= 200 && code < 400
				details = fmt.Sprintf("HTTPS %d err=%v", code, err)

			case "dns":
				up = checkDNS(t)
				details = "DNS lookup"

			default:
				fmt.Println("Unknown protocol")
				return
			}

			label := fmt.Sprintf("%s %s:%d %s", *proto, t, *port, details)

			if *jsonOut {
				printJSON(map[string]interface{}{
					"target":  t,
					"port":    *port,
					"proto":   *proto,
					"up":      up,
					"details": details,
					"time":    time.Now().Format(time.RFC3339),
				})
			} else {
				if up {
					output.Up(label)
				} else {
					output.Down(label)
				}
			}
		}

		if !*monitor {
			break
		}
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

/*
====================
HELPERS
====================
*/

// Expand localhost to all local interface IPs
func resolveTargets(target string) ([]string, error) {
	if target != "localhost" && target != "127.0.0.1" {
		return []string{target}, nil
	}

	var targets []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			if ip = ip.To4(); ip != nil {
				targets = append(targets, ip.String())
			}
		}
	}

	return append(targets, "127.0.0.1"), nil
}

// --- GEOIP ---
type GeoResult struct {
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	ASN      string  `json:"as"`
	Org      string  `json:"org"`
	Hostname string  `json:"reverse"`
}

func geoIPLookup(target string) (*GeoResult, error) {
	resp, err := http.Get("http://ip-api.com/json/" + target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Country string `json:"country"`
		Region  string `json:"regionName"`
		City    string `json:"city"`
		Lat     float64
		Lon     float64
		ASN     string `json:"as"`
		Org     string `json:"org"`
		Host    string `json:"reverse"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if data.Status != "success" {
		return nil, fmt.Errorf(data.Message)
	}

	return &GeoResult{
		City:     data.City,
		Region:   data.Region,
		Country:  data.Country,
		Lat:      data.Lat,
		Lon:      data.Lon,
		ASN:      data.ASN,
		Org:      data.Org,
		Hostname: data.Host,
	}, nil
}

// --- PHONE ---
func lookupPhone(number string) {
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	printJSON(map[string]interface{}{
		"raw":    number,
		"valid":  phonenumbers.IsValidNumber(num),
		"region": phonenumbers.GetRegionCodeForNumber(num),
		"type":   phonenumbers.GetNumberType(num),
	})
}

// --- NETWORK HELPERS ---
func scanPorts(host string, start, end int) []int {
	open := []int{}
	for port := start; port <= end; port++ {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 300*time.Millisecond)
		if err == nil {
			open = append(open, port)
			conn.Close()
		}
	}
	return open
}

func checkUDP(target string, port int) bool {
	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%d", target, port), time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	_, err = conn.Write([]byte("ping"))
	return err == nil
}

func checkDNS(target string) bool {
	_, err := net.LookupHost(target)
	return err == nil
}

func checkHTTP(target string, https bool) (int, error) {
	scheme := "http"
	if https {
		scheme = "https"
	}
	resp, err := http.Get(scheme + "://" + target)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func printJSON(v interface{}) {
	js, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(js))
}
