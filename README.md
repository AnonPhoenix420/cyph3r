# CYPH3R

```
  _____    __     __   ______    _    _    _____   ______
 / ____|  \ \   / /  | ___ \  | |  | |  |___  |  | ___ \
| |       \ \_/ /   | |_/ /  | |__| |    / /   | |_/ /
| |        \   /    |  __/   |  __  |  |_ \   |  _  \
| |____     | |     | |      | |  | |  ___) |  | | \ \
 \_____|    |_|     \_|      |_|  |_|  |____/   \_|  \_|
```

## ⚠️ Disclaimer

For Educational and Professional Services Use Only. Always ensure you have explicit permission before testing any network, host, or service you do not own. The author is not responsible for misuse.

## 🧠 Overview

**CYPH3R** is a modular Go-based network diagnostics and load testing platform. Features include:

* HTTP/HTTPS/TCP/ICMP testing
* Mixed scenario load runs
* WHOIS, DNS, ASN & CIDR expansion
* ICMP jitter & packet loss metrics
* Live terminal dashboard
* JSON & Prometheus metrics
* Grafana-ready endpoints
* ARM64/Termux friendly

## 🖥 Supported Platforms

* Parrot OS (Linux)
* Linux x86_64 / ARM64
* Termux (Android / aarch64)

## 📦 Requirements

* Go 1.22+ (1.23 recommended)
* git

## 🛠 Installation

### Clone the repository

```bash
git clone https://github.com/AnonPhoenix420/cyph3r
```
```
cd cyph3r
```

### Build

```bash

go mod tidy
```
```
go build -o cyph3r main.go intel.go

```

(Optional system-wide install on Linux)

```bash
sudo install -m 755 cyph3r /usr/local/bin/cyph3r
```

### ICMP Raw Sockets (Optional)

Requires root or `cap_net_raw`:

```bash
sudo setcap cap_net_raw+ep ./cyph3r
```

## 🚀 Usage

### Basic HTTP load test

```bash
./cyph3r -target example.com -proto http -rps 100 -duration 30s
```

### HTTPS POST with payload

```bash
./cyph3r -target example.com -proto https -method POST -payload '{"ping":"pong"}'
```

### ICMP jitter & packet loss

```bash
./cyph3r -target 8.8.8.8 -proto icmp
```

### Mixed scenario (HTTP + ICMP + TCP)

```bash
./cyph3r -target example.com -scenario mixed
```

### ASN fan-out testing

```bash
./cyph3r -target example.com -asn-fanout -rps 100
```

### WHOIS / DNS / ASN intel only

```bash
./cyph3r -target example.com -whois -dns -asn
```

### JSON output

```bash
./cyph3r -target example.com -json
```

## 📊 Monitoring & Dashboards

Prometheus metrics available at `http://localhost:2112/metrics`.
Web UI status endpoint at `http://localhost:2112/status`.
Metrics include latency, counts, ICMP loss/jitter, protocol/scenario/target labels.

## 🧪 Ramp Profiles & Thresholds

```bash
./cyph3r -target example.com -rps 500 -ramp 20s -latency 2s -failrate 0.05
```

Alerts if failure rate or p95 latency exceed thresholds.

## 📁 Project Structure

```
cyph3r/
├── main.go        # Scheduler, scenarios, load engine
├── intel.go       # WHOIS, DNS, ASN, CIDR, TLS, ICMP
├── output/        # Terminal UI & colors
├── go.mod
├── go.sum
└── README.md
```

## ⚖️ License

Educational & Professional Use Only. THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.

## 🧠 Notes

* No artificial caps, telemetry, or spyware
* Designed for diagnostics, learning, and authorized testing
* Use responsibly and legally
