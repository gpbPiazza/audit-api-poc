package health_check

type HealthCheckService struct {
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (this *HealthCheckService) Pong() string {
	return "Pong :)"
}
