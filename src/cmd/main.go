package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	// инициализируем сервис
	service := internal.NewAuth(config, logger)
	logger.Info("service initialized")

	// инициализируем обработчики
	handlers := transport.NewHandler(service)
	logger.Info("handlers initialized")

	// инициализируем сервер
	server, err := transport.NewServer(config.SrvConfig, handlers.InitRoutes(), logger)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("server initialized")
	}

	//инициализируем контекст
	ctx, cancel := context.WithTimeout(context.Background(), config.SrvConfig.Timeout)

	// запускаем сервер
	go func() {
		server.MustRun()
	}()

	//остановка сервера по сигналу системы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Waiting for SIGINT (pkill -2) or SIGTERM
	for {
		select {
		case <-stop:
			logger.Info("Server stopped by signal from system")
			server.Stop(ctx)
			cancel()
		case <-ctx.Done():
			logger.Info("Server stopped by timeout")
			server.Stop(ctx)
			cancel()
		}
	}

	//graceful shutdown, закрытие коннекта к базе, остановка сервера

	//TODO: Описать БД, инициализировать

	//TODO: инкапсулировать сервер и бд в один метод сервисного уровня

	//TODO: внедрить заависимости между уровнями

	//TODO: запустить сервер

}
