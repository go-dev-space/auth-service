package crypto

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

type CryptoService struct{}

func New() *CryptoService {
	return &CryptoService{}
}

func (s CryptoService) GenerateRandomString(length int) string {
	randomString := make([]byte, length)

	for i := 0; i < length; i++ {
		if i == 0 {
			// ASCI: 65-90 = A-Z
			randomString[i] = byte(s.GenerateRandomInt(65, 90))
			continue
		}
		// ASCI: 97-122: a-z
		randomString[i] = byte(s.GenerateRandomInt(97, 122))
	}
	return string(randomString)
}

func (s CryptoService) GenerateRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (s CryptoService) HashString(value string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(value), 12)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
