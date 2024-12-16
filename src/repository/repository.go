package repository

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/startup_krasnodar_test/src/entities"
	"github.com/startup_krasnodar_test/src/pkg/config"

	_ "github.com/lib/pq"
)

// интерфейс для уровня репозитория
type RepositoryHandler interface {
	AddNewUser(usr *entities.User, code string) (int, error)
	GetUserById(userId int) (*entities.User, error)
	GetCodeFromEmail(id int, code string) (bool, error)
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
	columnUsername    = "username"
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
func (p *PostgreRepository) AddNewUser(usr *entities.User, code string) (int, error) {

	//транзакция начинается
	transaction, err := p.db.Begin()

	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	//формируем запрос к БД для добавления новой записи в таблицу users
	queryAddToUsrsTable := fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s) VALUES ($1, $2, $3, $4) RETURNING %s`,
		tableForUsers, columnUsername, passwordHash, columnEmail, columnIsVerifiied, columnUserId)

	//выполняем запрос
	row := transaction.QueryRow(queryAddToUsrsTable, usr.Username, usr.Password_hash, usr.Email, 0)

	//получаем id новой записи
	if err := row.Scan(&usr.Id); err != nil {
		transaction.Rollback()
		return 0, err
	}

	//формируем запрос к БД для добавления новой записи в таблицу codes
	queryAddToCodesTable := fmt.Sprintf(`INSERT INTO %s (%s, %s) VALUES ($1, $2)`,
		tableForCodes, columnCode, columnUser)

	//ошибка - откат
	if _, err := transaction.Exec(queryAddToCodesTable, code, usr.Id); err != nil {
		transaction.Rollback()
		return 0, err
	}

	return usr.Id, transaction.Commit()
}

func (p *PostgreRepository) GetUserById(userId int) (*entities.User, error) {
	//TODO: функция возвращает пользователя по id
	return nil, nil
}

// функция принимает емайл и код, если код совпадает с кодом в таблице кодов
// поменять поле isVerifiied в таблице юзеров на true
func (p *PostgreRepository) GetCodeFromEmail(id int, code string) (bool, error) {
	fi := "repository.PostgreRepository.GetCodeFromEmail"

	var (
		codeFromDB string
	)
	//формирование запроса к базе
	querySelectCode := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = $1`,
		columnCode, tableForCodes, columnUser)

	//выполняем запрос, получаем запись
	row := p.db.QueryRow(querySelectCode, id)

	if err := row.Scan(&codeFromDB); err != nil {
		p.log.Debug("%s: %v", fi, err)
	}

	if codeFromDB == code {
		//формируем текст запроса
		queryToAddVerification := fmt.Sprintf(`UPDATE %s SET %s = true WHERE %s = $1`,
			tableForUsers, columnIsVerifiied, columnUserId)

		//выполняем запрос
		if _, err := p.db.Exec(queryToAddVerification, id); err != nil {
			p.log.Debug("%s: %v", fi, err)
			return false, err
		}
		return true, nil
	} else {
		return false, nil
	}
}
