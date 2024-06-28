package handler

import (
	t "my_module/genproto/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetBusSchedule(ctx *gin.Context){
	number := t.Number{}

	err := ctx.ShouldBindJSON(&number)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	bus,err := h.Transport.GetBusSchedule(ctx,&number)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,bus)
}

func (h *Handler) TrackBusLocation(ctx *gin.Context){
	number := t.Number{}
	err := ctx.ShouldBindJSON(&number)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	location,err := h.Transport.TrackBusLocation(ctx,&number)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,location)
}

func (h *Handler) ReportTrafficJam(ctx *gin.Context){
	location := t.Location{}
	err := ctx.ShouldBindJSON(&location)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	
	status,err := h.Transport.ReportTrafficJam(ctx,&location)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,status)
}