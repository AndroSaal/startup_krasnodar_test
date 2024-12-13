package internal

import (
	"fmt"
	"log/slog"

	"github.com/startup_krasnodar_test/src/entities"
	"github.com/startup_krasnodar_test/src/pkg/config"
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
	SendMail(toEmail, mailBody string) error
}

type EmailSender interface {
	SendMail(toEmail, mailBody string) error
}

// func NewService(conf *config.Config, log *slog.Logger) *Service {
// 	return &Service{
// 		Auth: NewAuth(conf, log),
// 	}
// }

type Auth struct {
	//интерфейс уровня репозитория TODO
	// repo repository.Repository
	//сущность для отправки писем
	EmailSender
}

func NewAuth(conf *config.Config, log *slog.Logger) *Auth {
	return &Auth{
		EmailSender: NewMailSender(conf.MailConfig, log),
	}
}

// TODO :реализовать
func (a *Auth) Login(email, password string) (int, error) {
	fmt.Printf("AT LOGIN")
	return 0, nil
}

func (a *Auth) Register(user *entities.User) (int, error) {
	fmt.Printf("AT REGISTER")
	a.EmailSender.SendMail(user.Email, "test")
	return 0, nil
}

func (a *Auth) CheckEmail(email string, code int) (bool, error) {
	fmt.Printf("AT CHECK EMAIL")
	return false, nil
}
