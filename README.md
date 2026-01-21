```text
  
  _____    __     __   ______    _    _    _____   ______
 / ____|  \ \   / /  | ___ \  | |  | |  |___  |  | ___ \
| |       \ \_/ /   | |_/ /  | |__| |    / /   | |_/ /
| |        \   /    |  __/   |  __  |  |_ \   |  _  \
| |____     | |     | |      | |  | |  ___) |  | | \ \
 \_____|    |_|     \_|      |_|  |_|  |____/   \_|  \_|

```

⚠️ Disclaimer  
For Educational and Professional Services Use Only  
The creator is NOT responsible for misuse of this tool  

Always ensure you have explicit permission before testing any network, host, or service you do not own.

---

## 🧠 Overview

**CYPH3R** is a modular, Go-based network diagnostics and monitoring utility designed for learning, troubleshooting, and professional network validation.

It supports continuous monitoring, downtime tracking, and multiple protocols while remaining lightweight and Termux-friendly.

Runs cleanly on:
- Parrot OS
- Linux (x86_64 & ARM64)
- Termux (Android / aarch64)

---

## ✨ Features

- TCP connectivity testing  
- UDP probing (local echo friendly)  
- HTTP / HTTPS status & latency checks  
- Continuous monitor mode (tracks downtime)  
- GeoIP / DNS resolution (IP or localhost)  
- Optional phone number metadata lookup  
- Colored terminal output (UP = blue, DOWN = red)  
- JSON output mode for scripting & automation  
- Version & author metadata flags  
- Single static binary build support  
- ARM64 / aarch64 compatible  

---

## 🚀 Usage & Help

To see all available options and usage info, run:

```bash
./cyph3r --help
```
or
```
./cyph3r -h
```
or just

```
./cyph3r


```
Example output:




CYPH3R: Network Diagnostics Utility

Usage:
  cyph3r [options]

Options:
  --target        Target host or IP (default: localhost)
  --port          Port number (default: 80)
  --proto         Protocol: tcp | udp | http | https | dns (default: tcp)
  --geoip         Lookup GeoIP and ASN info for the target
  --phone         Show phone number info (region, type, validity)
  --json          Print output in JSON
  --monitor       Continuously monitor the target
  --interval      Check interval in seconds (default: 5)
  --version       Show program version info
  --portscan      Scan ports on IP/localhost (default range 1-1024)
  --scanstart     Port scan range start (default: 1)
  --scanend       Port scan range end (default: 1024)
  -h, --help      Show this help and exit




📦 Requirements

Go 1.23.0

git

Supported Platforms:

Parrot OS

Linux (Debian-based recommended)

Termux on Android



🛠️ Installation
🔹 Parrot OS / Linux


```

curl -LO https://go.dev/dl/go1.23.0.linux-arm64.tar.gz

```

```

sudo tar -C /usr/local -xzf go1.23.0.linux-arm64.tar.gz

```
Verify:
```
go version

```

🔹 Termux (Android)

```

apt update && apt upgrade

```

```
apt install git
```

📥 Clone Repository

```

git clone https://github.com/AnonPhoenix420/cyph3r.git
```
```
cd cyph3r
```
```
git clone https://github.com/nyaruka/phonenumbers.git
```

🔧 Build & Install

Local build:

```

go mod tidy

```
```
go build -o cyph3r ./cmd/cyph3r

```

System-wide install:

```

sudo install -m 755 cyph3r /usr/local/bin/cyph3r
```
🚀 Usage Examples

Basic TCP test
```
./cyph3r --target localhost --port 80 --proto tcp
```
GeoIP Lookup
```
./cyph3r --target example.com --geoip

```
Monitor Mode

```
./cyph3r --target example.com --port 443 --proto https --monitor --interval 5
```
UDP Test
```
./cyph3r --target 127.0.0.1 --port 53 --proto udp

```
HTTP / HTTPS
```

./cyph3r --target example.com --port 443 --proto https

```
JSON Output

```

./cyph3r --target example.com --proto https --json

```
Phone Metadata

```
./cyph3r --phone +14155552671
```




🧪 Monitor Mode Explained

When ```--monitor``` is enabled:

• Tool runs indefinitely

• Tracks downtime

• Reports recovery time

• Stops only with Ctrl+C






📁 Project Structure


cyph3r/
├── cmd/cyph3r/main.go
├── internal/
│   ├── netcheck/
│   ├── geo/
│   ├── phone/
│   ├── output/
│   └── version/
├── install.sh
├── go.mod
├── LICENSE
└── README.md





⚖️ License

Educational & Professional Use Only
THE SOFTWARE IS PROVIDED "AS IS"




🧠 Notes

No ```--i-own-this``` flag required

No capped duration limits

Designed for diagnostics & learning

Do not test what you do not own

Happy hacking — responsibly 🧠🚀




