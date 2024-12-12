package internal

import "github.com/startup_krasnodar_test/src/entities"

type Service interface {
	Loginer
	Registerer
}

type Loginer interface {
	Login(email, password string) (int, error)
}

type Registerer interface {
	Register(user *entities.User) (int, error)
	CheckEmail(email string, code int) (bool, error)
}
