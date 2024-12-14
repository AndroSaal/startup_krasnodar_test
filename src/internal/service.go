package internal

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"

	"github.com/startup_krasnodar_test/src/entities"
	"github.com/startup_krasnodar_test/src/pkg/config"
	"github.com/startup_krasnodar_test/src/repository"
)

// интерфейс сервисного слоя
type ServiceHandler interface {
	Loginer
	Registerer
}

type Loginer interface {
	Login(email, password string) (int, error)
}

type Registerer interface {
	Register(user *entities.User) (int, error)
	VerifyEmail(email string, code int) (bool, error)
	SendMail(toEmail, mailBody string) error
}

type EmailSender interface {
	SendMail(toEmail, mailBody string) error
}

// имплементация интерфейса ServiceHandler
type Auth struct {
	//интерфейс - уровнь репозитория
	repository.RepositoryHandler

	//интерфейс - сущность для отправки писем
	EmailSender
}

func NewAuth(conf *config.Config, log *slog.Logger) *Auth {
	return &Auth{
		EmailSender:       NewMailSender(conf.MailConfig, log),
		RepositoryHandler: repository.NewPostgreRepository(conf.DBConfig, log),
	}
}

// TODO :реализовать
func (a *Auth) Login(email, password string) (int, error) {
	fmt.Printf("AT LOGIN")
	return 0, nil
}

func (a *Auth) Register(user *entities.User) (int, error) {
	fmt.Printf("AT REGISTER")

	code := (rand.Int() + rand.Int()) % (rand.Int() * 300)

	if err := a.EmailSender.SendMail(user.Email, strconv.Itoa(code)); err != nil {
		return 0, err
	}

	id, err := a.RepositoryHandler.AddNewUser(user, code)
	if err != nil {
		return 0, err
	}

	return id, nil
}
