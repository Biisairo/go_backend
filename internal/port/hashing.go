package port

type Hashing interface {
	HashingPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}
