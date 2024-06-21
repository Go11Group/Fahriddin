package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/modul"
	"strings"
	"time"
)
//Update qilishda filter qib ozgartiradi
type UpdateCourse struct {
	CourseId    *string `json:"course_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

//funcsiyalardi bazaga method qib olish uchun
type CoursesRepo struct {
	DB *sql.DB
}

func NewCoursesRepo(db *sql.DB) *CoursesRepo {
	return &CoursesRepo{DB: db}
}

//Course yaratish
func (c *CoursesRepo) Create(course modul.Course) error {
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = c.DB.Exec(`INSERT INTO Courses(title,description,created_at,updated_at)
	VALUES($1,$2,$3,$4)`, course.Title, course.Description, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

//Course id boyicha oqish
func (c *CoursesRepo) Read(Id string) (modul.Course, error) {
	tr, err := c.DB.Begin()
	if err != nil {
		return modul.Course{}, err
	}
	defer tr.Commit()

	course := modul.Course{}
	row := c.DB.QueryRow("SELECT course_id,title,description FROM Courses Where course_id = $1 AND deleted_at = 0", Id)
	err = row.Scan(&course.CourseId,&course.Title, &course.Description)

	if err != nil {
		return modul.Course{}, err
	}

	return course, nil
}

//Course ni filter qb o'zgartirish
func (c *CoursesRepo) Update(updateFilter UpdateCourse) error {
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT course_id
	  FROM Courses
	  WHERE deleted_at = 0 AND course_id = $1
	`
	err = c.DB.QueryRow(query, *updateFilter.CourseId).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Courses SET 
	`

	if updateFilter.Title != nil {
		params = append(params, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *updateFilter.Title)
	}

	if updateFilter.Description != nil {
		params = append(params, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, *updateFilter.Description)
	}

	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.CourseId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE course_id = $%d AND deleted_at = 0", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = c.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

//Course id bo'yicha o'chirish
func (c *CoursesRepo) DELETE(Id string) (error) {
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	res, err := c.DB.Exec(`UPDATE courses SET
		deleted_at = date_part('epoch', current_timestamp)::INT
	   where course_id = $1 and deleted_at = 0`, Id)
	if err != nil {
		return err
	}
	x,err := res.RowsAffected()
	
	if err != nil{
		return err
	}

	if x == 0{
		return fmt.Errorf("no such information is available or %s not found already deleted", Id)
	}


	return nil
}

//filter qib barcha ma'lumotlarni chiqarish
func (c *CoursesRepo) GetAll(fCourse modul.FilterCourse) ([]modul.Course, error) {
	tr, err := c.DB.Begin()
	if err != nil {
		return nil,err
	}
	defer tr.Commit()
	
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT course_id,title,description FROM courses WHERE deleted_at = 0 "

	if fCourse.Title != "" {
		params["title"] = fCourse.Title
		filter += "AND title = :title "
	}
	if fCourse.Description != "" {
		params["description"] = fCourse.Description
		filter += "AND description = :description "
	}

	if fCourse.Limit > 0 {
		params["limit"] = fCourse.Limit
		filter += " limit :limit "
	}

	if fCourse.Offset > 0 {
		params["offset"] = fCourse.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var courses []modul.Course
	for rows.Next() {
		var course modul.Course

		err = rows.Scan(&course.CourseId,&course.Title,&course.Description)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, err
}
