package handler

import (
	"fmt"
	"my_module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ga responce qaytaruvchi funksiyalar qo'shimcha 5ta API uchun

func (h *Handler) GetCoursesbyUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userCourse, err := h.Request.GetCoursesbyUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, userCourse)
}

func (h *Handler) GetLessonsbyCourse(ctx *gin.Context) {
	id := ctx.Param("id")

	courseLessons, err := h.Request.GetLessonsbyCourse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, courseLessons)
}

func (h *Handler) GetEnrolledUsersbyCourse(ctx *gin.Context) {
	id := ctx.Param("id")

	enrolledUsers, err := h.Request.GetEnrolledUsersbyCourse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, enrolledUsers)
}

func (h *Handler) SearchUsers(ctx *gin.Context) {
	search := postgres.SearchUser{}
	err := ctx.ShouldBindQuery(&search)
	fmt.Println(search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	results, err := h.Request.SearchUsers(search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (h *Handler) GetMostPopularCourses(ctx *gin.Context) {
	start_date, hasSta := ctx.GetQuery("start_date")
	if !hasSta {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "star_date not found",
		})
		return
	}
	end_date,hasSta := ctx.GetQuery("end_date")
	if !hasSta {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "end_date not found",
		})
		return
	}

	getCourses, err := h.Request.GetMostPopularCourses(postgres.SearchCourse{
		StartDate: start_date,
		EndDate: end_date,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, getCourses)
}
