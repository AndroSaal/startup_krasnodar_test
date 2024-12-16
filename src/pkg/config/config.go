package config

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// конфиг для сервиса
type Config struct {
	DBConfig   DBConfig
	SrvConfig  SrvConfig
	MailConfig ServerMailAuthConf
}

// конфиг для соединения с БД
type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

// конфиг для сервера
type SrvConfig struct {
	Port    string        `yaml:"port"`
	Host    string        `yaml:"host"`
	Env     string        `yaml:"env"`
	Timeout time.Duration `yaml:"timeout"`
}

type ServerMailAuthConf struct {
	Login    string
	Password string
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func MustLoadConfig() *Config {
	//если не полуичлось загрузить конфиг - падаем с паникой
	if Config, err := LoadConfig(); err != nil {
		panic("troubles loading config: " + err.Error())
	} else {
		return Config
	}

}

func LoadConfig() (*Config, error) {

	var (
		pathToConfigDir  string
		nameOfConfigFile string
		databaseConfig   DBConfig
		serverConfig     SrvConfig
		mailConfig       ServerMailAuthConf
	)

	//подгружаем переменные из .env файла
	if err := godotenv.Load(".env/.app.env"); err != nil {
		return nil, err
	}

	//получаем из argv
	pathToConfigDir, nameOfConfigFile = getConfigPathFromARGV()

	//если не получилось, получаем из ENV
	if pathToConfigDir == "" {
		pathToConfigDir, nameOfConfigFile = getConfigPathFromENV()
	}

	//пустой путь
	if pathToConfigDir == "" {
		return nil, errors.New("path to config dir is empty")
	}

	//пустое имя файла
	if nameOfConfigFile == "" {
		return nil, errors.New("name of config file is empty")
	}

	//проверяем, существует ли такая директория
	if _, err := os.Stat(pathToConfigDir); os.IsNotExist(err) {
		return nil, err
	}

	//инициализируем имя, папку и тип конфига
	viper.AddConfigPath(pathToConfigDir)
	viper.SetConfigName(nameOfConfigFile)
	viper.SetConfigType("yaml")

	//загружаем конфиг
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//заполняем структуру ДБ
	if err := viper.UnmarshalKey("db", &databaseConfig); err != nil {
		return nil, err
	}

	//заполняем структуру сервера
	if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
		return nil, err
	}

	//заполняем структуру для атвторизации в почте
	//получаем host и порт сервера smtp
	if err := viper.UnmarshalKey("mail", &mailConfig); err != nil {
		return nil, err
	}

	//получаем пароль и логин для почты, с которой будем посылать письма
	if mailConfig.Login, mailConfig.Password = getMailCredentials(); mailConfig.Password == "" || mailConfig.Login == "" {
		return nil, errors.New("mail credentials are empty")
	}

	return &Config{
		DBConfig:   databaseConfig,
		SrvConfig:  serverConfig,
		MailConfig: mailConfig,
	}, nil

}

// получение пути к папке с конфигами и имя конфига из ARGV
func getConfigPathFromARGV() (string, string) {
	var (
		pathToConfigDir  string
		nameOfConfigFile string
	)

	//получаем из argv в pathToConfigDir значение после флага --config_path=,
	//дефолтное значение - пустая строка
	flag.StringVar(&pathToConfigDir, "config_path", "", "path to directory with config file")

	//аналогично с именем
	flag.StringVar(&nameOfConfigFile, "config_name", "", "name of config file")

	//парсим всё это дело
	flag.Parse()

	return pathToConfigDir, nameOfConfigFile
}

// получение пути к папке с конфигами и имя конфига из ENV
func getConfigPathFromENV() (string, string) {

	//получение из переменных окружения
	pathToConfigDir := os.Getenv("CONFIG_PATH")
	nameOfConfigFile := os.Getenv("CONFIG_NAME")

	return pathToConfigDir, nameOfConfigFile
}

func getMailCredentials() (string, string) {

	login := os.Getenv("MAIL_LOGIN")
	password := os.Getenv("MAIL_PASSWORD")

	return login, password
}
