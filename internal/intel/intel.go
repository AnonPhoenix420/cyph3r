package intel

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
	"github.com/AnonPhoenix420/cyph3r/internal/probes"
)

func GetTargetIntel(input string) (models.IntelData, error) {
	var data models.IntelData
	data.TargetName = input
	data.NameServers = make(map[string][]string)

	ips, _ := net.LookupIP(input)
	for _, ip := range ips { 
		data.TargetIPs = append(data.TargetIPs, ip.String()) 
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() { 
		defer wg.Done()
		data.Subdomains = discoverSubdomains(input) 
	}()
	go func() { 
		defer wg.Done()
		data.Org = fetchWhois(input) 
	}()
	wg.Wait()

	data.NameServers["PORTS"] = probes.ScanPorts(input)
	return data, nil
}

func discoverSubdomains(domain string) []string {
	var found []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	subs := []string{"www", "mail", "vpn", "dev", "api", "ssh", "ftp"}

	for _, s := range subs {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			if _, err := net.LookupHost(sub + "." + domain); err == nil {
				mu.Lock()
				found = append(found, sub+"."+domain)
				mu.Unlock()
			}
		}(s)
	}
	wg.Wait()
	return found
}

// Internalized WHOIS logic - No external libraries needed
func fetchWhois(domain string) string {
	server := "whois.iana.org"
	if strings.HasSuffix(domain, ".ir") { server = "whois.nic.ir" }
	
	conn, err := net.DialTimeout("tcp", server+":43", 5*time.Second)
	if err != nil { return "CONNECTION_TIMEOUT" }
	defer conn.Close()

	fmt.Fprintf(conn, domain+"\r\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, "registrar:") || strings.Contains(line, "organization:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 { 
				return strings.TrimSpace(parts[1]) 
			}
		}
	}
	return "DATA_HIDDEN"
}
