package main

import (
	"github.com/startup_krasnodar_test/src/internal"
	"github.com/startup_krasnodar_test/src/pkg/config"
	mylog "github.com/startup_krasnodar_test/src/pkg/log"
	"github.com/startup_krasnodar_test/src/transport"
)

func main() {

	// инициализируем конфиг, загружаем его
	config := config.MustLoadConfig()
	// fmt.Println(config.DBConfig, config.SrvConfig)

	// инициализируем логгер
	logger := mylog.NewLogger(config.SrvConfig.Env)
	logger.Info("Logger initialized")

	//TODO: Описать сервер, инициализировать, добавить хендлеры

	service := internal.NewAuth()
	handlers := transport.NewHandler(service)
	server, err := transport.NewServer(config.SrvConfig, handlers.InitRoutes(), logger)
	if err != nil {
		logger.Error(err.Error())
	}

	server.MustRun()

	//TODO: Описать БД, инициализировать

	//TODO: инкапсулировать сервер и бд в один метод сервисного уровня

	//TODO: внедрить заависимости между уровнями

	//TODO: запустить сервер

}
