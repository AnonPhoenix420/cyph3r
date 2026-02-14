# CYPH3R: Tactical Network Intelligence System

**CYPH3R** is a high-performance reconnaissance engine built in Go, designed for deep intelligence gathering and tactical server testing. It focuses exclusively on remote target transparency while maintaining complete local host privacy.

---

## ⚡ Core Capabilities

### 🔍 Deep Recon Intelligence
* **Reversible Targeting:** Automatically detects and resolves both Domain-to-IP and IP-to-Domain (Reverse DNS).
* **IP Enumeration:** Discovers and displays every associated IPv4 and IPv6 address for a target node.
* **Recursive DNS Spidering:** Maps authoritative name servers and resolves their specific IP addresses in real-time.
* **Geo-Intelligence:** Pulls granular geographic data including Organization/ISP, City, State, Country Code, and Postal Code.

### 🛡️ Tactical Probing
* **Signal Identification:** Scans for open ports with active [ACK/SYN] signaling verification.
* **Protocol Fingerprinting:** Automatically identifies standard service protocols (SSH, HTTP, HTTPS, DNS, MySQL, etc.).
* **Speed-Optimized:** Utilizes 1.5s dial timeouts for efficient scanning without triggering basic rate-limiters.

---

## 🏗️ Project Architecture

```text
cyph3r/
├── bin/              # Compiled binary artifacts
├── cmd/
│   └── cyph3r/       # Entry point (main.go)
├── internal/
│   ├── intel/        # DNS and Geo-API logic
│   ├── models/       # Data structures and type definitions
│   ├── output/       # HUD rendering and UI status logic
│   └── probes/       # Tactical port scanning and protocol detection
├── Makefile          # Build, repair, and install automation
└── Dockerfile        # Multi-stage containerization
```

🚀 Deployment & Usage
​Fast Install
```
make install
```

Basic Recon
```
cyph3r -target google.com
```
Reverse IP Recon
```
cyph3r -target 8.8.4.4
```

🛠️ Maintenance Commands

```text
Command Action

*make repair Forces a deep-clean of Go caches and resets modules.

*make build Compiles a fresh binary to the ./bin directory.

*make backup Creates a timestamped .tar.gz archive of the source code.

*make docker Builds a minimal Alpine-based Docker image.
```

🔐 Privacy & Security Policy
CYPH3R is hard-coded to ignore local system information. It does not gather, print, or transmit your hostname, local interface IPs, or internal network topology. All output is strictly limited to the specified remote target.
