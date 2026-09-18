package internal

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

// ScopedPorts defines high-yield service, administrative, and SQL ports for Ghost scans.
var ScopedPorts = []int{
	80, 443, 8080, 8443,       // Web / App
	22, 3389,                  // Administrative (SSH, RDP)
	1433, 3306, 5432, 27017,   // Databases / SQL
	53,                        // Core Infrastructure
}

// NewGhostTransport initializes a resilient proxied HTTP transport routing through Tor/SOCKS5
// with hardened timeout controls, context enforcement, and request jitter to prevent deadlocks.
func NewGhostTransport(proxyAddr string) (*http.Transport, error) {
	if proxyAddr == "" {
		proxyAddr = "127.0.0.1:9050" // Default Tor SOCKS5 listener
	}

	baseDialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	// Ensure the proxy dialer supports context-aware timeouts
	contextDialer, ok := baseDialer.(proxy.ContextDialer)
	if !ok {
		contextDialer = &proxyContextDialerAdapter{dialer: baseDialer}
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Introduce subtle request jitter (0 to 150ms) to bypass timing signatures
			jitter := time.Duration(rand.Intn(150)) * time.Millisecond
			time.Sleep(jitter)

			// Enforce a strict per-connection timeout context for slow Tor circuits
			dialCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
			defer cancel()

			return contextDialer.DialContext(dialCtx, network, addr)
		},
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   20 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
		ForceAttemptHTTP2:     true,
	}

	return transport, nil
}

// proxyContextDialerAdapter wraps standard proxy dialers with context support to prevent infinite hangs.
type proxyContextDialerAdapter struct {
	dialer proxy.Dialer
}

func (a *proxyContextDialerAdapter) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	type dialResult struct {
		conn net.Conn
		err  error
	}
	ch := make(chan dialResult, 1)

	go func() {
		c, e := a.dialer.Dial(network, addr)
		ch <- dialResult{conn: c, err: e}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-ch:
		return res.conn, res.err
	}
}
