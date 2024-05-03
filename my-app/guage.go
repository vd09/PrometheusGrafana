package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func TestGaugeMetrics(reg *prometheus.Registry) {
	var activeConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gauge_metric",
		},
		[]string{"vd"},
	)
	reg.MustRegister(activeConnections)
	go func() {
		for {
			metricTags := prometheus.Labels{}
			metricTags["vd"] = "working"
			if rand.Int()%2 == 0 {
				metricTags["vd"] = "chilling"
			}
			//activeConnections.With(prometheus.Labels(metricTags)).Add(float64(8))
			activeConnections.With(prometheus.Labels(metricTags)).Set(float64(rand.Int() % 2))
			time.Sleep(50 * time.Millisecond)
		}
	}()
}
