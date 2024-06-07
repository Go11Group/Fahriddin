package postgres

import (
	"database/sql"
	"my_module/modul"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(User modul.User) error {
	tr, err := u.DB.Begin()
	defer tr.Commit()
	if err != nil {
		return err
	}
	_, err = u.DB.Exec("INSERT INTO Users(username,email,password) VALUES($1,$2,$3)",
		User.Username, User.Email, User.Password)
	if err != nil {
		return err
	}
	return nil
}

// func (u *UserRepo) GetAllProducts() ([]modul.User, error) {
// 	tr, err := u.DB.Begin()
// 	defer tr.Commit()

// 	if err != nil {
// 		return []modul.User{}, err
// 	}
// 	users := []modul.User{}
// 	user := modul.User{}

// 	rows, err := u.DB.Query("SELECT * FROM Users")
// 	if err != nil {
// 		return []modul.User{}, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
// 		if err != nil {
// 			return []modul.User{}, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

func (u *UserRepo) GetById(id int) (modul.User, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return modul.User{}, nil
	}
	defer tr.Commit()

	user := modul.User{}

	row := u.DB.QueryRow("Select * from Users Where id = $1", id)

	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return modul.User{}, nil
	}
	return user, nil
}

func (u *UserRepo) Update(user modul.User, id int) error {
	tr, err := u.DB.Begin()
	defer tr.Commit()

	if err != nil {
		return nil
	}
	_, err = u.DB.Exec(`Update Users Set 
	username = $1,
	email = $2,
	password = $3 
	Where id = $4`, user.Username, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) Delate(id int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = u.DB.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
