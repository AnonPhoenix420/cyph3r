# 🕵️‍♂️ Digital Privacy, Proxies & Tor: A Public Security & Anonymity Guide

In an era of deep packet inspection, automated tracking, and widespread telemetry, protecting your digital footprint during network reconnaissance and resilience testing is vital. This guide explores the mechanics of network anonymization, proxy routing, and how `Cyph3r` leverages these layers for secure operation.

# 🧭 Core Concepts: Proxies, SOCKS5, and Tor

Before running network diagnostics or audits, it is essential to understand how your traffic is exposed and how to shield it:

- **Standard IP Exposure:** When you send a request over the clear web, your Internet Service Provider (ISP) and target servers log your public IPv4/IPv6 address, exposing your physical location and network identity.

- **SOCKS5 Proxies:** Unlike basic HTTP proxies, SOCKS5 proxies handle any kind of TCP/UDP traffic at the transport layer. They can route raw socket data, DNS queries, and connections through an intermediate node, masking your real IP address.

- **The Onion Router (Tor):** Tor is a decentralized overlay network that routes your traffic through at least three encrypted relay nodes (Guard, Middle, and Exit node) globally. Each node only knows the IP address of the node immediately before and after it, making end-to-end tracking nearly impossible.

# 📦 Installation Guide: Tor, Torsocks & Proxychains

To route your system or command-line tools through Tor, you need the Tor daemon and wrapper utilities installed on your operating system.

# 🐧 For Linux (Debian, Ubuntu, Kali Linux, Parrot OS)

Run the following commands using `apt` and `sudo`:
```

sudo apt update
sudo apt install tor proxychains4 torsocks -y
```

# 📱 For Android (Termux)

If you are operating inside a mobile Termux environment, use `pkg`:

```
pkg update
pkg install tor proxychains-ng torsocks -y
```

# ⚙️ Service Control (Linux Systemd)

Make sure the Tor background service is running before attempting to tunnel connections:

```
sudo systemctl start tor
sudo systemctl enable tor
```

(Verify Tor is listening on its default local SOCKS port `127.0.0.1:9050:`)

```
ss -tulpn | grep 9050
```

# 🛠️ Tooling Overview: Proxychains & Torsocks

When running terminal applications that do not natively support proxy configurations, you can use wrapper utilities to force them through Tor:


1. **Torsocks** (Simple Single-Command Wrapper)
`torsocks` intercepts application socket calls and forces them through the Tor SOCKS5 proxy:

```
torsocks curl https://check.torproject.org/api/ip
```



2. **Proxychains** (Multi-Proxy Chain Routing)

`proxychains` allows you to route arbitrary applications through a sequence of proxies configured in `/etc/proxychains4.conf.`
To run a command through Proxychains:


# 👻 Integration with Cyph3r: Ghost Mode (`--ghost`)


Configuring external wrappers like `proxychains` or managing manual SOCKS bindings for heavy multi-threaded stress and OSINT engines can be tedious and prone to leaks.

**Cyph3r** includes a native Ghost Mode
(`--ghost`) layer that automates this process directly inside the application core.

# How Cyph3r Ghost Mode Works

When you append `--ghost` to any Cyph3r command:

1. Tor Proxy Binding: It automatically routes outbound HTTP requests, reconnaissance scans, and port sweeps through `127.0.0.1:9050`

2. Safety Guardrails: It actively detects and blocks operations incompatible with proxy tunneling (such as raw socket Layer 4 SYN floods) to prevent accidental IP leaks.

# Examples in Practice

Run an anonymous `OSINT` footprint extraction through Tor:

```
./cyph3r --target example.com --ghost --osint
```

Execute a stealth TCP port sweep across common service vectors anonymously:

```
./cyph3r --target 192.168.1.50 --ghost --scan
```

Run a high-performance benchmark via Tor SOCKS5 tunnels:

```
./cyph3r --target https://target.com --ghost --wrk -c 50 -d 30
```
