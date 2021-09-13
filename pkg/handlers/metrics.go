package handlers

import (
	"net/http"

	"github.com/Serj1c/yaml-to-openmetrics/pkg/parsers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metric represents one of possible app's metrics
type Metric struct {
	Path string
}

// NewMetric creates an exemplar of a metric
func NewMetric(Path string) *Metric {
	return &Metric{Path}
}

// Prepare "prepares" metrics which are to be sent
func (m *Metric) Prepare(rw http.ResponseWriter, r *http.Request) {
	yml, err := parsers.ParseYaml(m.Path)
	if err != nil {
		http.Error(rw, "unable to parse yaml file", http.StatusBadRequest)
	}
	reg := prometheus.NewRegistry()
	for _, v := range yml.Currencies {
		gauge := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: v.Name,
				Help: "This is the current value of " + v.Name + " against rub",
			})
		reg.MustRegister(gauge)
		gauge.Add(v.Value)
	}
	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	h.ServeHTTP(rw, r)
}
