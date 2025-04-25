package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidatorService struct {
	Validator *validator.Validate
}

func New() *ValidatorService {
	return &ValidatorService{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s *ValidatorService) Struct(data any) error {
	return s.Validator.Struct(data)
}

func (s *ValidatorService) Test(e error) (map[string]string, error) {
	found := make(map[string]string)
	if hits, ok := e.(validator.ValidationErrors); ok {
		for _, hit := range hits {
			switch hit.Tag() {
			case "required":
				found[hit.Field()] = fmt.Sprintf("%s is required", hit.Field())
			case "email":
				found[hit.Field()] = fmt.Sprintf("%s has wrong format", hit.Field())
			case "alphanum":
				found[hit.Field()] = fmt.Sprintf("%s must be alphanum", hit.Field())
			}
		}
		return found, errors.New("throw")
	}
	return found, nil
}
