package repository

import "github.com/startup_krasnodar_test/src/entities"

type Repository interface {
	AddNewUser(usr *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}
