# CYPH3R: Tactical Network Intelligence & Resilience Suite

**CYPH3R** is a high-performance reconnaissance, performance benchmarking, and infrastructure resilience suite built in Go, designed for deep intelligence gathering and tactical server testing. It focuses exclusively on remote target transparency while maintaining complete local host privacy.

---

## ⚡ Core Capabilities

### 🔍 Deep Recon Intelligence
* **Reversible Targeting:** Automatically detects and resolves both Domain-to-IP and IP-to-Domain (Reverse DNS).
* **IP Enumeration:** Discovers and displays every associated IPv4 and IPv6 address for a target node.
* **Recursive DNS Spidering:** Maps authoritative name servers and resolves their specific IP addresses in real-time.
* **Geo-Intelligence:** Pulls granular geographic data including Organization/ISP, City, State, Country Code, and Postal Code.

### 🛡️ Tactical Probing & Stress Suite
* **Signal Identification:** Scans for open ports with active [ACK/SYN] signaling verification.
* **Protocol Fingerprinting:** Automatically identifies standard service protocols (SSH, HTTP, HTTPS, DNS, MySQL, etc.).
* **Multi-Vector Stress Suite:** 7 integrated resilience engines (HULK, Slowloris, SYN Flood, Wrk Benchmarking, RUDY, HTTP/2 Rapid Reset, WebSocket Exhaustion).
* **Phone Decryption:** Metadata lookup for international phone vectors.

---

## 🏗️ Project Architecture

```text
cyph3r/
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── POLICY.md
├── install.sh
├── uninstall.sh
├── cmd/
│   └── cyph3r/       # Entry point (main.go)
└── internal/
    ├── intel/        # DNS, Geo, and Phone metadata logic
    ├── models/       # Data structures and type definitions
    ├── output/       # HUD rendering, banners, and status logic
    ├── probes/       # Tactical port scanning and protocol detection
    └── stress/       # Multi-vector resilience and benchmarking engines
```

# 🚀 Deployment & Usage

Fast Build & Sync

```make```

Basic Recon

```cyph3r --target google.com```

Full Recon & Port Scan

```cyph3r --target google.com --scan```

Multi-Vector Resilience Testing

```
cyph3r --target <host> --wrk -c 100 -d 30
```

# 🛠️ Maintenance & Cross-Compilation Commands



```make``` — Syncs dependencies 
(```go mod tidy```), cleans caches, and builds the production binary.

```make cross-compile``` — Builds standalone binaries for Linux, Windows, and macOS into the dist/ directory.

```make clean``` — Completely wipes out binaries, dist/ folder, and Go build caches.


# 🔐 Privacy, Security & Legal Policy

Use of this software is strictly governed by the Acceptable Use Policy (POLICY.md). CYPH3R is hard-coded to ignore local system information. It does not gather, print, or transmit your hostname, local interface IPs, or internal network topology. All output is strictly limited to authorized remote targets.
