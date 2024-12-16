package internal

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

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
	VerifyEmail(id int, code string) (bool, error)
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
	logger *slog.Logger
}

func NewAuth(conf *config.Config, log *slog.Logger) *Auth {
	return &Auth{
		EmailSender:       NewMailSender(conf.MailConfig, log),
		RepositoryHandler: repository.NewPostgreRepository(conf.DBConfig, log),
		logger:            log,
	}
}

// TODO :реализовать
func (a *Auth) Login(email, password string) (int, error) {
	fmt.Printf("AT LOGIN")
	return 0, nil
}

func (a *Auth) Register(user *entities.User) (int, error) {
	fi := "internal.Auth.Reister"

	code := generateCode()

	if err := a.EmailSender.SendMail(user.Email, code); err != nil {
		a.logger.Debug("%s: Error sending email: %v", fi, err)
		return 0, err
	}

	id, err := a.RepositoryHandler.AddNewUser(user, code)
	if err != nil {
		a.logger.Debug("%s: Error adding new user: %v", fi, err)
		return 0, err
	}

	return id, nil
}

func generateCode() string {
	//генерируем семя для генератора
	// seed := [32]byte{'s', 'o', 'm', 'e', 'k', 'e', 'y', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'a', 'b', 'c', 'd', 'e', 'f'}

	// // инициализируем новый генератор ChaCha8
	// generator := rand.NewChaCha8(seed)

	rand.Seed(time.Now().UnixNano())

	// Генерация четырех случайных чисел от 0 до 9
	code := []int{
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10),
	}

	// Соединение чисел в одну строку
	finalcode := fmt.Sprintf("%d%d%d%d", code[0], code[1], code[2], code[3])

	return finalcode

}
