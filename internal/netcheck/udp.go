package netcheck


import (
"net"
"time"
)


func UDP(host, port string) bool {
addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
if err != nil {
return false
}


conn, err := net.DialUDP("udp", nil, addr)
if err != nil {
return false
}
defer conn.Close()


conn.SetDeadline(time.Now().Add(2 * time.Second))
_, err = conn.Write([]byte("ping"))
return err == nil
}
