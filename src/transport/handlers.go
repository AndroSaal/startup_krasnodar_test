package transport

import "github.com/startup_krasnodar_test/src/internal"

// структура, обслуживающая все запросы от клиента
// в ней инкапсулируются вся бизнес-логика и работа
// с базой данных internal.Service - интерфейс бизнес логики
type Handler struct {
	service internal.ServiceHandler
}

func NewHandler(service internal.ServiceHandler) *Handler {
	return &Handler{
		service: service,
	}
}
