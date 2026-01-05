package netcheck


import (
"net"
"time"
)


func TCP(host, port string) (bool, int64) {
start := time.Now()
conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 3*time.Second)
if err != nil {
return false, 0
}
conn.Close()
return true, time.Since(start).Milliseconds()
}
