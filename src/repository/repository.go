package repository

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
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
	cfg config.DBConfig // конфигурация
	log *slog.Logger    // логгер
	db  *sqlx.DB        // коннект
}

// конструктор эотй импелментации
func NewPostgreRepository(cfg config.DBConfig, log *slog.Logger) *PostgreRepository {

	//функция сначала создает коннект, пингует его, если возникает ошибка, коннект закрывается
	//если коннект закрывается - вызывается паника
	db := sqlx.MustConnect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode))

	return &PostgreRepository{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (p *PostgreRepository) AddNewUser(usr *entities.User, code int) (int, error) {
	//TODO: функция добавляет нового юзера в таблицу, добавлет ему код
	//который отправляется на почту
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
