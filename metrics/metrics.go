package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	DbRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "Db_request_duration_seconds",
		Help: "The duration of the requests to the DB service.",
	})

	DbRequestsCurrent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Db_requests_current",
		Help: "The current number of requests to the DB service.",
	})

	DbClientErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Db_errors",
		Help: "The total number of DB client errors",
	})
)

func init() {
	prometheus.MustRegister(DbRequestsDuration)
	prometheus.MustRegister(DbClientErrors)
	prometheus.MustRegister(DbRequestsCurrent)
}

// MetricsHandler defined for collecting application metrics
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
