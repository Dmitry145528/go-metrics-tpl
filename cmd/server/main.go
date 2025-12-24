package main

import (
	"log"
	"net/http"

	"github.com/Dmitry145528/go-metrics-tpl.git/internal/handler"
	"github.com/Dmitry145528/go-metrics-tpl.git/internal/repository"
	"github.com/Dmitry145528/go-metrics-tpl.git/internal/service"
)

func main() {
	storage := repository.NewMemStorage()
	metricsService := service.NewMetricsService(storage)

	handler.SetMetricsService(metricsService)

	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, handler.UpdateMetric)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
