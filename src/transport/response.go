package transport

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

// структурка для ошибки в ответе
type errorResponse struct {
	Message string `json:"message"`
}

// структура для статуса в ответе
// type ststusResponse struct {
// 	Status string `json:"status"`
// }

// обработчик ошибки - вернуть ошибку и закрыть соединение
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	//возвращение ошибки внутри логгера (чтобы мы увидели)
	slog.Error(message)
	//возварщение ошибки в качестве ответа (чтобы увидел клиент)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
