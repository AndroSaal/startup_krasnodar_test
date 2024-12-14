package repository

import (
	"log/slog"

	"github.com/startup_krasnodar_test/src/entities"
	"github.com/startup_krasnodar_test/src/pkg/config"
)

// интерфейс для уровня репозитория
type RepositoryHandler interface {
	AddNewUser(usr *entities.User, code int) (int, error)
	GetUserById(userId int) (*entities.User, error)
	GetCodeFromEmail(email string, code int) (bool, error)
}

// имплементация этого интерфейа
type PostgreRepository struct {
	cfg config.DBConfig
	log *slog.Logger
}

// конструктор эотй импелментации
func NewPostgreRepository(cfg config.DBConfig, log *slog.Logger) *PostgreRepository {
	return &PostgreRepository{
		cfg: cfg,
		log: log,
	}
}

func (p *PostgreRepository) AddNewUser(usr *entities.User, code int) (int, error) {
	//TODO: функция добавляет нового юзера в таблицу, добавлет ему код
	return 0, nil
}

func (p *PostgreRepository) GetUserById(userId int) (*entities.User, error) {
	//TODO: функция возвращает пользователя по id
	return nil, nil
}

func (p *PostgreRepository) GetCodeFromEmail(email string, code int) (bool, error) {
	//TODO: функция принимает емайл и код, если код совпадает с кодом в таблице юзера
	//поменять поле isVerifiied на true
	return false, nil
}
