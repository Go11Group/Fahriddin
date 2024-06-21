package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/modul"
)
//Qo'shimcha Api lar

//1-apidagi userga tegishli course uchun struct 
type UserCourse struct {
	UserId  string         `json:"user_id"`
	Courses []modul.Course `json:"courses"`
}

//2-apidagi course ga tegishli lessons lar uchun struct
type CourseLessons struct {
	CourseId string         `json:"course_id"`
	Lessons  []Lesson `json:"lessons"`
}
//2-apidagi lesson uchun alohida struct
type Lesson struct{
	LessonId string `json:"lesson_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

//3-apidagi user boyicha course ga yozilganlarni olish uchun struct
type EnrolledUsers struct {
	CourseId string       `json:"course_id"`
	Users    []modul.User `json:"enrolled_users"`
}

//4-apidagi user ni age,email yoki name bo'yicha qidirish uchun struct
type SearchUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	AgeTo   int    `json:"age_to"`
	AgeFrom int    `json:"age_from"`
}
//4-apidagi qidiruvdan keyin malumotlarni olish uchun  struct
type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

//4-apidagi natijani olish uchun struct
type Results struct {
	Users []User `json:"results"`
}

//5-apidagi Course ni start_date va end_date boyicha qidirish uchun struct
type SearchCourse struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

//5-apidagi natijani olish uchun struct
type Course struct{
	CourseId string `json:"course_id"`
	CourseTitle string 	`json:"course_title"`
	EnrollmentsCount int `json:"enrollments_count"`
}
//5-api natijasi olish uchun struct
type GetCourses struct{
	TimePeriod SearchCourse `json:"time_period"`
	PopularCourses []Course `json:"popular_courses"`
}

//Api larni olish uchun postgresqlga ulash uchun struct
type RequestRepo struct {
	DB *sql.DB
}

func NewRequestRepo(db *sql.DB) *RequestRepo {
	return &RequestRepo{DB: db}
}

//1-api
func (r *RequestRepo) GetCoursesbyUser(Id string) (UserCourse, error) {

	rows, err := r.DB.Query(`
		SELECT 
			C.course_id, 
			C.title, 
			C.description
		FROM
			Users AS U
		JOIN 
			Enrollments as E
		ON 
		 	U.user_id = E.user_id
		JOIN
			Courses as C 
		ON 
			E.course_id = C.course_id
		WHERE 
			U.deleted_at = 0 AND C.deleted_at = 0 AND E.deleted_at = 0 AND U.user_id = $1
		`, Id)

	if err != nil {
		return UserCourse{}, err
	}

	var courses []modul.Course

	for rows.Next() {
		var course modul.Course
		err = rows.Scan(&course.CourseId, &course.Title, &course.Description)
		if err != nil {
			return UserCourse{}, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return UserCourse{}, err
	}
	userCourse := UserCourse{
		UserId:  Id,
		Courses: courses,
	}
	return userCourse, nil
}

//2-api
func (r *RequestRepo) GetLessonsbyCourse(Id string) (CourseLessons, error) {
	rows, err := r.DB.Query(`
		SELECT 
			L.lesson_id,L.title,L.content
		FROM
			Courses AS C
		JOIN
			Lessons AS L
		ON 
			C.course_id = L.course_id
		WHERE
			L.deleted_at = 0 AND C.deleted_at = 0 AND C.course_id = $1
		`, Id)
	if err != nil {
		return CourseLessons{}, err
	}

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		err = rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content)
		if err != nil {
			return CourseLessons{}, err
		}
		lessons = append(lessons, lesson)
	}
	courseLessons := CourseLessons{
		CourseId: Id,
		Lessons:  lessons,
	}
	return courseLessons, nil
}

//3-api
func (h *RequestRepo) GetEnrolledUsersbyCourse(Id string) (EnrolledUsers, error) {
	rows, err := h.DB.Query(`
		SELECT
			U.user_id,U.name,U.email
		FROM
			Users AS U
		JOIN
			Enrollments AS E
		ON
			U.user_id = E.user_id
		JOIN
			Courses AS C
		ON
			C.course_id = E.course_id
		WHERE
			U.deleted_at = 0 AND C.deleted_at = 0 AND E.deleted_at = 0 AND C.course_id = $1
	`, Id)
	if err != nil {
		return EnrolledUsers{}, err
	}
	var users []modul.User
	for rows.Next() {
		user := modul.User{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Email)
		if err != nil {
			return EnrolledUsers{}, err
		}
		users = append(users, user)
	}
	enrolledUsers := EnrolledUsers{
		CourseId: Id,
		Users:    users,
	}
	return enrolledUsers, nil
}


//4-api
func (r *RequestRepo) SearchUsers(search SearchUser) (Results, error) {
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)
	query := `SELECT user_id,name,email FROM Users WHERE deleted_at = 0 `

	if search.Name != "" {
		params["name"] = search.Name
		filter += " AND name = :name "
	}

	if search.Email != "" {
		params["email"] = search.Email
		filter += " AND email = :email "
	}

	if search.AgeTo > 0 && search.AgeFrom > 0 {
		params["age_to"] = search.AgeTo
		params["age_from"] = search.AgeFrom

		filter += " AND EXTRACT(YEAR FROM age(birthday)) between :age_to AND :age_from "
	}

	query += filter
	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return Results{}, err
	}
	fmt.Println(query, args)
	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.Name, &user.Email)
		if err != nil {
			return Results{}, err
		}
		users = append(users, user)
	}
	results := Results{
		Users: users,
	}
	return results, nil
}

//5-api
func(r *RequestRepo)GetMostPopularCourses(searchCourse SearchCourse) (GetCourses,error){

	// if searchCourse.StartDate == "" && searchCourse.EndDate == ""{
	// 	return GetCourses{},fmt.Errorf("start_date and end_date didn't come")
	// } 

	rows,err := r.DB.Query(`
	SELECT 
      c.course_id,
      c.title,
      COUNT(c.course_id) AS enrollment_count
    FROM 
      courses AS c 
    INNER JOIN
      enrollments AS e ON c.course_id = e.course_id
    WHERE 
      e.enrollment_date BETWEEN $1 AND $2
    GROUP BY
      c.course_id, c.title
    HAVING
      COUNT(c.course_id) = (
        SELECT 
          MAX(enroll_count) 
        FROM (
          SELECT 
            COUNT(e.course_id) AS enroll_count
          FROM 
            courses AS c 
          INNER JOIN
            enrollments AS e ON c.course_id = e.course_id
          GROUP BY
            c.course_id
        ) AS counts
      )
	`,searchCourse.StartDate,searchCourse.EndDate)
	if err != nil{
		return GetCourses{},err
	}

	var courses []Course

	for rows.Next(){
		var course Course
		err = rows.Scan(&course.CourseId,&course.CourseTitle,&course.EnrollmentsCount)
		if err != nil{
			return GetCourses{},err
		}
		courses = append(courses,course)
	}
	getCourses := GetCourses{
		TimePeriod: SearchCourse{},
		PopularCourses: courses,
	}
	return getCourses,err
}
