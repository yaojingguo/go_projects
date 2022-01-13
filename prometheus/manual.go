package main

// Integrate with Prometheus

import (
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func histogram() prometheus.Histogram {
	temps := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "x_pond_temperature_celsius",
		Help:    "x The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(10, 5, 5),    // 5 buckets, each 5 centigrade wide.
	})

	go func() {
		for i := 0; i < 30; i++ {
			time.Sleep(1 * time.Second)
			temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
		}
		for {
			time.Sleep(time.Second * 1)
			temps.Observe(40)
		}
	}()

	return temps
}

func histogram2_1() prometheus.Histogram {
	temps := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "my_hist_3",
		Help:    "my little historgram",             // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(1, 1, 10), // 5 buckets, each 5 centigrade wide.
	})

	go func() {
		observe(&temps)
	}()

	return temps
}

func histogram2_2() prometheus.Histogram {
	temps := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "my_hist_2",
		Help:    "my little historgram",            // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(5, 2, 5), // 5 buckets, each 5 centigrade wide.
	})

	go func() {
		observe(&temps)
	}()

	return temps
}

func observe(hist *prometheus.Histogram) {
	(*hist).Observe(1)
	(*hist).Observe(1)
	(*hist).Observe(1)
	(*hist).Observe(3)
	(*hist).Observe(3)

	(*hist).Observe(4)
	(*hist).Observe(4)
	(*hist).Observe(4)
	(*hist).Observe(7)
	(*hist).Observe(8)
}

func counter() prometheus.Counter {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_counter",
		Help: "my example counter",
	})

	go func() {
		for {
			time.Sleep(time.Second * 1)
			counter.Inc()
		}
	}()

	return counter
}

func gauge() prometheus.Gauge {
	g := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "my_gauge",
		Help: "my example gauge",
	})

	go func() {
		for {
			time.Sleep(time.Second * 1)
			g.Add(10.0)
		}
	}()

	return g
}

func counterVec() *prometheus.CounterVec {
	httpReqs := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "how many requests",
	}, []string{"code", "method"},
	)

	go func() {
		httpReqs.WithLabelValues("200", "GET").Inc()
		httpReqs.WithLabelValues("300", "GET").Inc()
		httpReqs.WithLabelValues("400", "GET").Inc()

		httpReqs.WithLabelValues("400", "POST").Inc()
		httpReqs.WithLabelValues("400", "POST").Inc()
		httpReqs.WithLabelValues("400", "POST").Inc()
	}()

	return httpReqs
}

func main() {
	r := prometheus.NewRegistry()

	r.MustRegister(histogram2_1())
	r.MustRegister(histogram2_2())
	r.MustRegister(counter())
	r.MustRegister(gauge())
	r.MustRegister(counterVec())

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)
	http.ListenAndServe(":2112", nil)
}
