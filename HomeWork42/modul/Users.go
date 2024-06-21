package modul

import "time"

//User uchun struct
type User struct {
	UserId   string    `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
}
