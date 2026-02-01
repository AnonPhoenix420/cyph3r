# 
```text
   ______      ____  __  __ _____ ____ 
  / ____/_  __/ __ \/ / / /|__  // __ \
 / /   / / / / /_/ / /_/ /  /_ </ /_/ /
/ /___/ /_/ / ____/ __  / ___/ / _, _/ 
\____/\__, /_/   /_/ /_/ /____/_/ |_|  
     /____/         NETWORK_INTEL_SYSTEM

  v2.6 [STABLE] // Wireframe HUD Edition
  ───────────────────────────────────────
```

---


## 🧠 Overview
---

CYPH3R is a professional-grade network reconnaissance and monitoring tool built in Go. It operates on a **"Zero-Key" philosophy**, providing deep OSINT (ISP, Geo, Metadata) and multi-protocol connectivity testing without requiring external API subscriptions.

---


## 📂 Architecture Mapping


---

#
```text
cyph3r/
├── cmd/
│   └── cyph3r/
│       └── main.go
├── internal/
│   ├── intel/
│   ├── output/
│   └── probes/
├── .gitignore       
├── LICENSE          
├── Dockerfile       
├── go.mod
├── go.sum
├── Makefile
├── install.sh
└── uninstall.sh
```

---


CYPH3R uses a modular internal structure to ensure high-speed execution and zero dependency clashing:

#
```text
* `cmd/cyph3r/`: The primary CLI entry point.
* `internal/intel/`: OSINT logic for IP/Domain and Phone metadata.
* `internal/output/`: The HUD system (Split into Banners, Colors, and Status).
* `internal/probes/`: The network engine (TCP/UDP/HTTP/ACK socket logic).
```


---


## 🛠️ Installation & Self-Repair

---


CYPH3R includes a built-in **Self-Repair** system via `Makefile`. This is the recommended way to install to ensure your `go.sum` and dependencies are perfectly synced.


---


### 1. Requirements
* **Go:** 1.23+
* **Make:** For automated building.


---


### 2. Build Process
Open your terminal in the project root and run:
bash
```
make repair
```

---


# This cleans the cache, resyncs dependencies, and compiles the binary


---


🛠️ CYPH3R Installation


---



****Clone the official repo****

Standard Build (Recommended)

```
git clone https://github.com/AnonPhoenix420/cyph3r.git
```
```
cd cyph3r
go mod tidy
go build -o cyph3r ./cmd/cyph3r

```

---


🧹⚙️🗑️
Advanced Maintenance (```Makefile```)
Self-Repair: ```make repair``` (Cleans cache and forces dependency sync)

Uninstall/Clean: ```make clean``` (Removes binary and clears build cache)


---


Automated Install (Linux/macOS Only)

---


```

chmod +x install.sh
./install.sh
```

---


****SINGLE STEP INSTRUCTIONS***

---


1. Install Go (The Language)
CYPH3R requires Go 1.23 or higher.


---


For Linux (Ubuntu/Debian/Kali)
Run these commands in your terminal:


```
sudo apt update
sudo apt install golang -y
```

---

Verify with: go version

---

For macOS
If you have Homebrew:
```
brew install go
```

---


For Windows
Download the MSI installer from go.dev/dl.

Run the installer and follow the prompts.

Restart your terminal/PowerShell.


---



****Run these commands one by one to remove old versions and install Go 1.23****

---


# 1. Download the Go 1.23.0 Archive
```
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
```
---


# 2. Remove any previous Go installation and extract the new one
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
```
---


# 3. Add Go to your PATH (Environment Variables)
```
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.zshrc
```
---


# 4. Refresh your terminal
```
source ~/.bashrc
```
---


**"IF YOU DON'T HAVE GO YET JUST DOWNLOAD IT FORM THE PATH ABOVE AND INSTALL IT YOURSELF***

---

```
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz

```
```
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
```
```
export PATH=$PATH:/usr/local/go/bin
```

---


****Verify the Version****

Once you have run the code above, type this to confirm it worked:
```
go version
```
---


Expected Output:
```go version go1.23.0 linux/amd64```

---


2. Install Build Tools
CYPH3R uses a ```Makefile``` to handle the "```Self-Repair```" and automated building features.


---


For Linux/macOS
Most systems have this, but if not:

---


Linux: 
```
sudo apt install build-essential -y

```

---

macOS: 
```
xcode-select --install
```

---

For Windows
Windows doesn't have make by default. You have two choices:


---


The Easy Way: Skip make and just run 
```
go build -o cyph3r.exe ./cmd/cyph3r manually.
```

---


The Pro Way: Install Chocolatey and run 
```
choco install make.
```

---



🚀 Tool Usage Guide
CYPH3R contains three primary "tools" in one binary. Here is how to use each.

---


🛡️ Tool 1: Target Intelligence (OSINT)
Retrieve ISP, Organization, City, Zip, and GPS coordinates for any IP or Domain.

---


Command: ```./cyph3r --target <host>```

---


Example: ./cyph3r --target 8.8.8.8


---



📡 Tool 2: Continuous Monitor (HUD Feed)
Track the uptime and latency of a target over time. Perfect for stress testing or uptime verification.


---


Command: ```./cyph3r --target <host> --proto <type> --monitor```


---


Example: ./cyph3r --target google.com --proto https --monitor --interval 5s


---


Protocols supported: ```tcp```, ```udp```, ```http```, ```https```, ```ack```, ```ping```.


---


📱 Tool 3: Phone Metadata Lookup
Validate international phone numbers and retrieve regional/carrier metadata.

Command: ```./cyph3r --phone <number>```

Example: ./cyph3r --phone +14155552671


---  


***UNINSTALL***

Since CYPH3R v2.6 is a modular Go tool, it doesn't scatter files all over your system like a standard installer might. However, to keep your workspace pristine, a dedicated uninstaller is a professional touch.


---


## 🗑️ Uninstallation
To remove the binary and clean build artifacts:

```
chmod +x uninstall.sh
./uninstall.sh
```

---


💻⌨️🧑‍💻 ```HAPPY HACKING``` 📀🖥️🖱️


---


📜 MIT License
Copyright (c) 2026 AnonPhoenix420

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.



---



⚖️ Disclaimer
For authorized security testing and educational purposes only. Misuse of this tool is strictly the responsibility of the end user.

