package handler

import "my_module/storage/postgres"

//API ga responce qaytaruvchi funksiyalar uchun data bazani chaqirib olish

type Handler struct{
	Users *postgres.UsersRepo
	Courses *postgres.CoursesRepo
	Lessons *postgres.LesseonsRepo
	Enrollments *postgres.EnrollmentsRepo
	Request *postgres.RequestRepo
}

