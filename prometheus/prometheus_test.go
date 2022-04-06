package prometheus

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	os.Setenv(PROMETHEUS_PATH, "/metrics")
	os.Setenv(PROMETHEUS_PORT, "9100")

	prom, err := Start()
	if err != nil {
		t.Errorf("failed to start prom endpoint: %s", err.Error())
		return
	}

	res, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d%s", GetPort(), GetPath()))
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %v", res.StatusCode)
	}

	if err := prom.Close(); err != nil {
		t.Errorf("failed to close prometheus endpoint: %s", err.Error())
	}
}
