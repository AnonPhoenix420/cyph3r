package probes
import (
	"net"
	"net/http"
	"time"
)
func ExecuteProbe(proto, target string, port int) (bool, time.Duration) {
	start := time.Now()
	addr := fmt.Sprintf("%s:%d", target, port)
	if proto == "http" || proto == "https" {
		c := &http.Client{Timeout: 2 * time.Second}
		_, err := c.Get(proto + "://" + target)
		return err == nil, time.Since(start)
	}
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err == nil { conn.Close() }
	return err == nil, time.Since(start)
}
