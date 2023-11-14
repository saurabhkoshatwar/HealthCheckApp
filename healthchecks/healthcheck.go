package healthchecks

import (
	"myhealthcheckapp/dtos"
	"net/http"
	"time"
)

type HealthCheckResult struct {
	URL             string
	IsUp            bool
	ResponseLatency time.Duration
}

func PerformHealthCheck(endpoint dtos.Endpoint) HealthCheckResult {
	startTime := time.Now()
	httpClient := http.Client{}

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, nil)
	if err != nil {
		return HealthCheckResult{endpoint.URL, false, 0}
	}

	if endpoint.Headers != nil {
		for key, value := range endpoint.Headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return HealthCheckResult{endpoint.URL, false, 0}
	}

	latency := time.Since(startTime)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 && latency < 500*time.Millisecond {
		return HealthCheckResult{endpoint.URL, true, latency}
	}

	return HealthCheckResult{endpoint.URL, false, latency}
}
