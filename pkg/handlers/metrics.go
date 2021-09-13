package handlers

import (
	"net/http"

	"github.com/Serj1c/yaml-to-openmetrics/pkg/parsers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics ...
type Metrics struct {
	Path string
}

// NewMetrics creates an exemplar of metrics
func NewMetrics(Path string) *Metrics {
	return &Metrics{Path}
}

// Prepare "prepares" metrics which are to be sent
func (m *Metrics) Prepare(rw http.ResponseWriter, r *http.Request) {
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
