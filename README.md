```
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / / / / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM

 ⚡ v2.6 [STABLE] // Wireframe HUD Edition
  ───────────────────────────────────────
```
![Version](https://img.shields.io/badge/Version-2.6--STABLE-cyan?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-magenta?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-blue?style=for-the-badge)

 ⚠️ **IMPORTANT LEGAL NOTICE:** Use of this software is strictly governed by the [Acceptable Use Policy (POLICY.md)](./POLICY.md). By cloning, building, or running CYPH3R, you agree to use this tool exclusively for authorized auditing, resilience testing, and educational purposes on systems you own or have explicit written permission to test. Unauthorized usage violates federal and international law.


# 🧠 Overview

CYPH3R is a professional-grade network reconnaissance, monitoring, and infrastructure resilience suite built in Go. It operates on a "Zero-Key" philosophy, providing deep OSINT (ISP, Geo, Metadata), multi-protocol connectivity testing, and an advanced 7-vector stress suite without requiring external API subscriptions.

# 🚀 CORE CAPABILITIES

- **Node Intelligence:** Automated registrar, ISP, and geographic coordinate mapping.

- **Accelerated Port Scanner:** Concurrent TCP scanning of common services (SSH, HTTP, DBs, etc.).

- **Live HUD Feed:** Real-time latency and status monitoring with a wireframe terminal UI.

- **Multi-Vector Stress Suite:** 7 integrated resilience engines (HULK, Slowloris, SYN Flood, Wrk Benchmarking, RUDY, HTTP/2 Rapid Reset, WebSocket Exhaustion).

- **Phone Decryption:** Metadata lookup for international phone vectors.

- **Go 1.23 Native:** Optimized for the stable Go 1.23 runtime with zero external binary dependencies.

# 📂 Architecture Mapping

```
cyph3r/
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── install.sh
├── uninstall.sh
├── cmd/cyph3r/main.go
├── internal/
│   ├── models/models.go
│   ├── intel/ (intel.go, dns.go)
│   ├── probes/ (probes.go, scanner.go)
│   ├── stress/ (stress.go)
│   └── output/ (banner.go, colors.go, pulse.go, render.go, status.go)
```

CYPH3R uses a modular internal structure to ensure high-speed execution and zero dependency clashing:

* `cmd/cyph3r/`: The primary CLI entry point.

* `internal/intel/`: OSINT logic for IP/Domain and Phone metadata.

* `internal/output/`: The HUD system (Split into Banners, Colors, and Status).

* `internal/probes/`: The network engine (TCP/UDP/HTTP/ACK socket logic).

* `internal/stress/`: The multi-vector resilience and benchmarking suite.


# 🛠️ Installation & Self-Repair

CYPH3R includes a built-in Self-Repair and cross-compilation system via Makefile. This is the recommended way to install to ensure your go.sum and dependencies are perfectly synced.

1. Requirements

Go: 1.23+

Make: For automated building and cross-compilation.

2. Build Process

Open your terminal in the project root and run:
bash

```make```

This syncs dependencies, cleans caches, and builds the production binary.


# 🛠️ CYPH3R Installation

```
git clone https://github.com/AnonPhoenix420/cyph3r.git
```
```cd cyph3r```
```
go mod tidy
go build -o cyph3r ./cmd/cyph3r
```

# 🧹⚙️🗑️ Advanced Maintenance (Makefile)

Self-Repair / Sync: make repair (Cleans cache and forces dependency sync)

Cross-Compile All Platforms: make cross-compile (Builds Linux, Windows, & macOS binaries into dist/)

Uninstall/Clean: make clean (Removes binaries, dist/ folder, and clears build cache)


Automated Install (Linux/macOS Only)

```
chmod +x install.sh
./install.sh
```

*SINGLE STEP INSTRUCTIONS

Install Go (The Language)

CYPH3R requires Go 1.23 or higher.

Parrot 🦜 OS 🖥️ Termux 📱

```
wget https://go.dev/dl/go1.23.5.linux-arm64.tar.gz
```
```
#tar -C /usr/local -xzf go1.23.5.linux-arm64.tar.gz
```
```
#echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
```go version```

you should get :

go version go1.23.5 linux/arm64

For Linux (Ubuntu/Debian/Kali)

Run these commands in your terminal:

```sudo apt update```

```
sudo apt install golang -y
```


Verify with: go version

For macOS
If you have Homebrew:

```brew install go```


For Windows
Download the MSI installer from ```go.dev/dl```

Run the installer and follow the prompts.

Restart your terminal/PowerShell.

Verify the Version

Once you have run the code above, type this to confirm it worked:

```go version```

Expected Output:

go version go1.23.0 linux/amd64


Install Build Tools:

CYPH3R uses a Makefile to handle automated building and cross-compilation.

For Linux/macOS
Most systems have this, but if not:

Linux:

```sudo apt install build-essential -y```

macOS:

```xcode-select --install```

For Windows
Windows doesn't have make by default. 

You have two choices:

The Easy Way: Skip make and just run

```
go build -o cyph3r.exe ./cmd/cyph3r manually
```

The Pro Way: Install Chocolatey and run

```choco install make```

# 🚀 Tool Usage Guide
CYPH3R contains multiple primary tools packed into a single binary. Here is how to use each.

🛡️ Tool 1: Target Intelligence & OSINT Reconnaissance
Extract comprehensive intelligence reports, unmasked real IPs, discovered subdomains, harvested emails, phone vectors, and digital footprints.

Command:

./cyph3r --target <host> --osint

Example:

./cyph3r --target google.com --osint

Accelerated Tactical Port Scan

Command:

./cyph3r --target <host> --scan

Example:

./cyph3r --target google.com --scan

📡 Tool 2: Continuous Monitor (HUD Feed) 
Track the uptime and latency of a target over time. Perfect for stress testing or uptime verification.

Command:

./cyph3r --target <host> --proto <type> --monitor

Example:

./cyph3r --target google.com --proto https --monitor --interval 5s

Protocols supported: tcp, udp, http, https, ack, ping.

📱 Tool 3: Phone Metadata Lookup 
Validate international phone numbers and retrieve regional/carrier metadata.

Command:

./cyph3r --phone <number>

Example:

./cyph3r --phone +14155552671

⚡ Tool 4: Multi-Vector Stress & Benchmarking Suite
Evaluate infrastructure resilience using 7 high-performance concurrency engines.

HULK HTTP Flood:
./cyph3r --target <host> --hulk

Slowloris Header Exhaustion:
./cyph3r --target <host> --slowloris

Layer 4 SYN/State Flood:
./cyph3r --target <host> --synflood

Wrk-Style Benchmark (with RPS & Latency tracking):
./cyph3r --target <host> --wrk -c 100 -d 30

RUDY Slow-POST Exhaustion:
./cyph3r --target <host> --rudy

HTTP/2 Rapid Reset Multiplexing:
./cyph3r --target <host> --h2

WebSocket Connection & Frame Pool Exhaustion:
./cyph3r --target <host> --ws




Wrk-Style Benchmark (with RPS & Latency tracking):

```
./cyph3r --target <host> --wrk -c 100 -d 30
```

RUDY Slow-POST Exhaustion:

```./cyph3r --target <host> --rudy```

HTTP/2 Rapid Reset Multiplexing: 

```./cyph3r --target <host> --h2```

WebSocket Connection & Frame Pool Exhaustion: 

```./cyph3r --target <host> --ws```

UNINSTALL

Since CYPH3R v2.6 is a modular Go tool, it doesn't scatter files all over your system like a standard installer might. However, to keep your workspace pristine, a dedicated uninstaller is included.

# 🗑️ Uninstallation

To remove the binary, local distribution artifacts, and clean global pathways:

```
chmod +x uninstall.sh
./uninstall.sh
```

💻⌨️🧑‍💻 HAPPY HACKING 📀🖥️🖱️


# 📜 MIT License

Copyright (c) 2026 AnonPhoenix420
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.


# ⚖️ Disclaimer
For authorized security testing and educational purposes only. Misuse of this tool is strictly the responsibility of the end user.


# READ POLICY.MD BEFORE USE !!

