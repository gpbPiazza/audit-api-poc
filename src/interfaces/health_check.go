package interfaces

type HealthCheckService interface {
	Pong() string
}
