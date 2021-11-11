package routes

import (
	"github/Enocp/webapi-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("ap1/v1")
	{
		books := main.Group("book")
		{
			books.GET("/", controllers.ShowBook)
		}
	}

	return router
}