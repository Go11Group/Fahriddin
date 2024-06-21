package main

import (
	"my_module/api"
	"my_module/api/handler"
	"my_module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	//Postgresql ni ulab qoyish
	us := postgres.NewUsersRepo(db)
	cr := postgres.NewCoursesRepo(db)
	ls := postgres.NewLessonsRepo(db)
	en := postgres.NewEnrollmentsRepo(db)
	rq := postgres.NewRequestRepo(db)

	handler := handler.Handler{
		Users:       us,
		Courses:     cr,
		Lessons:     ls,
		Enrollments: en,
		Request:     rq,
	}

	router := api.Router(handler)

	router.Run(":8080")
}
