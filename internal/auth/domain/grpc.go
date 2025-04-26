package domain

type GRPC interface {
	Send(string, int) (string, error)
}
