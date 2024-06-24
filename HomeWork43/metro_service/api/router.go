package api

import (
	"my_module/api/handler"
	"github.com/gin-gonic/gin"
)

func Router(handler handler.Handler) *gin.Engine{
	router := gin.Default()

	card := router.Group("/card")

	card.GET("/:id",handler.GetByIdCard)
	card.POST("",handler.CreateCard)
	card.PUT("",handler.UpdateCard)
	card.DELETE("/:id",handler.DeleteCard)
	card.GET("/getall",handler.GetAllCard)

	station := router.Group("/station")

	station.GET("/:id",handler.GetByIdStation)
	station.POST("",handler.CreateStation)
	station.PUT("",handler.UpdateStation)
	station.DELETE("/:id",handler.DeleteStation)
	station.GET("/getall",handler.GetAllStation)

	return router
}