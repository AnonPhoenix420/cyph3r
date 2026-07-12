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
  --target      Set target URL/IP (e.g. 192.168.1.1)
  -p            Set target port (default 80)
  --method      Verb (GET/POST) for tests
  -c            Concurrency level (threads/streams)
  -d            Duration of operations in seconds

[+] INTELLIGENCE / RECON:
  --monitor     Engage HUD monitor loop
  --phone       Execute standalone telephony lookup
  --test-integrity Engage validation suite
  --json        Format output as raw JSON matrix

[+] RESILIENCE / STRESS TESTING:
  --hulk        Engage extreme resilience stress testing
  --proto       Protocol mode: tcp, udp, or http
  
[+] SYSTEM:
  -v            Enable full logging debug tracing
  --help        Display this menu

[!] EXAMPLES:
  Recon: ./cyph3r --target google.com
  Stress: ./cyph3r --target 192.168.1.50 -p 443 --hulk --proto tcp -c 500 -d 30
`)
	os.Exit(0)
}
