package api

import (
	"my_module/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handler handler.Handler) *gin.Engine {
	router := gin.Default()

	//User uchun API lar chiqarish
	user := router.Group("/user")
	user.GET("/:id", handler.ReadUsers)
	user.POST("", handler.CreateUsers)
	user.PUT("", handler.UpdateUsers)
	user.DELETE("/:id", handler.DeleteUsers)
	user.GET("/getall", handler.GetAllUsers)

	//Course uchun API lar chiqarish
	course := router.Group("/course")
	course.GET("/:id", handler.ReadCourses)
	course.POST("", handler.CreateCourses)
	course.PUT("", handler.UpdateCourses)
	course.DELETE("/:id", handler.DeleteCourses)
	course.GET("/getall", handler.GetAllCourses)

	//Enrollment uchun API lar chiqarish
	enrollment := router.Group("/enrollment")
	enrollment.GET("/:id", handler.ReadEnrollments)
	enrollment.POST("", handler.CreateEnrollments)
	enrollment.PUT("", handler.UpdateEnrollments)
	enrollment.DELETE("/:id", handler.DeleteEnrollments)
	enrollment.GET("/getall", handler.GetAllEnrollments)

	//Lesson uchun API lar chiqarish
	lesson := router.Group("/lesson")
	lesson.GET("/:id", handler.ReadLessons)
	lesson.POST("", handler.CreateLessons)
	lesson.PUT("", handler.UpdateLessons)
	lesson.DELETE("/:id", handler.DeleteLessons)
	lesson.GET("/getall", handler.GetAllLessons)

	//Qo'shimcha 5 ta API lar
	request := router.Group("/request")
	request.GET("/user/:id", handler.GetCoursesbyUser)
	request.GET("/course/:id",handler.GetLessonsbyCourse)
	request.GET("/enrollment/:id",handler.GetEnrolledUsersbyCourse)
	request.GET("/user",handler.SearchUsers)
	request.GET("/course",handler.GetMostPopularCourses)

	return router
}
