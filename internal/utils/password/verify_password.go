package password

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
