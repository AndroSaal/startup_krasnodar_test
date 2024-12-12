package main

import (
	"github.com/startup_krasnodar_test/src/pkg/config"
	mylog "github.com/startup_krasnodar_test/src/pkg/log"
)

func main() {

	// инициализируем конфиг, загружаем его
	config := config.MustLoadConfig()
	// fmt.Println(config.DBConfig, config.SrvConfig)

	// инициализируем логгер
	logger := mylog.NewLogger(config.SrvConfig.Env)
	logger.Info("Logger initialized")

	//TODO: Описать сервер, инициализировать, добавить хендлеры
	


	//TODO: Описать БД, инициализировать

	//TODO: инкапсулировать сервер и бд в один метод сервисного уровня

	//TODO: внедрить заависимости между уровнями

	//TODO: запустить сервер

}
