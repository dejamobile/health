package health

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func RegisterHealthCheck(router *mux.Router, modules ...*Module) (*mux.Router) {
	healthCheck := &HealthCheck{Status: Up}
	healthCheck.Modules = append(healthCheck.Modules, modules...)
	router.
		HandleFunc("/health", healthCheckHandler(healthCheck)).
		Methods("GET")
	return router

}

func NewHandler(modules ...*Module) http.HandlerFunc {
	healthCheck := &HealthCheck{Status: Up}
	healthCheck.Modules = append(healthCheck.Modules, modules...)
	return func(writer http.ResponseWriter, request *http.Request) {
		healthCheck.process()
		healthCheckBytes, err := json.Marshal(healthCheck)
		if err != nil {
			log.Fatal("Application health check handler failed", err)
		}
		fmt.Fprintln(writer, string(healthCheckBytes))
	}
}

func healthCheckHandler(healthCheck *HealthCheck) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// processing healthcheck
		healthCheck.process()
		// encoding struct as json
		healthCheckBytes, err := json.Marshal(healthCheck)
		if err != nil {
			log.Fatal("Application health check handler failed", err)
		}
		fmt.Fprintln(w, string(healthCheckBytes))
	}
}
