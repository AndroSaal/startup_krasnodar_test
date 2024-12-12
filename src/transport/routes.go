package transport

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp) //регистрация
		// auth.POST("/signin", h.SignIn) //TODO - логин
	}

	return router

}
