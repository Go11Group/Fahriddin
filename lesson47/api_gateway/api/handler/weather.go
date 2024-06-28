package handler

import (
	w "my_module/genproto/weather"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	time := w.Time{}
	err := ctx.ShouldBindJSON(&time)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	weather, err := h.Weather.GetCurrentWeather(ctx, &time)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, weather)
}

func (h *Handler) GetWeatherForecast(ctx *gin.Context) {
	day := w.Day{}
	err := ctx.ShouldBindJSON(&day)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	weather, err := h.Weather.GetWeatherForecast(ctx, &day)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, weather)
}

func (h *Handler) ReportWeatherCondition(ctx *gin.Context) {
	weather := w.Weather{}
	err := ctx.ShouldBindJSON(&weather)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	status, err := h.Weather.ReportWeatherCondition(ctx, &weather)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, status)
}
