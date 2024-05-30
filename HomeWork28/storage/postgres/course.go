package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) GetAllCourse() ([]model.Course, error) {
	rows, err := c.DB.Query("Select * from Course")
	if err != nil {
		return nil, err
	}

	var courses []model.Course
	var course model.Course
	for rows.Next() {
		err = rows.Scan(&course.Id,&course.Name,&course.Field)
		if err != nil {
			return nil, err
		}
		courses= append(courses,course)
	}

	return courses, nil
}

func (c * CourseRepo) GetByID(id string) (model.Course, error) {
	var Course model.Course

	row := c.DB.QueryRow(`Select * From Course where id = $1`,id)
	err := row.Scan(&Course.Id,&Course.Name,&Course.Field)
	if err != nil {
		return model.Course{}, err
	}

	return Course, nil
}

func (c *CourseRepo)Create(course model.Course) error {
	fmt.Print("Course Name : ")
	fmt.Scan(&course.Name)
	fmt.Print("Field : ")
	fmt.Scan(&course.Field)
	_,err := c.DB.Exec("Insert Into Course(Name,Field) Values($1,$2)",
	course.Name,course.Field)

	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo)Update(Id string,course model.Course) error {
	fmt.Print("Course Name : ")
	fmt.Scan(&course.Name)
	fmt.Print("Field : ")
	fmt.Scan(&course.Field)
	_,err := c.DB.Exec("Update Course Set Name = $1, Field = $2 Where Id = $3",
	course.Name,course.Field,Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *CourseRepo)Delete(id string) error {

	_,err := c.DB.Exec("Delete from Course Where Id = $1",id)
	if err != nil {
		return err
	}

	return nil
}






