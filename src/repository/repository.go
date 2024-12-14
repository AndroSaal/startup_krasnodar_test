package repository

import (
	"log/slog"

	"github.com/startup_krasnodar_test/src/entities"
	"github.com/startup_krasnodar_test/src/pkg/config"
)

// интерфейс для уровня репозитория
type RepositoryHandler interface {
	AddNewUser(usr *entities.User) (int, error)
	GetUserById(userId string) (*entities.User, error)
	AddCodeForEmail(email, code string) error
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

func (p *PostgreRepository) AddNewUser(usr *entities.User) (int, error) {
	return 0, nil
}

func (p *PostgreRepository) GetUserById(userId string) (*entities.User, error) {
	return nil, nil
}

func (p *PostgreRepository) AddCodeForEmail(email, code string) error {
	return nil
}
