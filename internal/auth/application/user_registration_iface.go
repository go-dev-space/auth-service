package application

import (
	"context"

	"github.com/auth-service/internal/auth/interfaces/dto"
)

type Registrar interface {
	Execute(context.Context, *dto.Payload) (map[string]string, error)
}
