//go:generate mockgen -source=registry.go -package=metrics -destination=registry.mock.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Registry interface {
	prometheus.Registerer

	NewCounter(opts prometheus.CounterOpts) prometheus.Counter
	NewCounterVec(opts prometheus.CounterOpts, labels []string) *prometheus.CounterVec
	NewCounterFunc(opts prometheus.CounterOpts, function func() float64) prometheus.CounterFunc
	NewGauge(opts prometheus.GaugeOpts) prometheus.Gauge
	NewGaugeVec(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec
	NewGaugeFunc(opts prometheus.GaugeOpts, function func() float64) prometheus.GaugeFunc
	NewSummary(opts prometheus.SummaryOpts) prometheus.Summary
	NewSummaryVec(opts prometheus.SummaryOpts, labelNames []string) *prometheus.SummaryVec
	NewHistogram(opts prometheus.HistogramOpts) prometheus.Histogram
	NewHistogramVec(opts prometheus.HistogramOpts, labelNames []string) *prometheus.HistogramVec
	NewUntypedFunc(opts prometheus.UntypedOpts, function func() float64) prometheus.UntypedFunc
}

type PromRegistry struct {
	namespace string
	prom      prometheus.Registerer
}

func NewDefaultRegistry(cfg Config) *PromRegistry {
	return NewRegistry(cfg, prometheus.DefaultRegisterer)
}

func NewRegistry(cfg Config, registerer prometheus.Registerer) *PromRegistry {
	return &PromRegistry{
		namespace: cfg.Namespace,
		prom:      registerer,
	}
}

func (r *PromRegistry) NewCounter(opts prometheus.CounterOpts) prometheus.Counter {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewCounter(opts)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewCounterVec(opts prometheus.CounterOpts, labels []string) *prometheus.CounterVec {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewCounterVec(opts, labels)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewCounterFunc(opts prometheus.CounterOpts, function func() float64) prometheus.CounterFunc {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewCounterFunc(opts, function)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewGauge(opts prometheus.GaugeOpts) prometheus.Gauge {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewGauge(opts)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewGaugeVec(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewGaugeVec(opts, labelNames)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewGaugeFunc(opts prometheus.GaugeOpts, function func() float64) prometheus.GaugeFunc {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewGaugeFunc(opts, function)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewSummary(opts prometheus.SummaryOpts) prometheus.Summary {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewSummary(opts)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewSummaryVec(opts prometheus.SummaryOpts, labelNames []string) *prometheus.SummaryVec {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewSummaryVec(opts, labelNames)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewHistogram(opts prometheus.HistogramOpts) prometheus.Histogram {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewHistogram(opts)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewHistogramVec(opts prometheus.HistogramOpts, labelNames []string) *prometheus.HistogramVec {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewHistogramVec(opts, labelNames)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) NewUntypedFunc(opts prometheus.UntypedOpts, function func() float64) prometheus.UntypedFunc {
	if opts.Namespace == "" {
		opts.Namespace = r.namespace
	}

	collector := prometheus.NewUntypedFunc(opts, function)

	r.MustRegister(collector)

	return collector
}

func (r *PromRegistry) Register(collector prometheus.Collector) error {
	return r.prom.Register(collector)
}

func (r *PromRegistry) MustRegister(collector ...prometheus.Collector) {
	r.prom.MustRegister(collector...)
}

func (r *PromRegistry) Unregister(collector prometheus.Collector) bool {
	return r.prom.Unregister(collector)
}
