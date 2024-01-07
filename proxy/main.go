package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Request total",
		},
		[]string{"endpoint"},
	)

	cacheDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cache_duration_seconds",
			Help:    "Cache duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	dbDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_duration_seconds",
			Help:    "DB duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	externalAPIDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "external_api_duration_seconds",
			Help:    "API duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func main() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(cacheDuration)
	prometheus.MustRegister(dbDuration)
	prometheus.MustRegister(externalAPIDuration)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	http.ListenAndServe(":8080", nil)

}
