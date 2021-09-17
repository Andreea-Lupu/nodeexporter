package api

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/anuvu/zot/pkg/extensions/monitoring"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	up = prometheus.NewDesc(
		"zot_up",
		"Connection to zot server was successfully established.",
		nil, nil,
	)
	invalidChars = regexp.MustCompile("[^a-zA-Z0-9:_]")
	controller *Controller
)

type ZotCollector struct {
}

// Implements prometheus.Collector.
func (c ZotCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
}

// Implements prometheus.Collector.
func (c ZotCollector) Collect(ch chan<- prometheus.Metric) {
	config := monitoring.ZotMetricsConfig{}
	config.Address = controller.Address

	zot, err := monitoring.NewMetricsClient(&config, controller.Config.HTTP.Host, controller.Log)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}

	metrics, err := zot.GetMetrics()

	if err != nil {
		fmt.Println(err)
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)

	fmt.Println("Metrics: %#v", metrics)

	for _, g := range metrics.Gauges {
		name := invalidChars.ReplaceAllLiteralString(g.Name, "_")
		desc := prometheus.NewDesc(name, "Zot metric "+g.Name, g.LabelNames, nil)
		ch <- prometheus.MustNewConstMetric(
			desc, prometheus.GaugeValue, float64(g.Value), g.LabelValues...)
	}

	for _, c := range metrics.Counters {
		fmt.Println("Counter: %#v", c)
		name := invalidChars.ReplaceAllLiteralString(c.Name, "_")
		fmt.Println("Counter name : %s", name)
		desc := prometheus.NewDesc(name+"_total", "Zot metric "+c.Name, c.LabelNames, nil)
		fmt.Println("Counter desc : %#v", desc)
		ch <- prometheus.MustNewConstMetric(
			desc, prometheus.CounterValue, float64(c.Count), c.LabelValues...)
	}

	for _, s := range metrics.Samples {
		// All samples are times in milliseconds, we convert them to seconds below.
		name := invalidChars.ReplaceAllLiteralString(s.Name, "_")
		countDesc := prometheus.NewDesc(
			name+"_count", "Zot metric "+s.Name, s.LabelNames, nil)
		ch <- prometheus.MustNewConstMetric(
			countDesc, prometheus.CounterValue, float64(s.Count), s.LabelValues...)
		sumDesc := prometheus.NewDesc(
			name+"_sum", "Zot metric "+s.Name, s.LabelNames, nil)
		ch <- prometheus.MustNewConstMetric(
			sumDesc, prometheus.CounterValue, s.Sum, s.LabelValues...)
	}
}

func RunZotExporter(ctrl *Controller) {
	controller = ctrl

	c := ZotCollector{}
	prometheus.MustRegister(c)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", controller.Config.Port), nil))
}
