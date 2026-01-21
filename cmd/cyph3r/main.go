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
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/output"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/* ================= CONFIG ================= */

type Config struct {
	Target            string
	Port              int
	Proto             string
	Method            string
	Payload           string
	Headers           headerList
	Duration          time.Duration
	RPS               int
	Workers           int
	JSON              bool
	Monitor           bool
	Interval          time.Duration
	FailRateThreshold float64
	LatencyThreshold  time.Duration
}

/* ================= METRICS ================= */

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

func (m *Metrics) percentiles() (time.Duration, time.Duration, time.Duration) {
	m.latMu.Lock()
	defer m.latMu.Unlock()

	if len(m.lats) == 0 {
		return 0, 0, 0
	}

	sort.Slice(m.lats, func(i, j int) bool {
		return m.lats[i] < m.lats[j]
	})

	idx := func(q float64) int {
		return int(float64(len(m.lats)-1) * q)
	}

	return time.Duration(m.lats[idx(0.50)]) * time.Microsecond,
		time.Duration(m.lats[idx(0.95)]) * time.Microsecond,
		time.Duration(m.lats[idx(0.99)]) * time.Microsecond
}

/* ================= PROMETHEUS ================= */

var (
	pSent = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cyph3r_requests_total",
	})
	pSuccess = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cyph3r_success_total",
	})
	pFail = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cyph3r_failure_total",
	})
	pLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "cyph3r_latency_ms",
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

/* ================= HTTP CLIENT ================= */

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        2500,
		MaxIdleConnsPerHost: 1100,
		IdleConnTimeout:     90 * time.Second,
	},
}

/* ================= WORKER ================= */

func worker(
	ctx context.Context,
	jobs <-chan struct{},
	cfg Config,
	m *Metrics,
	wg *sync.WaitGroup,
) {
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
				conn, err := net.DialTimeout(
					"tcp",
					fmt.Sprintf("%s:%d", cfg.Target, cfg.Port),
					2*time.Second,
				)
				if err == nil {
					conn.Close()
					ok = true
				}

			case "http", "https":
				url := cfg.Proto + "://" + cfg.Target
				var body io.Reader
				if cfg.Payload != "" {
					body = strings.NewReader(cfg.Payload)
				}

				req, err := http.NewRequest(cfg.Method, url, body)
				if err == nil {
					req.Header.Set("User-Agent", "cyph3r")
					for k, v := range cfg.Headers {
						req.Header.Set(k, v)
					}
					resp, err := httpClient.Do(req)
					if err == nil {
						ok = resp.StatusCode >= 200 && resp.StatusCode < 400
						resp.Body.Close()
					}
				}
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

/* ================= DASHBOARD ================= */

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

			_, p95, _ := m.percentiles()

			output.Info(
				fmt.Sprintf(
					"sent=%d ok=%d fail=%d avg=%s p95=%s",
					sent, ok, fail, avg, p95,
				),
			)

			if cfg.FailRateThreshold > 0 && sent > 0 {
				failRate := float64(fail) / float64(sent)
				if failRate > cfg.FailRateThreshold {
					output.Down("FAILURE THRESHOLD BREACHED")
				}
			}

			if cfg.LatencyThreshold > 0 && p95 > cfg.LatencyThreshold {
				output.Down("LATENCY THRESHOLD BREACHED")
			}
		}
	}
}

/* ================= MAIN ================= */

func main() {
	output.Banner()

	cfg := Config{}

	flag.StringVar(&cfg.Target, "target", "localhost", "target host")
	flag.IntVar(&cfg.Port, "port", 80, "port")
	flag.StringVar(&cfg.Proto, "proto", "http", "tcp|http|https")
	flag.StringVar(&cfg.Method, "method", "GET", "HTTP method")
	flag.StringVar(&cfg.Payload, "payload", "", "HTTP body")
	flag.Var(&cfg.Headers, "H", "HTTP header key:value")
	flag.DurationVar(&cfg.Duration, "duration", 30*time.Second, "test duration")
	flag.IntVar(&cfg.RPS, "rps", 100, "requests per second")
	flag.IntVar(&cfg.Workers, "workers", runtime.NumCPU()*4, "workers")
	flag.BoolVar(&cfg.JSON, "json", false, "json output")
	flag.BoolVar(&cfg.Monitor, "monitor", true, "dashboard")
	flag.DurationVar(&cfg.Interval, "interval", 2*time.Second, "dashboard interval")
	flag.Float64Var(&cfg.FailRateThreshold, "failrate", 0.0, "fail threshold")
	flag.DurationVar(&cfg.LatencyThreshold, "latency", 0, "latency threshold")
	flag.Parse()

	initProm()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()

var cancel context.CancelFunc
if !cfg.Monitor {
    // Only use duration timeout if NOT in monitor mode
    ctx, cancel = context.WithTimeout(ctx, cfg.Duration)
    defer cancel()
}


	metrics := &Metrics{lats: make([]uint64, 0, cfg.RPS)}
	jobs := make(chan struct{}, cfg.Workers)

	var wg sync.WaitGroup
	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, cfg, metrics, &wg)
	}

	if cfg.Monitor {
		go dashboard(ctx, metrics, cfg)
	}

	tick := time.NewTicker(time.Second / time.Duration(maxInt(1, cfg.RPS)))
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()

			p50, p95, p99 := metrics.percentiles()
			result := map[string]interface{}{
				"sent":    metrics.Sent,
				"success": metrics.Success,
				"failure": metrics.Failure,
				"p50":     p50.String(),
				"p95":     p95.String(),
				"p99":     p99.String(),
			}

			if cfg.JSON {
				b, _ := json.MarshalIndent(result, "", "  ")
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

/* ================= HELPERS ================= */

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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
