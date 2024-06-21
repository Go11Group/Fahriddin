package handler

import (
	"my_module/modul"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar Lesson uchun

func (h *Handler) CreateLessons(ctx *gin.Context) {
	lesson := modul.Lesson{}
	err := ctx.ShouldBindJSON(&lesson)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := modul.Lesson{}
	if lesson == res{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not found",
		})
		return
	}
	err = h.Lessons.Create(lesson)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create lessons",
		})
		return
	}

	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lessons created successfully",
	})

}

func (h *Handler) ReadLessons(ctx *gin.Context) {
	id := ctx.Param("id")
	lesson, err := h.Lessons.Read(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lesson)
}

func (h *Handler) UpdateLessons(ctx *gin.Context) {
	updateLesson := postgres.UpdateLesson{}
	err := ctx.ShouldBindJSON(&updateLesson)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = h.Lessons.Update(updateLesson)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Update not found",
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lesson updated successfully",
	})
}

func (h *Handler) DeleteLessons(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Lessons.DELETE(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lesson deleted successfully",
	})
}

func (h *Handler) GetAllLessons(ctx *gin.Context) {
	var lesson modul.FilterLesson
	err := ctx.ShouldBindQuery(&lesson)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	lessons, err := h.Lessons.GetAll(lesson)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, lessons)
}
