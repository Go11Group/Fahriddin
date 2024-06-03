package main

import (
	"mymode/postgres"
)

func main() {
	Db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	// pr := postgres.NewProductRepo(Db)
	// us := postgres.NewUsersRepo(Db)

	// product := modul.Product{
	// 	Name:           "Smartfon",
	// 	Description:    "Uyali aloqa",
	// 	Price:          150000.0,
	// 	Stock_quantity: 200,
	// }
	// err = pr.Create(product)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Add Product Succes!!!")

	// res, err := pr.GetAllProducts()
	// if err != nil {
	// 	panic(err)
	// }
	// for _,r := range res{
	// 	fmt.Println(r)
	// }

	// product := modul.Product{
	// 	Id:16,
	// 	Name:           "Iphone",
	// 	Description:    "Smartfon",
	// 	Price:          9000000.0,
	// 	Stock_quantity: 100,
	// }
	// err = pr.Update(product,product.Id)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Succes Update Products!!!")

	// id := 16
	// err = pr.Delate(id)
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Printf("ID - %d Delete Products\n", id)

	// user := modul.User{
	// 	Username: "Jasur",
	// 	Email: "schkda@gmail.com",
	// 	Password: "12356",
	// }
	// err = us.Create(user)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Succes Add Users")

	// res,err := us.GetAllProducts()
	// if err != nil{
	// 	panic(err)
	// }

	// for _,r := range res{
	// 	fmt.Println(r)
	// }

	// user := modul.User{
	// 	Id:       16,
	// 	Username: "Ahmad",
	// 	Email:    "schkasda@gmail.com",
	// 	Password: "12345",
	// }
	// err = us.Update(user, user.Id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("ID-%d Update Success", user.Id)

	// id := 16
	// err = us.Delate(id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("ID - %d Users Delate\n", id)
}
