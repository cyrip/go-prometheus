package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of HTTP requests processed, labeled by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		httpRequestsTotal.WithLabelValues("200", r.Method).Inc()
	})

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		if r.URL.Query().Get("error") == "true" {
			http.Error(w, "Simulated error", http.StatusInternalServerError)
			httpRequestsTotal.WithLabelValues("500", r.Method).Inc()
			return
		}
		w.Write([]byte("Work done"))
		httpRequestsTotal.WithLabelValues("200", r.Method).Inc()
	})

	log.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
