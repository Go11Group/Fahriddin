package handler

import (
	"my_module/model"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStation(ctx *gin.Context) {
	station := model.Station{}
	err := ctx.ShouldBindJSON(&station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	res := model.Station{}

	if station != res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Not Found",
		})
	}
	err = h.Station.Create(station)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Station created succesfully",
	})
}

func (h *Handler) GetByIdStation(ctx *gin.Context) {
	id := ctx.Param("id")
	station, err := h.Station.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *Handler) UpdateStation(ctx *gin.Context) {
	updateStation := postgres.UpdateStation{}
	err := ctx.ShouldBindJSON(&updateStation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	err = h.Station.Update(updateStation)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Updated station succesfully",
	})
}

func (h *Handler) DeleteStation(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Station.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"Message" : "Deleted Station succesfully",
	})
}

func (h *Handler) GetAllStation(ctx *gin.Context){
	fStation := model.FilterStation{}
	err := ctx.ShouldBindQuery(&fStation)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	stations,err := h.Station.GetAll(fStation)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,stations)
}