package domain

type Crypto interface {
	GenerateRandomString(int) string
	GenerateRandomInt(min, max int) int
	HashString(string) (string, error)
}
