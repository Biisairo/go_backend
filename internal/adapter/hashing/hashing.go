package hashing

import "golang.org/x/crypto/bcrypt"

type HashingImpl struct {
}

func (h *HashingImpl) HashingPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err

}

func (h *HashingImpl) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
