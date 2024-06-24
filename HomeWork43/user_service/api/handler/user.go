package handler

import (
	"my_module/model"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar User uchun

func (h *Handler) CreateUsers(ctx *gin.Context) {
	user := model.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := model.User{}
	if user == res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not found",
		})
		return
	}

	err = h.User.Create(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

}

func (h *Handler) GetByIdUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.User.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUsers(ctx *gin.Context) {
	updateUser := postgres.UpdateUser{}
	err := ctx.ShouldBindJSON(&updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = h.User.Update(updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Update not found",
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *Handler) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.User.DELETE(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	user := model.FilterUser{}
	err := ctx.ShouldBindQuery(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	users, err := h.User.GetAll(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR1": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
