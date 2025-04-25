package application

type HealthcheckUseCase struct {
}

func NewHealthcheckUseCase() *HealthcheckUseCase {
	return &HealthcheckUseCase{}
}

func (uc *HealthcheckUseCase) Execute() (string, error) {
	return "pong", nil
}
