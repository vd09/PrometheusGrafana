package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Testing validation of stackoverflow theory
// https://stackoverflow.com/questions/55162093/understanding-histogram-quantile-based-on-rate-in-prometheus
func TestingHistogramQuantile(reg prometheus.Registerer) {
	histogram := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "histogram_metric",
		Buckets: []float64{10, 50, 100, 200, 300, 500, 1000},
	})
	reg.MustRegister(histogram)

	bucket := []int{0, 10, 50, 100, 200, 300, 500, 1000, 2000}
	count := []int{3000, 3000, 1500, 1000, 800, 400, 200, 100}
	go func() {
		for {
			for j := 0; j < len(count); j++ {
				diff := bucket[j+1] - bucket[j]
				for i := 0; i < count[j]; i++ {
					histogram.Observe(float64(rand.Int()%diff) + float64(bucket[j]) + 1)
				}
			}
			time.Sleep(time.Minute + (3 * time.Second))
		}
	}()
}
