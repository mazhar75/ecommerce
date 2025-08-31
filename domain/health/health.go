package health

type Health struct {
	Status  int    `json:"health"`
	Message string `json:"message"`
}

type HealthRepository interface {
	GetHealth() Health
}
