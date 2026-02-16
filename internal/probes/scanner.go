package probes

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func ScanPorts(target string) []string {
	var results []string
	ports := []int{80, 443, 8080, 2083, 2087, 22, 21}
	for _, p := range ports {
		addr := fmt.Sprintf("%s:%d", target, p)
		time.Sleep(100 * time.Millisecond) // Jitter
		conn, err := net.DialTimeout("tcp", addr, 1200*time.Millisecond)
		if err != nil { continue }
		
		info := grabBanner(conn)
		if info == "" { info = getHTTPInfo(target, p) }
		conn.Close()

		if strings.Contains(info, "7.2") || strings.Contains(info, "Apache/2.2") {
			info += " [!]"
		}
		results = append(results, fmt.Sprintf("%d (%s)", p, info))
	}
	return results
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(800 * time.Millisecond))
	b, _ := bufio.NewReader(conn).ReadString('\n')
	return strings.TrimSpace(b)
}

func getHTTPInfo(target string, p int) string {
	url := fmt.Sprintf("http://%s:%d", target, p)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := (&http.Client{Timeout: 1 * time.Second}).Do(req)
	if err != nil { return "OPEN" }
	defer resp.Body.Close()
	return fmt.Sprintf("HTTP %d | %s", resp.StatusCode, resp.Header.Get("Server"))
}
