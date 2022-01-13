package main

// Integrate with Prometheus

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			delta := rand.Int63n(10) - 5
			opsQueued.Add(float64(delta))
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	opsQueued = promauto.NewGauge(prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "ops_queued",
		Help: "Number of blob storage operations waiting to be processed.",
	})
)

func processHistogram() {
	temps := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "x_pond_temperature_celsius",
		Help:    "x The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5),    // 5 buckets, each 5 centigrade wide.
	})

	// Simulate some observations.
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}

	// Just for demonstration, let's check the state of the histogram by
	// (ab)using its Write method (which is usually only used by Prometheus
	// internally).
	//metric := &dto.Metric{}
	//temps.Write(metric)
	//fmt.Println(proto.MarshalTextString(metric))
}

func quantile() {
	temps := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "y_my_func",
		Help:    "My Function",
		Buckets: prometheus.LinearBuckets(2, 2, 5),
	})

	for i := 1; i <= 10; i++ {
		temps.Observe(float64(i))
	}
}

func summary() {
	temps := promauto.NewSummary(prometheus.SummaryOpts{
		Name:       "a_my_sliding",
		Help:       "a_The temperature of the frog pond.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	// Simulate some observations.
	//for i := 0; i < 1000; i++ {
	//	temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	//}

	temps.Observe(1)
	temps.Observe(3)
	temps.Observe(8)
	temps.Observe(9)
	temps.Observe(10)
}

func main() {
	recordMetrics()

	go processHistogram()
	go quantile()
	go summary()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
