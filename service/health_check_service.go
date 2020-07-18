package service

import (
	"github.com/DimitryEf/incrementer-api/config"
	"io"
	"log"
	"net/http"
)

type HealthCheckService struct {
	config *config.Config
}

func NewHealthCheckService(config *config.Config) *HealthCheckService {
	return &HealthCheckService{
		config: config,
	}
}

func (s *HealthCheckService) Run() {
	http.HandleFunc("/health-check", s.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(s.config.HealthCheckPort, nil))
}

func (s *HealthCheckService) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		s.config.Logger.Log.Error(err)
	}
}
