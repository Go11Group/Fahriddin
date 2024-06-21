package modul

//Lesson uchun struct
type Lesson struct {
	LessonId string `json:"lesson_id"`
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
