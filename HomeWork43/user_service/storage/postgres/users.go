package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/model"
	"strings"
)

type UpdateUser struct {
	Id    *string `json:"id"`
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
	Age   *int    `json:"age"`
}

type UsersRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{DB: db}
}

func (u *UsersRepo) Create(user model.User) error {
	tr, err := u.DB.Begin()

	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = u.DB.Exec(`INSERT INTO Users(name,phone,age) 
	VALUES($1,$2,$3)`, user.Name, user.Phone, user.Age)

	if err != nil {
		return err
	}

	return nil
}

func (u *UsersRepo) GetById(Id string) (model.User, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return model.User{}, err
	}
	defer tr.Commit()
	user := model.User{}

	row := u.DB.QueryRow("SELECT user_id,name,phone,age FROM Users WHERE id = $1 ", Id)

	err = row.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UsersRepo) Update(updateFilter UpdateUser) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT id
	  FROM Users
	  WHERE user_id = $1
	`

	if err := u.DB.QueryRow(query, *updateFilter.Id).Err(); err != nil {
		return fmt.Errorf("user by this id not found: %v", err)
	}

	query = `
	  UPDATE Users SET 
	`

	if updateFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *updateFilter.Name)
	}

	if updateFilter.Phone != nil {
		params = append(params, fmt.Sprintf("phone = $%d", len(args)+1))
		args = append(args, *updateFilter.Phone)
	}

	if updateFilter.Age != nil {
		params = append(params, fmt.Sprintf("age = $%d", len(args)+1))
		args = append(args, *updateFilter.Age)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.Id)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d ", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = u.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (u *UsersRepo) DELETE(Id string) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = u.DB.Exec("DELETE FROM Users WHERE id = $1 ", Id)

	return err
}

func (u *UsersRepo) GetAll(fUser model.FilterUser) ([]model.User, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT id, name, phone, age  FROM users WHERE true  "

	if fUser.Name != "" {
		params["name"] = fUser.Name
		filter += "AND name = :name "
	}
	if fUser.Phone != "" {
		params["phone"] = fUser.Phone
		filter += "AND phone = :phone "
	}
	if fUser.Age > 0 {
		params["age"] = fUser.Age
		filter += "AND age = :age"
	}

	if fUser.Limit > 0 {
		params["limit"] = fUser.Limit
		filter += " limit :limit "
	}

	if fUser.Offset > 0 {
		params["offset"] = fUser.Offset
		filter += " offset :offset "
	}

	query += filter
	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := u.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
