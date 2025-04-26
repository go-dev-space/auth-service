package grpc

import (
	"context"
	"time"

	registration "github.com/auth-service/internal/auth/infrastructure/grpc/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCServiceRegistration is the data struct type
type GRPCServiceRegistration struct {
	target string
}

// NewGRPCServiceRegistration is the grpc service constructor function
func NewGRPCServiceRegistration(t string) *GRPCServiceRegistration {
	return &GRPCServiceRegistration{
		target: t,
	}
}

// Send creates a new grpc connection and client to the target and sends the email and userid
func (s *GRPCServiceRegistration) Send(email string, userid int) (string, error) {
	// create a new connection to target
	conn, err := grpc.NewClient(s.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	// create a new client
	client := registration.NewRegistrationClient(conn)
	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// send via grpc
	res, err := client.Do(ctx, &registration.Req{
		Email:  email,
		Userid: int64(userid),
	})
	if err != nil {
		return "", err
	}

	return res.Response, nil
}
