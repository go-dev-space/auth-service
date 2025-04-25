package application

import (
	"context"

	"github.com/auth-service/internal/auth/domain"
	"github.com/auth-service/internal/auth/infrastructure/crypto"
	"github.com/auth-service/internal/auth/infrastructure/mailer"
	"github.com/auth-service/internal/auth/infrastructure/validator"
	"github.com/auth-service/internal/auth/interfaces/dto"
	"github.com/auth-service/pkg/logs"
)

type RegistrationUserUseCase struct {
	store     domain.Store
	logger    *logs.Logwriter
	crypto    domain.Crypto
	mailer    domain.Mailer
	validator domain.Validator
}

func NewRegistrationUserUseCase(s domain.Store, l *logs.Logwriter) *RegistrationUserUseCase {
	return &RegistrationUserUseCase{
		store:     s,
		logger:    l,
		crypto:    crypto.New(),
		mailer:    mailer.New("", "", "", 465),
		validator: validator.New(),
	}
}

func (uc *RegistrationUserUseCase) Execute(ctx context.Context, p *dto.Payload) (map[string]string, error) {

	// validate payload
	err := uc.validator.Struct(p)
	if err != nil {
		msg, err := uc.validator.Test(err)
		if err != nil {
			uc.logger.Error.Println(err.Error())
			return msg, err
		}
		return nil, err
	}

	// generate username
	username := uc.crypto.GenerateRandomString(8)

	// hash password
	hash, err := uc.crypto.HashString(p.Password)
	if err != nil {
		uc.logger.Error.Println(err.Error())
		return nil, err
	}

	// create new user
	user := domain.NewUser(username, p.Email, hash)

	// save user & profile to db
	if err := uc.store.Save(ctx, user); err != nil {
		uc.logger.Error.Println(err.Error())
		return map[string]string{"store": err.Error()}, err
	}

	// TODO: send grpc / mail / push to nats message broker

	return nil, nil
}
