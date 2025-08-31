package usecase

import "github/ecommerce/domain/health"

type HealthService struct {
	Repo health.HealthRepository
}

func NewHealthService(repo health.HealthRepository) *HealthService {
	return &HealthService{Repo: repo}
}
func (r *HealthService) GetHealth() health.Health {
	return r.Repo.GetHealth()
}
