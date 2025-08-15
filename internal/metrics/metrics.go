package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	ActiveWorkers prometheus.Gauge
}

func NewMetric(reg prometheus.Registerer) *Metric {
	m := &Metric{
		ActiveWorkers: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "workerpool",
			Name: "active_worker_count",
			Help: "number of active workers of cimri workerpool microservice",
		}),
	}
	reg.MustRegister(m.ActiveWorkers)
	return m
}

func (m Metric) IncrementActiveWorkerCount() {
	m.ActiveWorkers.Add(1)
}

func (m Metric) DecrementActiveWorkerCount() {
	m.ActiveWorkers.Add(-1)
}