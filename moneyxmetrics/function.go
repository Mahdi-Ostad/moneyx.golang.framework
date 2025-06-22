package MoneyxMetrics

import (
	"context"
	"time"
)

// Accepts a MultiOutputFunc and prints the outputs
func FuncWithMetrics(fn func()) {
	defer func(start time.Time) {
		duration := time.Since(start)
		metrics := getGlobalMetrics()
		metrics.RecordHistogram(context.Background(), "func", duration.Seconds())
	}(time.Now())
	fn()
}
