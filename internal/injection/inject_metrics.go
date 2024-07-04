package injection

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/bnb-chain/token-recover-app/internal/metrics"
	collector "github.com/bnb-chain/token-recover-app/internal/metrics/prometheus"
)

func InitPrometheusRegister() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func InitMetrics(registry *prometheus.Registry) metrics.Metrics {
	return collector.NewCollector(registry)
}
