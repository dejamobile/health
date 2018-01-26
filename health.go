package health

const (
	Up   HealthCheckStatus = "UP"
	Down HealthCheckStatus = "DOWN"
)

type HealthCheckStatus string

type HealthCheck struct {
	Status  HealthCheckStatus `json:"status"`
	Modules []*Module         `json:"modules,omitempty"`
}

type Module struct {
	Name          string            `json:"name"`
	Status        HealthCheckStatus `json:"status"`
	HealthChecker HealthChecker     `json:"-"`
}

type HealthChecker interface {
	CheckHealth() HealthCheckStatus
}

// return a pointer to a Module struct
func NewModule(name string, healthChecker HealthChecker) *Module {
	return &Module{Name: name, HealthChecker: healthChecker}
}

// process the health check on all modules
func (healthCheck *HealthCheck) process() {
	numberOfModules := len(healthCheck.Modules)
	moduleStatusChannel := make(chan HealthCheckStatus, numberOfModules)
	for _, module := range healthCheck.Modules {
		go func(module *Module) {
			module.Status = module.HealthChecker.CheckHealth()
			moduleStatusChannel <- module.Status
		}(module)
	}
	for i := 0; i < numberOfModules; i++ {
		status := <-moduleStatusChannel
		if status == Down {
			// if one module is down, considers application down
			healthCheck.Status = Down
		}
	}
}
