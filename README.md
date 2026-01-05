# cyph3r

⚠️ **For Educational and Professional Services Use Only**
⚠️ **The creator is NOT responsible for misuse of this tool**

`cyph3r` is a modular network diagnostics and testing utility written in Go. It supports TCP, UDP, HTTP, and HTTPS checks, optional phone metadata lookup, GeoIP/DNS resolution, JSON output, and runs cleanly on **Parrot OS**, **Linux**, and **Termux (Android)**.

---

## ✨ Features

* TCP / UDP connectivity testing
* HTTP / HTTPS status & latency checks
* GeoIP & DNS resolution (IP or localhost)
* Optional phone number metadata lookup
* Colored terminal output (UP = blue, DOWN = red)
* JSON output mode for scripting
* ARM64 (aarch64) compatible
* Works with sudo (installer)

---

## 📦 Requirements

### Supported Platforms

* Parrot OS (Linux)
* Termux on Android (aarch64)
* Standard Linux PCs

### Required Packages

* Go **1.20 or newer**
* git

---

## 🛠️ Installation (Parrot OS on Termux)

### 1️⃣ Install dependencies

```bash
pkg update && pkg upgrade
pkg install git golang
```

Verify Go:

```bash
go version
```

---

### 2️⃣ Clone the repository 

```
git clone https://github.com/AnonPhoenix420/Cyph3r
```
```
cd cyph3r
```
```
git clone https://github.com/nyaruka/phonenumbers v1.0.59
```

### 3️⃣ Build and install

```bash
go mod tidy
go build -o cyph3r ./cmd/cyph3r
```

(Optional system-wide install – requires sudo on Linux PCs)

```bash
sudo install -m 755 cyph3r /usr/local/bin/cyph3r
```

On **Termux**, just run it locally:

```bash
./cyph3r
```

---

## 🚀 Usage

### Basic TCP test

```bash
cyph3r --target localhost --port 80 --proto tcp
```

### UDP test

```bash
cyph3r --target 127.0.0.1 --port 53 --proto udp
```

### HTTP / HTTPS test

```bash
cyph3r --target example.com --port 443 --proto https
```

### JSON output mode

```bash
cyph3r --target example.com --proto https --json
```

### Phone metadata lookup (optional)

```bash
cyph3r --phone +14155552671
```

### Version info

```bash
cyph3r --version
```

---

## 🧪 Single-File / Portable Build

To create a **single static binary** (recommended for Termux):

```bash
go build -o cyph3r ./cmd/cyph3r
```

You can now copy `cyph3r` anywhere and run it without the source tree.

---

## 📁 Project Structure

```
cyph3r/
├── cmd/cyph3r/main.go
├── internal/
│   ├── netcheck/   # TCP / UDP / HTTP logic
│   ├── geo/        # GeoIP / DNS lookup
│   ├── phone/      # Phone metadata
│   ├── output/     # Color + JSON output
│   └── version/    # Version metadata
├── install.sh
├── go.mod
├── LICENSE
└── README.md
```

---

## ⚖️ License & Disclaimer

```
Educational & Professional Use Only

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.
THE CREATOR IS NOT RESPONSIBLE FOR MISUSE OR DAMAGE CAUSED BY THIS TOOL.
```

---

## 📸 Example Output

```
==============================
 cyph3r — Network Utility
 Educational use only ⚠️
==============================
🌍 Resolved IPs:
 - 127.0.0.1
[UP] localhost:80 (12 ms)
```

---

## 🧠 Notes

* No special flags like `--i-own-this` are required
* Designed for diagnostics, learning, and professional testing
* Always have permission before testing networks you do not own

---

Happy hacking — responsibly 🧠🚀

