package internal

import (
	"fmt"

	"github.com/startup_krasnodar_test/src/entities"
)

// интерфейс сервисного слоя
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

type Auth struct {
	//интерфейс уровня репозитория
	// repo repository.Repository
}

func NewAuth() *Auth {
	return &Auth{}
}

// TODO :реализовать
func (a *Auth) Login(email, password string) (int, error) {
	fmt.Printf("AT LOGIN")
	return 0, nil
}

func (a *Auth) Register(user *entities.User) (int, error) {
	fmt.Printf("AT REGISTER")
	return 0, nil
}

func (a *Auth) CheckEmail(email string, code int) (bool, error) {
	fmt.Printf("AT CHECK EMAIL")
	return false, nil
}
