package geo


import (
"fmt"
"net"
)


func Lookup(host string) {
ips, err := net.LookupIP(host)
if err != nil {
fmt.Println("GeoIP lookup failed")
return
}


fmt.Println("🌍 Resolved IPs:")
for _, ip := range ips {
fmt.Println(" -", ip.String())
}
}
