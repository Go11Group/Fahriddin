package handler

import (
	"my_module/modul"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar Course uchun

func (h *Handler) CreateCourses(ctx *gin.Context) {
	course := modul.Course{}
	err := ctx.ShouldBindJSON(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := modul.Course{}
	if course == res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not Found",
		},
		)
		return
	}
	err = h.Courses.Create(course)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create course",
		})
		return
	}

	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course created successfully",
	})

}

func (h *Handler) ReadCourses(ctx *gin.Context) {
	id := ctx.Param("id")
	course, err := h.Courses.Read(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (h *Handler) UpdateCourses(ctx *gin.Context) {
	updateCourse := postgres.UpdateCourse{}
	err := ctx.ShouldBindJSON(&updateCourse)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = h.Courses.Update(updateCourse)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Update not found",
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course updated successfully",
	})
}

func (h *Handler) DeleteCourses(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Courses.DELETE(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	// Agar xatolik bo'lmasa, davom eting
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course deleted successfully",
	})
}

func (h *Handler) GetAllCourses(ctx *gin.Context) {
	var course modul.FilterCourse
	err := ctx.ShouldBindQuery(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	courses, err := h.Courses.GetAll(course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}
