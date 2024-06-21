package modul

import "time"

//Hamma table uchun struct la mavjud


type FilterUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"Age"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type FilterCourse struct {
	Title       string `json:"title"`
	Description string `jons:"description"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

type FilterLesson struct {
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type FilterEnrollment struct {
	UserID         string     `json:"user_id"`
	CourseID       string     `json:"course_id"`
	EnrollmentDate *time.Time `json:"enrollment_date"`
	Offset         int        `json:"offset"`
	Limit          int        `json:"limit"`
}
