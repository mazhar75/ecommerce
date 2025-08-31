package memory

import "github/ecommerce/domain/health"

type HealthRepo struct {
	Health health.Health
}

var _ health.HealthRepository = &HealthRepo{}

func (r *HealthRepo) GetHealth() health.Health {
	return r.Health
}
