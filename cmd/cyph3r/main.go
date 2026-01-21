package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	
    "github.com/AnonPhoenix420/cyph3r/internal/intel"
    "github.com/AnonPhoenix420/cyph3r/internal/output"


	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"sort"
)

// ================= CONFIG =================

type Config struct {
	Target            string
	Port              int
	Proto             string
	Method            string
	Payload           string
	Headers           headerList
	Whois             bool
	Duration          time.Duration
	RPS               int
	Workers           int
	Ramp              time.Duration
	JSON              bool
	Monitor           bool
	Interval          time.Duration
	FailRateThreshold float64
	LatencyThreshold  time.Duration
	Scenario          string
	ASNFanout         bool
}
// ---------- FIXED METRICS / STATS BLOCK ----------

start := time.Now()

sent, received := runTrafficTest(target, port)

// uint64-safe max function
func maxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

total := maxUint64(sent, received)

// Latency stats (use values so Go is happy)
p50, p99 := calculateLatencyStats()

fmt.Printf("Latency p50: %v ms\n", p50)
fmt.Printf("Latency p99: %v ms\n", p99)

elapsed := time.Since(start)
fmt.Printf("Elapsed time: %s\n", elapsed)

fmt.Printf("Packets sent: %d\n", sent)
fmt.Printf("Packets received: %d\n", received)
fmt.Printf("Total packets: %d\n", total)

// ================= PROMETHEUS =================

var (
	pSent    = prometheus.NewCounter(prometheus.CounterOpts{Name: "tester_requests_total"})
	pSuccess = prometheus.NewCounter(prometheus.CounterOpts{Name: "tester_success_total"})
	pFail    = prometheus.NewCounter(prometheus.CounterOpts{Name: "tester_failure_total"})
	pLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tester_latency_ms",
		Buckets: prometheus.DefBuckets,
	})
)

func initProm() {
	prometheus.MustRegister(pSent, pSuccess, pFail, pLatency)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()
}

// ================= HTTP CLIENT =================

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	},
}

// ================= WORKER =================

func worker(ctx context.Context, jobs <-chan struct{}, cfg Config, m *Metrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-jobs:
			start := time.Now()
			ok := false

			switch cfg.Proto {
			case "tcp":
				c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", cfg.Target, cfg.Port), 2*time.Second)
				if err == nil {
					c.Close()
					ok = true
				}
			case "http", "https":
				scheme := cfg.Proto
				url := scheme + "://" + cfg.Target
				var body io.Reader
				if cfg.Payload != "" {
					body = strings.NewReader(cfg.Payload)
				}
				req, err := http.NewRequest(cfg.Method, url, body)
				if err == nil {
					req.Header.Set("User-Agent", "cyph3r/2.1")
					for k, v := range cfg.Headers {
						req.Header.Set(k, v)
					}
					resp, err := httpClient.Do(req)
					if err == nil {
						ok = resp.StatusCode >= 200 && resp.StatusCode < 400
						resp.Body.Close()
					}
				}
			case "icmp":
				ok = intel.CheckICMP(cfg.Target)
			}

			lat := time.Since(start)
			m.record(lat, ok)
			pSent.Inc()
			pLatency.Observe(float64(lat.Milliseconds()))
			if ok {
				pSuccess.Inc()
			} else {
				pFail.Inc()
			}
		}
	}
}

// ================= DASHBOARD =================

func dashboard(ctx context.Context, m *Metrics, cfg Config) {
	t := time.NewTicker(cfg.Interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			sent := atomic.LoadUint64(&m.Sent)
			ok := atomic.LoadUint64(&m.Success)
			fail := atomic.LoadUint64(&m.Failure)
			avg := time.Duration(0)
			if sent > 0 {
				avg = time.Duration(atomic.LoadUint64(&m.LatSumUs)/sent) * time.Microsecond
			}
			p50, p95, p99 := m.percentiles()
			output.Info(fmt.Sprintf("sent=%d ok=%d fail=%d avg=%s p95=%s", sent, ok, fail, avg, p95))
			failRate := float64(fail) / float64(max(1, sent))
			if cfg.FailRateThreshold > 0 && failRate > cfg.FailRateThreshold {
				output.Down("FAILURE THRESHOLD BREACHED")
			}
			if cfg.LatencyThreshold > 0 && p95 > cfg.LatencyThreshold {
				output.Down("LATENCY THRESHOLD BREACHED")
			}
		}
	}
}

// ================= MAIN =================

func main() {
	output.Banner()

	cfg := Config{}
	flag.StringVar(&cfg.Target, "target", "localhost", "Target host")
	flag.IntVar(&cfg.Port, "port", 80, "Port")
	flag.StringVar(&cfg.Proto, "proto", "http", "tcp|http|https|icmp")
	flag.StringVar(&cfg.Method, "method", "GET", "HTTP method")
	flag.StringVar(&cfg.Payload, "payload", "", "HTTP payload")
	flag.Var(&cfg.Headers, "H", "HTTP header key:value")
	flag.BoolVar(&cfg.Whois, "whois", false, "WHOIS lookup")
	flag.DurationVar(&cfg.Duration, "duration", 30*time.Second, "Test duration")
	flag.IntVar(&cfg.RPS, "rps", 200, "Requests per second")
	flag.IntVar(&cfg.Workers, "workers", runtime.NumCPU()*4, "Workers")
	flag.DurationVar(&cfg.Ramp, "ramp", 10*time.Second, "Ramp up duration")
	flag.BoolVar(&cfg.JSON, "json", false, "JSON output")
	flag.BoolVar(&cfg.Monitor, "monitor", true, "Live dashboard")
	flag.DurationVar(&cfg.Interval, "interval", 2*time.Second, "Dashboard interval")
	flag.Float64Var(&cfg.FailRateThreshold, "failrate", 0.1, "Failure threshold")
	flag.DurationVar(&cfg.LatencyThreshold, "latency", 2*time.Second, "p95 latency threshold")
	flag.StringVar(&cfg.Scenario, "scenario", "", "Mixed scenario")
	flag.BoolVar(&cfg.ASNFanout, "asn-fanout", false, "ASN fan-out mode")
	flag.Parse()

	initProm()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, cfg.Duration)
	defer cancel()

	metrics := &Metrics{lats: make([]uint64, 0, cfg.RPS*int(cfg.Duration.Seconds()))}
	jobs := make(chan struct{}, cfg.Workers)

	var wg sync.WaitGroup
	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, cfg, metrics, &wg)
	}

	if cfg.Monitor {
		go dashboard(ctx, metrics, cfg)
	}

	start := time.Now()
	tick := time.NewTicker(time.Second / time.Duration(max(1, cfg.RPS)))
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			p50, p95, p99 := metrics.percentiles()
			out := map[string]interface{}{
				"sent":    metrics.Sent,
				"success": metrics.Success,
				"failure": metrics.Failure,
				"p50":     p50.String(),
				"p95":     p95.String(),
				"p99":     p99.String(),
			}
			if cfg.JSON {
				b, _ := json.MarshalIndent(out, "", "  ")
				fmt.Println(string(b))
			} else {
				output.Success("Completed")
			}
			return
		case <-tick.C:
			jobs <- struct{}{}
		}
	}
}

// ================= HELPERS =================

type headerList map[string]string

func (h *headerList) String() string { return "" }
func (h *headerList) Set(v string) error {
	if *h == nil {
		*h = make(map[string]string)
	}
	parts := strings.SplitN(v, ":", 2)
	if len(parts) == 2 {
		(*h)[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return nil
}

func max(a, b int) int { if a > b { return a }; return b }
