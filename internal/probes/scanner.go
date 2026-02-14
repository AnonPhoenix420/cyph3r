package probes

import (
	"fmt"
	"net"
	"time"
)

func DialTarget(target string, port int) (bool, string) {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false, "CLOSED"
	}
	defer conn.Close()
	return true, "ALIVE"
}
