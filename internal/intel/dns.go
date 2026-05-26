package intel

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// DNSRecord holds the structured results of a DNS lookup operation
type DNSRecord struct {
	Type    string   `json:"type"`
	Records []string `json:"records"`
	Error   string   `json:"error,omitempty"`
}

// DNSResult bundles all gathering metrics for a target domain
type DNSResult struct {
	Domain     string               `json:"domain"`
	A          []string             `json:"a,omitempty"`
	AAAA       []string             `json:"aaaa,omitempty"`
	MX         []string             `json:"mx,omitempty"`
	TXT        []string             `json:"txt,omitempty"`
	CNAME      string               `json:"cname,omitempty"`
	Subdomains map[string][]string  `json:"subdomains,omitempty"`
}

// RunDNSLookup performs standard DNS queries against the target domain
func RunDNSLookup(ctx context.Context, domain string) (*DNSResult, error) {
	result := &DNSResult{
		Domain:     domain,
		Subdomains: make(map[string][]string),
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Helper for concurrent lookups
	lookup := func(lookupType string, fn func()) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				fn()
			}
		}()
	}

	// 1. A Records
	lookup("A", func() {
		if ips, err := net.LookupIP(domain); err == nil {
			var aRecords []string
			for _, ip := range ips {
				if ip.To4() != nil {
					aRecords = append(aRecords, ip.String())
				}
			}
			mu.Lock()
			result.A = aRecords
			mu.Unlock()
		}
	})

	// 2. AAAA Records
	lookup("AAAA", func() {
		if ips, err := net.LookupIP(domain); err == nil {
			var aaaaRecords []string
			for _, ip := range ips {
				if ip.To4() == nil && ip.To16() != nil {
					aaaaRecords = append(aaaaRecords, ip.String())
				}
			}
			mu.Lock()
			result.AAAA = aaaaRecords
			mu.Unlock()
		}
	})

	// 3. MX Records
	lookup("MX", func() {
		if mxRecords, err := net.LookupMX(domain); err == nil {
			var mx []string
			for _, record := range mxRecords {
				mx = append(mx, fmt.Sprintf("%s (Pref: %d)", record.Host, record.Pref))
			}
			mu.Lock()
			result.MX = mx
			mu.Unlock()
		}
	})

	// 4. TXT Records
	lookup("TXT", func() {
		if txtRecords, err := net.LookupTXT(domain); err == nil {
			mu.Lock()
			result.TXT = txtRecords
			mu.Unlock()
		}
	})

	// 5. CNAME Record
	lookup("CNAME", func() {
		if cname, err := net.LookupCNAME(domain); err == nil {
			// net.LookupCNAME often returns the domain itself if no CNAME exists
			if strings.TrimSuffix(cname, ".") != domain {
				mu.Lock()
				result.CNAME = cname
				mu.Unlock()
			}
		}
	})

	wg.Wait()
	return result, nil
}

// BruteForceSubdomains reads a wordlist and tests for valid subdomains concurrently
func BruteForceSubdomains(ctx context.Context, domain string, wordlistPath string, workers int) (map[string][]string, error) {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer file.Close()

	subdomainsChan := make(chan string, workers)
	resultsChan := make(chan struct {
		sub string
		ips []string
	}, workers)

	var wg sync.WaitGroup
	foundSubdomains := make(map[string][]string)
	var resultsMu sync.Mutex

	// Start worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for subPrefix := range subdomainsChan {
				select {
				case <-ctx.Done():
					return
				default:
					target := fmt.Sprintf("%s.%s", subPrefix, domain)
					// Use a custom resolver timeout to keep lookups snappy
					resolver := &net.Resolver{
						PreferGo: true,
						Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
							d := net.Dialer{Timeout: 2 * time.Second}
							return d.DialContext(ctx, network, "8.8.8.8:53")
						},
					}

					ips, err := resolver.LookupHost(ctx, target)
					if err == nil && len(ips) > 0 {
						resultsChan <- struct {
							sub string
							ips []string
						}{sub: target, ips: ips}
					}
				}
			}
		}()
	}

	// Result collector
	doneCollecting := make(chan struct{})
	go func() {
		for res := range resultsChan {
			resultsMu.Lock()
			foundSubdomains[res.sub] = res.ips
			resultsMu.Unlock()
		}
		close(doneCollecting)
	}()

	// Feed workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}
		select {
		case <-ctx.Done():
			break
		case subdomainsChan <- text:
		}
	}

	close(subdomainsChan)
	wg.Wait()
	close(resultsChan)
	<-doneCollecting

	return foundSubdomains, scanner.Err()
}

// PrintDNSReport formats and dumps the reconnaissance data cleanly to stdout
func PrintDNSReport(res *DNSResult) {
	fmt.Printf("\n[\033[34m*\033[0m] DNS Intelligence Report for: \033[1;36m%s\033[0m\n", res.Domain)
	fmt.Println(strings.Repeat("-", 50))

	if res.CNAME != "" {
		fmt.Printf("\033[33mCNAME:\033[0m %s\n", res.CNAME)
	}

	if len(res.A) > 0 {
		fmt.Println("\033[32mIPv4 Address Records (A):\033[0m")
		for _, ip := range res.A {
			fmt.Printf("  └─ %s\n", ip)
		}
	}

	if len(res.AAAA) > 0 {
		fmt.Println("\033[32mIPv6 Address Records (AAAA):\033[0m")
		for _, ip := range res.AAAA {
			fmt.Printf("  └─ %s\n", ip)
		}
	}

	if len(res.MX) > 0 {
		fmt.Println("\033[32mMail Exchange Records (MX):\033[0m")
		for _, mx := range res.MX {
			fmt.Printf("  └─ %s\n", mx)
		}
	}

	if len(res.TXT) > 0 {
		fmt.Println("\033[32mText Records (TXT):\033[0m")
		for _, txt := range res.TXT {
			fmt.Printf("  └─ %s\n", txt)
		}
	}

	if len(res.Subdomains) > 0 {
		fmt.Printf("\033[32mDiscovered Subdomains (%d):\033[0m\n", len(res.Subdomains))
		for sub, ips := range res.Subdomains {
			fmt.Printf("  └─ \033[1m%s\033[0m -> [%s]\n", sub, strings.Join(ips, ", "))
		}
	}
	fmt.Println()
}
