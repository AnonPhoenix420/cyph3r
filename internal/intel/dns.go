package intel

import "net"

func LookupNodes(target string) ([]string, []string) {
	var ips, nss []string
	addr, _ := net.LookupHost(target)
	ips = addr
	nsRecords, _ := net.LookupNS(target)
	for _, ns := range nsRecords {
		nss = append(nss, ns.Host)
	}
	return ips, nss
}
