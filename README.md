

 [ CYPH3R : Network Intel & Diagnostics ]
 
           [ Hacker at Work ]
               _________
              |  _____  |
              | |     | |
              | | CLI | |
              | |_____| |
              |_________|
                  ||
      (•_•)       ||
     <|   |>======||
      / \         ||
                 /__\

            

------------------------------------------

## 🧠 Overview

**CYPH3R** is a high-performance network diagnostic and intelligence tool written in Go. Unlike basic scanners, CYPH3R combines deep **OSINT (Open-Source Intelligence)** with a multi-protocol **Probing Engine**. 

It is designed to be "Zero-Config"—no API keys, no passwords, and no complex setup required.

## ✨ Updated Features
- **Deep Intel:** Real-time lookup of ISP Handlers, Organization names, and Reverse DNS.
- **Precision GeoIP:** Returns Country, City, and **Zip Code** with GPS coordinates.
- **Maps Integration:** Generates a direct Google Maps link for the target's physical location.
- **Multi-Protocol Engine:** - **TCP/ACK:** Handshake validation.
  - **UDP:** Connectionless probing.
  - **HTTP/HTTPS:** Full status code and latency reporting.
- **Phone Intelligence:** International number validation and carrier region metadata.
- **Worker Pool:** High-speed concurrency support for stress testing.

---

## 🛠️ Installation

### 1. Requirements
- **Go 1.23+** (Required for modern networking libraries)
- **Git**

### 2. Fast-Track Build (Recommended)
Clone the repository and run the automated installer:

```bash
git clone https://github.com/AnonPhoenix420/cyph3r.git
cd cyph3r
chmod +x install.sh
./install.sh
```

Manual Compilation
If you prefer to build manually:

```

go mod tidy
go build -o cyph3r ./cmd/cyph3r

```

🚀 Usage Guide
Deep Intelligence Lookup
Get ISP, Maps, Zip, and WHOIS for a domain or IP:

```
./cyph3r --target example.com --intel

```
Protocol Probing
Test specific ports using different methods:
```

# TCP Probe (Default)
./cyph3r --target 1.1.1.1 --port 53 --proto tcp

# UDP Probe
./cyph3r --target 8.8.8.8 --port 53 --proto udp

# HTTP Latency Check
./cyph3r --target google.com --proto https

```
Phone Metadata
Verify international number status:
```
./cyph3r --phone +14155552671

```

📁 Project Structure

```

cyph3r/
├── cmd/cyph3r/main.go     # The Commander (Entry Point)
├── internal/
│   ├── intel/             # The Brain (GeoIP, Maps, WHOIS, Phone)
│   ├── probes/            # The Hands (TCP, UDP, ACK, HTTP)
│   └── output/            # The Voice (Terminal UI & Colors)
├── go.mod                 # Dependency Manifest
└── install.sh             # Automation Script

```

⚖️ Disclaimer
For Educational and Professional Services Use Only. The creator is NOT responsible for misuse. Always ensure you have explicit permission before testing any network or service you do not own.


Happy Hacking — Responsibly 🧠🚀



🚨🚨 TECHNICAL ERRORS 🚨🚨



How to generate the new go.sum "if you are having problems and a go.sum was automatically generated
Run these three commands in your terminal inside the cyph3r directory:

# 1. Remove any old, conflicting sum files
```
rm -f go.sum
```
# 2. Download and verify the exact versions required by the new code
```
go mod tidy
```
# 3. Verify the hashes are locked
```
go mod verify
```
# 1. Initialize the module (if not already done)
```
go mod init github.com/AnonPhoenix420/cyph3r 2>/dev/null || true
```
# 2. Fetch the specific OSINT and Phone libraries
```
go get github.com/nyaruka/phonenumbers
go get github.com/prometheus/client_golang
```
# 3. Clean and verify the go.sum file
```
go mod tidy
```
# 4. Build the executable
```
go build -o cyph3r ./cmd/cyph3r
MIT License

Copyright (c) 2026 AnonPhoenix420

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
