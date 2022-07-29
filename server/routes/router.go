package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	// caminho central da API
	main := router.Group("api/v1")
	{
		// caminho especifico
		users := main.Group("users")
		{
			users.GET("/")
		}
	}

	return router
}
