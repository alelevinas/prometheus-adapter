package monitoring

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/klog/v2"
)

// ServePrometheusMetrics async starts the server to report metrics to
func ServePrometheusMetrics(port int) {
	mux := http.NewServeMux()
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	mux.Handle("/metrics", promhttp.Handler())
	klog.Infof("Serving Prometheus metrics on 127.0.0.1:%d/metrics", port)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
