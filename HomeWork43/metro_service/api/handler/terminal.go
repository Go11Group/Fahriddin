package handler

import (
	"my_module/model"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTerminal(ctx *gin.Context) {
	terminal := model.Terminal{}
	err := ctx.ShouldBindJSON(&terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	res := model.Terminal{}

	if terminal != res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Not Found",
		})
	}
	err = h.Terminal.Create(terminal)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Terminal created succesfully",
	})
}

func (h *Handler) GetByIdTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	terminal, err := h.Terminal.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (h *Handler) UpdateTerminal(ctx *gin.Context) {
	updateTerminal := postgres.UpdateTerminal{}
	err := ctx.ShouldBindJSON(&updateTerminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	err = h.Terminal.Update(updateTerminal)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Updated terminal succesfully",
	})
}

func (h *Handler) DeleteTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Terminal.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"Message" : "Deleted terminal succesfully",
	})
}

func (h *Handler) GetAllTerminal(ctx *gin.Context){
	fTerminal := model.FilterTerminal{}
	err := ctx.ShouldBindQuery(&fTerminal)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	terminals,err := h.Terminal.GetAll(fTerminal)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,terminals)
}