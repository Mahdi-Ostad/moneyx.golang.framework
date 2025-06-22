package MoneyxMetrics

import "gofr.dev/pkg/gofr/metrics"

var globalMetrics metrics.Manager

func SetGlobalMetrics(metrics metrics.Manager) {
	globalMetrics = metrics
}

func getGlobalMetrics() metrics.Manager {
	return globalMetrics
}
