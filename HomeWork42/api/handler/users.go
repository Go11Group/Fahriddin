package handler

import (
	"my_module/modul"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar User uchun


func (h *Handler)CreateUsers(ctx *gin.Context){
	user := modul.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			})
		return
	}
	res := modul.User{}
	if user == res{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not found",
			})
		return
	}

	err = h.Users.Create(user)

	if err != nil{
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

func (h *Handler)ReadUsers(ctx *gin.Context){
	id := ctx.Param("id")
	user,err := h.Users.Read(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK,user)
}

func (h *Handler)UpdateUsers(ctx *gin.Context){
	updateUser := postgres.UpdateUser{}
	err :=ctx.ShouldBindJSON(&updateUser)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	err = h.Users.Update(updateUser)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":"Update not found",
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *Handler) DeleteUsers(ctx *gin.Context){
	id := ctx.Param("id")
	err :=h.Users.DELETE(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":err.Error(),
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func (h *Handler) GetAllUsers(ctx *gin.Context){
	user := modul.FilterUser{}
	err := ctx.ShouldBindQuery(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":err.Error(),
		})
		return
	}

	users,err := h.Users.GetAll(user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR1":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,users)
}

