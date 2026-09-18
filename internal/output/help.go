package output

import (
	"fmt"
	"os"
)

func DisplayHelp() {
	fmt.Println(`
  ______      ____  __  __ _____ ____
 / ____/_  __/ __ \/ / / /|__  // __ \
/ /   / / / / /_/ / /_/ /  /_ </ /_/ /
\____/\__, /_/   /_/ /_/ /____/_/ |_|
     
     NETWORK_INTEL_SYSTEM // RESILIENCE SUITE
     ---------------------------------------
     
[!] USAGE: cyph3r --target <host> [options]

[+] OPERATIONAL VECTORS:
  --target      Set target URL, domain, or IP (e.g. example.com, 192.168.1.1)
  -p            Set target port (Auto-detected from URL or defaults to 80/443)
  -c            Concurrency pool size (Default: 500 sockets/workers)
  -d            Test duration in seconds (Default: 0 for infinite/continuous)
  --interval    HUD monitor ping interval (Default: 2s)
  --ghost       Engage Ghost Mode (SOCKS5/Tor stealth tunneling with jitter & timeout controls)

[+] INTELLIGENCE & RECON:
  --osint       Extract comprehensive intelligence report, unmasked IPs, & footprints
  --scan        Engage accelerated TCP port scanner and service probes
  --monitor     Engage live HUD connection monitor loop
  --phone       Execute standalone international phone metadata lookup

[+] STRESS & BENCHMARKING SUITE:
  --hulk        Engage continuous resilience HULK stress engine
  --slowloris   Engage Slowloris slow-rate exhaustion engine
  --synflood    Engage Layer 4 TCP state/SYN flood engine
  --wrk         Engage Wrk-style high-performance benchmarking engine
  --rudy        Engage RUDY slow-POST exhaustion engine
  --h2          Engage HTTP/2 rapid reset stream engine
  --ws          Engage WebSocket connection & frame exhaustion engine
  --proto       Wire protocol mode: tcp or udp (for monitor)

[+] SYSTEM:
  --help        Display this menu

[!] EXAMPLES:
  OSINT Recon:  ./cyph3r --target google.com --ghost --osint
  Port Scan:    ./cyph3r --target google.com --ghost --scan
  Phone Lookup: ./cyph3r --phone +14155552671
  Stress Test:  ./cyph3r --target 192.168.1.50 --ghost --hulk -c 500 -d 30
`)
	os.Exit(0)
}
