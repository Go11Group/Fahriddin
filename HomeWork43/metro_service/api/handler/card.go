package handler

import (
	"my_module/model"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCard(ctx *gin.Context) {
	card := model.Card{}
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	res := model.Card{}

	if card != res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Not Found",
		})
	}
	err = h.Card.Create(card)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Card created succesfully",
	})
}

func (h *Handler) GetByIdCard(ctx *gin.Context) {
	id := ctx.Param("id")
	card, err := h.Card.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *Handler) UpdateCard(ctx *gin.Context) {
	updateCard := postgres.UpdateCard{}
	err := ctx.ShouldBindJSON(&updateCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	err = h.Card.Update(updateCard)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Updated card succesfully",
	})
}

func (h *Handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Card.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"Message" : "Deleted card succesfully",
	})
}

func (h *Handler) GetAllCard(ctx *gin.Context){
	fCard := model.FilterCard{}
	err := ctx.ShouldBindQuery(&fCard)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	cards,err := h.Card.GetAll(fCard)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,cards)
}