package transport

import "github.com/startup_krasnodar_test/src/internal"

// структура, обслуживающая все запросы от клиента
// в ней инкапсулируются вся бизнес-логика и работа
// с базой данных internal.Service - интерфейс бизнес логики
type Handler struct {
	service internal.Service
}

func NewHandler(service internal.Service) *Handler {
	return &Handler{
		service: service,
	}
}
