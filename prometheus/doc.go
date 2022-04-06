package prometheus

// Prometheus package to expose application metrics
//
// This package create a dedicated HTTP server using CC_METRICS_PROMETHEUS_PORT
// because, it might not be interesting to expose internal metrics world wide.
// In most of the case, the default CC_METRICS_PROMETHEUS_PORT is 8080.
// So you won't be able to start your main application server on the same port.
// The easier way to use this module is to CC_METRICS_PROMETHEUS_PORT to another port (Ex: 9100)
