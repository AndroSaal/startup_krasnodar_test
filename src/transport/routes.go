package transport

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.singUp) //регистрация

		verify := auth.Group("/verify-email")
		{
			verify.POST("", h.verifyEmail)        //TODO - верификация почты
			verify.POST("/resend", h.resendEmail) //TODO - повторная отправка письма
		}
		// auth.POST("/signin", h.SignIn) //TODO - логин
	}

	return router

}
