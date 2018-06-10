package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	peersLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "peers_discovery_latency",
		Help:    "Latency to discover peers",
		Buckets: []float64{0.5, 2, 5, 10, 20, 30},
	}, []string{"peers"})
	started = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "peers_discovery_started",
		Help: "Number of healtcheck started",
	})
	failed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "peers_discovery_failed",
		Help: "Number of healtcheck failed",
	})
)

func init() {
	prometheus.MustRegister(peersLatency, started, failed)
}

// Stats represents messages' statistics.
type Stats struct {
}

// For verifying the service is up
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// NewStats returns new empty Stats object.
func NewStats(statsPort string) Stats {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// Add most trivial healthcheck
		http.HandleFunc("/health", healthHandler)
		log.Fatal(http.ListenAndServe(statsPort, nil))
	}()
	return Stats{}
}

func (s Stats) Discovered(count int, d time.Duration) {
	peersLatency.WithLabelValues(strconv.Itoa(count)).Observe(float64(d) / float64(time.Second))
}

func (s Stats) Started() {
	started.Add(1)
}

func (s Stats) Failed() {
	failed.Add(1)
}
