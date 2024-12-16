package repository

import (
	"fmt"
	"log/slog"
	"strconv"

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

// таблица
const tableForUsers = "users"

// поля таблицы
const (
	columnUserId      = "id"
	columnEmail       = "email"
	conumnUsername    = "username"
	passwordHash      = "password_hash"
	columnIsVerifiied = "is_email_verified"
)

// таблица кодов и верификаций
const tableForCodes = "codes"

// поля таблиц
const (
	columnCode = "email_code"
	columnUser = "user_id"
)

// функция добавляет нового юзера в таблицу, добавлет ему код
// который отправляется на почту
func (p *PostgreRepository) AddNewUser(usr *entities.User, code int) (int, error) {

	//транзакция начинается
	transaction, err := p.db.Begin()

	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	//формируем запрос к БД для добавления новой записи в таблицу users
	queryAddToUsrsTable := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3) RETURNING %s",
		tableForUsers, conumnUsername, passwordHash, columnEmail, columnUserId)

	//выполняем запрос
	row := transaction.QueryRow(queryAddToUsrsTable, usr.Username, usr.Password_hash, usr.Email)

	//получаем id новой записи
	if err := row.Scan(&usr.Id); err != nil {
		transaction.Rollback()
		return 0, err
	}

	//формируем запрос к БД для добавления новой записи в таблицу codes
	queryAddToCodesTable := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2)",
		tableForCodes, columnCode, strconv.Itoa(usr.Id))

	//ошибка - откат
	if _, err := transaction.Exec(queryAddToCodesTable, code); err != nil {
		transaction.Rollback()
		return 0, err
	}

	return usr.Id, transaction.Commit()
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
