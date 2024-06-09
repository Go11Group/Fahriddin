package postgres

import (
	"database/sql"
	"my_module/modul"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo)Create(user modul.User) error{
	_,err := u.DB.Exec("INSERT INTO Users(name,age,email,pasword) ($1,$2,$3,$4)",
		user.Name,user.Age,user.Email,user.Password)

	if err != nil{
		return err
	}

	return nil
}
func (u *UserRepo)GetById(Id string) (modul.User,error){

	row := u.DB.QueryRow("Select * From Users Where id = $1",Id)
	user := modul.User{}
	err := row.Scan(&user.Id,&user.Name,&user.Age,&user.Email,&user.Password)

	if err != nil{
		return modul.User{},err
	}

	return user,nil
}

func (u *UserRepo)Update(user modul.User) error{
	_,err := u.DB.Exec(`Update Users Set 
	name = $1,
	age = $2,
	email = $3,
	password = $4
	Where id = $5
	`,user.Name,user.Age,user.Email,user.Password,user.Id)

	return err
}

func (u *UserRepo)Delete(Id string) error {
	_,err := u.DB.Exec("Delete From Users Where Id = $1",Id)
	return err
}