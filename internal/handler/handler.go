package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Dmitry145528/go-metrics-tpl.git/internal/service"
)

var metricsService *service.MetricsService

func SetMetricsService(s *service.MetricsService) {
	metricsService = s
}

func UpdateMetric(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(req.URL.Path, "/update/") {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	path := strings.TrimPrefix(req.URL.Path, "/update/")
	parts := strings.Split(path, "/")

	if len(parts) != 3 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	metricType := parts[0]
	metricName := parts[1]
	metricValue := parts[2]

	if metricName == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	switch metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		metricsService.UpdateGauge(metricName, value)

	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		metricsService.UpdateCounter(metricName, value)

	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(metricType + " " + metricName + " " + metricValue))
}
