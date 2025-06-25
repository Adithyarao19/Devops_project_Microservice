package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus counter
var requestCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "service1_http_requests_total",
	Help: "Total HTTP requests handled by Service 1",
})

func init() {
	// Register the Prometheus metric
	prometheus.MustRegister(requestCount)
}

func main() {
	// Instrument existing handlers without changing logic
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		requestCount.Inc()
		jsonResponse(w, map[string]string{
			"status":  "ok",
			"service": "1",
		})
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		requestCount.Inc()
		jsonResponse(w, map[string]string{
			"message": "Hello from Service 1",
		})
	})

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Service 1 listening on port 8001...")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
