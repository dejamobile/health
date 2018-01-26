package health

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func RegisterHealthCheck(router *mux.Router, modules ...*Module) {
	healthCheck := &HealthCheck{Status: Up}
	healthCheck.Modules = append(healthCheck.Modules, modules...)
	router.
		HandleFunc("/health", heathCheckHandler(healthCheck)).
		Methods("GET")
}

func heathCheckHandler(healthCheck *HealthCheck) func(w http.ResponseWriter, r *http.Request) {
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
