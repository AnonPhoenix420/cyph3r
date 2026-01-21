package main

// MAIN MODULE
// Unified Load Tester + Mixed Scenario Scheduler

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"cyph3r/output"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ================= CONFIG =================

type Config struct {
	Target   string
	Port     int
	Proto    string
	Method   string
	Payload  string
	Headers  headerList
	Whois    bool
	Duration time.Duration
	RPS      int
	Workers  int
	Ramp     time.Duration
	JSON     bool
	Monitor  bool
	Interval time.Duration

	FailRateThreshold float64
	LatencyThreshold  time.Duration
	ASNFanoutLimit    int
}

// ================= METRICS =================

type Metrics struct {
	Sent     uint64
	Success  uint64
	Failure  uint64
	LatSumUs uint64

	latMu sync.Mutex
	lats  []uint64
}

func (m *Metrics) record(lat time.Duration, ok bool) {
	atomic.AddUint64(&m.Sent, 1)
	if ok {
		atomic.AddUint64(&m.Success, 1)
	} else {
		atomic.AddUint64(&m.Failure, 1)
	}
	atomic.AddUint64(&m.LatSumUs, uint64(lat.Microseconds()))

	m.latMu.Lock()
	m.lats = append(m.lats, uint64(lat.Microseconds()))
	m.latMu.Unlock()
}

func (m *Metrics) percentiles() (p50, p95, p99 time.Duration) {
	m.latMu.Lock()
	defer m.latMu.Unlock()
	if len(m.lats) == 0 {
		return
	}
	sort.Slice(m.lats, func(i, j int) bool { return m.lats[i] < m.lats[j] })
	idx := func(q float64) int { return int(float64(len(m.lats)-1) * q) }
	p50 = time.Duration(m.lats[idx(0.50)]) * time.Microsecond
	p95 = time.Duration(m.lats[idx(0.95)]) * time.Microsecond
	p99 = time.Duration(m.lats[idx(0.99)]) * time.Microsecond
	return
}

// ================= PROMETHEUS =================

var (
	pSent    = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "tester_requests_total"}, []string{"proto", "scenario", "target"})
	pSuccess = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "tester_success_total"}, []string{"proto", "scenario", "target"})
	pFail    = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "tester_failure_total"}, []string{"proto", "scenario", "target"})
	pLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "tester_latency_ms", Buckets: prometheus.DefBuckets}, []string{"proto", "scenario", "target"})
)

func initProm() {
	prometheus.MustRegister(pSent, pSuccess, pFail, pLatency)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
			status := map[string]interface{}{"timestamp": time.Now().Format(time.RFC3339)}
			json.NewEncoder(w).Encode(status)
		})
		_ = http.ListenAndServe(":2112", nil)
	}()
}

// ================= WORKER =================

func worker(ctx context.Context, jobs <-chan string, cfg Config, m *Metrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case scenario, ok := <-jobs:
			if !ok { return }
			start := time.Now()
			ok := false

			switch scenario {
			case "http", "https":
				// call http worker
				ok = httpWorker(cfg, scenario)
			case "tcp":
				ok = tcpProbe(cfg.Target, cfg.Port)
			case "icmp":
				_, ok = icmpPing(cfg.Target, 2*time.Second)
			}

			lat := time.Since(start)
			m.record(lat, ok)
			pSent.WithLabelValues(scenario, scenario, cfg.Target).Inc()
			pLatency.WithLabelValues(scenario, scenario, cfg.Target).Observe(float64(lat.Milliseconds()))
			if ok {
				pSuccess.WithLabelValues(scenario, scenario, cfg.Target).Inc()
			} else {
				pFail.WithLabelValues(scenario, scenario, cfg.Target).Inc()
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
			// output metrics
			sent := atomic.LoadUint64(&m.Sent)
			ok := atomic.LoadUint64(&m.Success)
			fail := atomic.LoadUint64(&m.Failure)
			avg := time.Duration(0)
			if sent > 0 {
				avg = time.Duration(atomic.LoadUint64(&m.LatSumUs)/sent) * time.Microsecond
			}
			p50, p95, p99 := m.percentiles()
			output.Info(fmt.Sprintf("sent=%d ok=%d fail=%d avg=%s p95=%s", sent, ok, fail, avg, p95))
		}
	}
}

// ================= MAIN =================

func main() {
	output.Banner()

	cfg := Config{}
	flag.StringVar(&cfg.Target, "target", "localhost", "target host")
	flag.IntVar(&cfg.Port, "port", 80, "port")
	flag.StringVar(&cfg.Proto, "proto", "http", "tcp|http|https|icmp")
	flag.IntVar(&cfg.ASNFanoutLimit, "asnlimit", 50, "ASN fan-out limit")
	flag.DurationVar(&cfg.Duration, "duration", 30*time.Second, "test duration")
	flag.IntVar(&cfg.RPS, "rps", 200, "requests per second")
	flag.IntVar(&cfg.Workers, "workers", runtime.NumCPU()*4, "workers")
	flag.DurationVar(&cfg.Ramp, "ramp", 10*time.Second, "ramp up duration")
	flag.BoolVar(&cfg.JSON, "json", false, "json output")
	flag.BoolVar(&cfg.Monitor, "monitor", true, "live dashboard")
	flag.DurationVar(&cfg.Interval, "interval", 2*time.Second, "dashboard interval")
	flag.Float64Var(&cfg.FailRateThreshold, "failrate", 0.1, "failure threshold")
	flag.DurationVar(&cfg.LatencyThreshold, "latency", 2*time.Second, "p95 latency threshold")
	flag.Parse()

	initProm()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, cfg.Duration)
	defer cancel()

	metrics := &Metrics{lats: make([]uint64, 0, cfg.RPS*int(cfg.Duration.Seconds()))}
	jobs := make(chan string, cfg.Workers)

	var wg sync.WaitGroup
	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, cfg, metrics, &wg)
	}

	if cfg.Monitor {
		go dashboard(ctx, metrics, cfg)
	}

	// Mixed scenario scheduler
	scenarios := []string{"http", "tcp", "icmp"}
	tick := time.NewTicker(time.Second / time.Duration(max(1, cfg.RPS)))
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return
		case <-tick.C:
			// Round-robin mixed scenario
			s := scenarios[time.Now().Second()%len(scenarios)]
			jobs <- s
		}
	}
}

type headerList map[string]string

func (h *headerList) String() string { return "" }
func (h *headerList) Set(v string) error {
	if *h == nil { *h = make(map[string]string) }
	parts := strings.SplitN(v, ":", 2)
	if len(parts) == 2 { (*h)[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1]) }
	return nil
}

func max(a, b int) int { if a > b { return a } ; return b }
