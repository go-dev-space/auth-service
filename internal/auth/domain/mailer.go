package domain

type Mailer interface {
	Send(string, string, string) error
}
