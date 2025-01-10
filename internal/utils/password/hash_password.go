package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error) {
	var hashPassword string
	if password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		hashPassword = string(hash)
	}
	return &hashPassword, nil
}
