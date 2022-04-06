package prometheus

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetPort() int {
	str := os.Getenv(PROMETHEUS_PORT)
	port, err := strconv.Atoi(str)
	if err != nil {
		return 9100
	}

	return port
}

func GetPath() string {
	return os.Getenv(PROMETHEUS_PATH)
}

func GetUser() string {
	return os.Getenv(PROMETHEUS_USER)
}

func GetPassword() string {
	return os.Getenv(PROMETHEUS_PASSWORD)
}

// Start a Prometheus HTTP server to serve metrics
func Start(options ...func(*option)) (Closer, error) {
	srv := http.NewServeMux()
	addr := fmt.Sprintf("127.0.0.1:%d", GetPort())
	user, pass := GetUser(), GetPassword()
	opts := &option{
		ctx: context.Background(),
	}

	for _, opt := range options {
		opt(opts)
	}

	if user != "" && pass != "" {
		srv.HandleFunc(GetPath(), basicAuth(promhttp.Handler().ServeHTTP))
	} else {
		srv.Handle(GetPath(), promhttp.Handler())
	}

	lc := net.ListenConfig{}
	listen, err := lc.Listen(opts.ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	go http.Serve(listen, srv)

	return listen, nil
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		if username == GetUser() && password == GetPassword() {
			next.ServeHTTP(w, r)
			return
		}
	})
}
