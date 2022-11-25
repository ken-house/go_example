package meta

import "github.com/prometheus/client_golang/prometheus"

// CustomizedCounterMetric 自定义指标收集器
var CustomizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "demo_server_say_hello_method_handle_count",
	Help: "Total number of RPCs handled on the server.",
}, []string{"name"})
