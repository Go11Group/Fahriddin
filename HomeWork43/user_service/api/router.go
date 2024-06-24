package api

import (
	"my_module/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handler handler.Handler) *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")

	user.GET("/:id", handler.GetByIdUsers)
	user.POST("", handler.CreateUsers)
	user.PUT("", handler.UpdateUsers)
	user.DELETE("/:id", handler.DeleteUsers)
	user.GET("/getall", handler.GetAllUsers)

	return router
}
