package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/startup_krasnodar_test/src/entities"
)

func (h *Handler) singUp(c *gin.Context) {
	var (
		user entities.User
	)

	//заполняем структуру из пришедшего запроса
	if err := c.BindJSON(&user); err != nil {
		//ошибка в сериализации - отправляем ответ-ошибку
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	//создаем пользователя, вызывая метод сервисного слоя
	id, err := h.service.Register(&user)
	if err != nil {
		//ошибка в сервисе - отправляем ответ-ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//отправляем успешный ответ
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
