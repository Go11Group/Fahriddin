package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from student s
					left join course c on c.id = s.course_id `)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course_Id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
					left join course c on c.id = s.course_id where s.id = $1`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course_Id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *StudentRepo)Create(user model.User) error {
	fmt.Print("Name : ")
	fmt.Scan(&user.Name)
	fmt.Print("Age : ")
	fmt.Scan(&user.Age)
	fmt.Print("Gender -> 'M' or 'F': ")
	fmt.Scan(&user.Gender)
	fmt.Print("Course Id : ")
	fmt.Scan(&user.Course_Id)
	_,err := u.Db.Exec("Insert Into Student(Name,Age,Gender,Course_id) Values($1,$2,$3,$4)",
	user.Name,user.Age,user.Gender,user.Course_Id)

	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepo)Update(Id string,user model.User) error {
	fmt.Print("Name : ")
	fmt.Scan(&user.Name)
	fmt.Print("Age : ")
	fmt.Scan(&user.Age)
	fmt.Print("Gender -> 'M' or 'F' : "  )
	fmt.Scan(&user.Gender)
	fmt.Print("Course Id : ")
	fmt.Scan(&user.Course_Id)

	_,err := u.Db.Exec("Update Student Set Name = $1, Age = $2, Gender = $3,Course_id = $4 Where Id = $5",
	user.Name,user.Age,user.Gender,user.Course_Id,Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *StudentRepo)Delete(id string) error {

	_,err := u.Db.Exec("Delete from student Where Id = $1",id)
	if err != nil {
		return err
	}

	return nil
}
