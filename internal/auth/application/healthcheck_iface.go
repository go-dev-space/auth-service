package application

type Healthchecker interface {
	Execute() (string, error)
}
