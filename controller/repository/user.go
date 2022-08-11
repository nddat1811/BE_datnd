//go:generate mockgen -source user.go -destination ../testdata/mock_repository/user_gen.go
package repository

import "BE_datnd/data"

type UserRepository interface {
	FindUserByEmailAndPassword(email string, password string) (*data.User, error)
	FindUserByEmail(email string) (*data.User, error)
	FindUserByID(id int) (*data.User, error)
	HashPassword(password string) (string, error)
	UpdateHashedPassword(hasedPassword string, userID int) error
	CheckPasswordHash(password string, hash string) error
	CheckRole() int
}
