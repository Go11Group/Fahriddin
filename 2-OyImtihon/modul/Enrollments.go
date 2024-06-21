package modul

import "time"

//Enrollment uchun struct
type Enrollment struct {
	EnrollmentId   string    `json:"enrollment_id"`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}
