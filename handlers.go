package health

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func HeathCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheck := checkApplicationHealth()
	healthCheckBytes, err := json.Marshal(healthCheck)
	if err != nil {
		log.Fatal("Application health check handler failed", err)
	}
	fmt.Fprintln(w, string(healthCheckBytes))
}

func checkApplicationHealth() *HealthCheck {
	return &HealthCheck{GlobalStatus: HealthCheckStatusUp}
}

func RegisterHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", HeathCheckHandler).Methods("GET")
}
