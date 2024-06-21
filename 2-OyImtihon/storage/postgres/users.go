package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/modul"
	"strings"
	"time"
)
//o'zgartirishlardi filter qiladi ya'ni qaysilarini o'zgartirish kerakligini tekshiradi
type UpdateUser struct {
	UserId   *string    `json:"user_id"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Birthday *time.Time `json:"birthday"`
	Password *string    `json:"password"`
}

//Ekranga chiqedigon malumotlar filterlab olish uchun struct
type Filter struct {
	User_id    string
	Name       string
	Email      string
	Birthday   time.Time
	Password   string
	Deleted_at time.Time
}

//user uchun postgresql method qilib ulash uchun struct
type UsersRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{DB: db}
}

//user yaratish
func (u *UsersRepo) Create(user modul.User) error {
	tr, err := u.DB.Begin()

	//u.DB.Begin() muvaffaqiyatsiz bo'lsa, defer statement ishlamaydi.
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = u.DB.Exec(`INSERT INTO Users(name,email,birthday,password,created_at,updated_at) 
	VALUES($1,$2,$3,$4,$5,$6)`, user.Name, user.Email, user.Birthday, user.Password, time.Now(), time.Now())
	
	//Malumot kiritilmasa error qaytaradi
	if err != nil {
		return err
	}

	return nil
}
//user ni oqish ID bo'yicha
func (u *UsersRepo) Read(Id string) (modul.User, error) {
	tr, err := u.DB.Begin()
	//u.DB.Begin() muvaffaqiyatsiz bo'lsa, defer statement ishlamaydi.
	if err != nil {
		return modul.User{}, err
	}
	defer tr.Commit()
	user := modul.User{}

	row := u.DB.QueryRow("SELECT user_id,name,email,birthday,password FROM Users WHERE user_id = $1 AND deleted_at = 0", Id)

	err = row.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password)
	if err != nil {
		return modul.User{}, err
	}

	return user, nil
}

//user ni o'zgartirish 
func (u *UsersRepo) Update(updateFilter UpdateUser) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT user_id
	  FROM Users
	  WHERE deleted_at = 0 AND user_id = $1
	`

	if err := u.DB.QueryRow(query, *updateFilter.UserId).Err(); err != nil {
		return fmt.Errorf("user by this id not found: %v", err)
	}

	query = `
	  UPDATE Users SET 
	`

	if updateFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *updateFilter.Name)
	}

	if updateFilter.Email != nil {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *updateFilter.Email)
	}

	if updateFilter.Birthday != nil {
		params = append(params, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, *updateFilter.Birthday)
	}

	if updateFilter.Password != nil {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, *updateFilter.Password)
	}
	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.UserId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE user_id = $%d AND deleted_at = 0", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = u.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

// user ni o'chirish
func (u *UsersRepo) DELETE(Id string) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	res, err := u.DB.Exec(`UPDATE Users SET
		deleted_at = date_part('epoch', current_timestamp)::INT
	   where user_id = $1 and deleted_at = 0`, Id)
	if err != nil {
		return err
	}
	x, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if x == 0 {
		return fmt.Errorf("no such information is available or %s not found already deleted", Id)
	}
	return nil
}

//userlarni filterlab chiqarish
func (u *UsersRepo) GetAll(fUser modul.FilterUser) ([]modul.User, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return nil,err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT user_id, name, email, birthday, password FROM users WHERE deleted_at = 0 "

	if fUser.Name != "" {
		params["name"] = fUser.Name
		filter += "AND name = :name "
	}
	if fUser.Email != "" {
		params["email"] = fUser.Email
		filter += "AND email = :email "
	}
	if fUser.Age > 0 {
		params["age"] = fUser.Age
		filter += "AND EXTRACT(YEAR FROM age(birthday)) = :age"
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

	var users []modul.User
	for rows.Next() {
		var user modul.User

		err = rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday,&user.Password)

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
