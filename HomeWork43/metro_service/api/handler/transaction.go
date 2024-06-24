package handler

import (
	"my_module/model"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	transaction := model.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	res := model.Transaction{}

	if transaction != res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Not Found",
		})
	}
	err = h.Transaction.Create(transaction)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Transaction created succesfully",
	})
}

func (h *Handler) GetByIdTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := h.Transaction.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	updateTransaction := postgres.UpdateTransaction{}
	err := ctx.ShouldBindJSON(&updateTransaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	err = h.Transaction.Update(updateTransaction)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Updated transaction succesfully",
	})
}

func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Transaction.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"Message" : "Deleted transaction succesfully",
	})
}

func (h *Handler) GetAllTransaction(ctx *gin.Context){
	fTransaction := model.FilterTransaction{}
	err := ctx.ShouldBindQuery(&fTransaction)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	tracsactions,err := h.Transaction.GetAll(fTransaction)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,tracsactions)
}