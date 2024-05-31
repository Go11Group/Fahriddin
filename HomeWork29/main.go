package main

import (
	"my_module/postgres"
)

func main() {
	Db, err := postgres.ConnectGORM()
	if err != nil {
		panic(err)
	}
	us := postgres.NewUserRepo(Db)

	// user := model.User{FirstName: "Fahriddin", Lastname: "Rahimberdiyev", Email: "nswfn@gmail.com", Password: "1234", Age: 20, Field: "Programer", Gender: "M", IsEmployee: true}
	// us.Create(user)

	// users, err := us.GetAllUsers()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	// user := model.User{
	// 	Model: gorm.Model{
	// 		ID:        1,
	// 		CreatedAt: time.Now(),
	// 		UpdatedAt: time.Now(),
	// 	},
	// 	FirstName: "Diyorbek",
	// 	Lastname: "Nematov",
	// 	Field: "Programmer",
	// }
	// us.Update(user)

	// id := uint(1)
	// us.Delete(id)
}
