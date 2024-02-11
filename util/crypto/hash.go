package crypto

import "golang.org/x/crypto/bcrypt"

func Hash(raw string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func ValidateHash(raw, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	return err
}
