package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/golang-api/controller"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controller.ShowBook)
			books.POST("/", controller.CreateBook)
			books.GET("/", controller.ShowBooks)
			books.PUT("/", controller.UpdateBook)
			books.DELETE("/:id", controller.DeleteBook)
		}
	}

	return router
}
