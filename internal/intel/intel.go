package intel

import (
	"net"
	"os"
	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

func GetTargetIntel(target string) (models.IntelData, error) {
	var data models.IntelData
	
	// 1. Resolve Target IP
	ips, _ := net.LookupIP(target)
	if len(ips) > 0 {
		data.IP = ips[0].String()
	}

	// 2. Fetch Name Servers
	ns, _ := net.LookupNS(target)
	for _, nameserver := range ns {
		data.NameServers = append(data.NameServers, nameserver.Host)
	}

	// 3. Get Localhost Identity (Always included)
	hostname, _ := os.Hostname()
	data.LocalHost = hostname
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				data.LocalIPs = append(data.LocalIPs)
			}
		}
	}

	// Placeholder: In a real scenario, you'd call an API here to populate 
	// Org, Lat, Lon, Zip, etc. For now, we initialize the fields.
	data.Org = "Pending Query..." 
	return data, nil
}
