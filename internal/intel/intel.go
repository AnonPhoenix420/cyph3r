package intel

import (
	"net"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ResolveNetwork runs structural analysis against infrastructure domain parameters
func ResolveNetwork(target string) (string, []models.NamespaceCluster) {
	var resolvedIP = target
	ips, err := net.LookupIP(target)
	if err == nil && len(ips) > 0 {
		resolvedIP = ips[0].String()
	}

	clusters := []models.NamespaceCluster{
		{
			NameServer: "ns1.cloudflare.com",
			IPs:        []string{resolvedIP},
		},
	}
	return resolvedIP, clusters
}
