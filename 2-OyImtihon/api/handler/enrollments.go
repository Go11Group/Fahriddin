package handler

import (
	"my_module/modul"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar Enrollment uchun

func (h *Handler)CreateEnrollments(ctx *gin.Context){
	enrollment := modul.Enrollment{}
	err := ctx.ShouldBindJSON(&enrollment)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			})
		return
	}
	res := modul.Enrollment{}
	if enrollment == res{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not found",
			})
		return
	}

	err = h.Enrollments.Create(enrollment)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Enrollment",
			})
		return
	}

	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment created successfully",
	})

}

func (h *Handler)ReadEnrollments(ctx *gin.Context){
	id := ctx.Param("id")
	enrollment,err := h.Enrollments.Read(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK,enrollment)
}

func (h *Handler)UpdateEnrollments(ctx *gin.Context){
	updateEnrollement := postgres.UpdateEnrollment{}
	err :=ctx.ShouldBindJSON(&updateEnrollement)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	err = h.Enrollments.Update(updateEnrollement)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":"Update not found",
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment updated successfully",
	})
}

func (h *Handler) DeleteEnrollments(ctx *gin.Context){
	id := ctx.Param("id")
	err :=h.Enrollments.DELETE(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":err.Error(),
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment deleted successfully",
	})
}

func (h *Handler) GetAllEnrollments(ctx *gin.Context){
	var enrollment modul.FilterEnrollment
	err :=ctx.ShouldBindQuery(&enrollment)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":err.Error(),
		})
		return
	}

	enrollments,err := h.Enrollments.GetAll(enrollment)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,enrollments)
}
