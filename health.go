package health

const (
	HealthCheckStatusUp   HealthCheckStatus = "UP"
	HealthCheckStatusDown HealthCheckStatus = "DOWN"
)

type HealthCheckStatus string

type HealthCheck struct {
	GlobalStatus HealthCheckStatus `json:"globalStatus"`
}
