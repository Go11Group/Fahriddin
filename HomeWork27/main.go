package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Cars struct {
	Id    string
	Brand string
	Year  int
	Price int
}

type Join struct {
	Name1 string
	Name2 string
	Year  int
	Price int
}

func main() {
	db, err := sql.Open("postgres", "host = localhost port = 5432 user = postgres dbname = nt password = 0412")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 3 - misol

	// cars := Cars{}

	// rows,err := db.Query("Select * from Cars")
	// if err != nil{
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next(){
	// 	err = rows.Scan(&cars.Id,&cars.Brand,&cars.Year,&cars.Price)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// 	fmt.Println(cars)
	// }

	//3.1-misol
	//work.sql da 2 ta yangi table ochilishi tartiibi va ozgaruvchilari ko'rsatilgan

	// InsertQueary := "Insert Into Cars(Brand,Year,Price) Values($1, $2, $3)"

	// cars := []struct {
	// 	brand string
	// 	year  int
	// 	price int
	// }{
	// 	{"Toyota", 2020, 20000},
	// 	{"Honda", 2019, 18000},
	// 	{"Ford", 2021, 25000},
	// 	{"Chevrolet", 2018, 22000},
	// 	{"Nissan", 2020, 21000},
	// 	{"BMW", 2021, 35000},
	// 	{"Mercedes", 2019, 40000},
	// 	{"Volkswagen", 2020, 27000},
	// 	{"Audi", 2021, 33000},
	// 	{"Hyundai", 2019, 19000},
	// }
	// for _,car := range cars {
	// 	_,err = db.Exec(InsertQueary,car.brand,car.year,car.price)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("10 new cars added successfully!")

	// InsertQueary = "Insert into Users(Name, Car_id) Values($1, $2)"

	// Users := []struct {
	// 	Name   string
	// 	Car_id string
	// }{
	// 	{"Jamshid", "daba7966-8de2-4ec5-804e-1e7111b07ebb"},
	// 	{"Ali", "b60c9833-decc-4004-aa1f-06da7f46b450"},
	// 	{"Sarvar", "00053c5b-0b18-430a-a878-fe19fce64129"},
	// 	{"Muhriddin", "6e917376-00e7-4113-bb2f-d09f5253ef17"},
	// 	{"Bekzod", "9be214de-3de1-4b8f-9bd5-1fadff6d9207"},
	// 	{"Nurmuhammad", "01e6a6d6-7ba3-437e-bbfd-c02af8443fd1"},
	// 	{"Sanjarbek", "99f17925-43f3-431e-9d88-cf28aa7f1e37"},
	// 	{"Hamidjon", "3f8278ee-e3cf-472a-9859-9bc0cd41c335"},
	// 	{"Diyorbek", "1145b367-9d1d-4982-b331-2ed6de82f294"},
	// 	{"Fahriddin", "be034c81-863c-4663-9ba7-90460f8c80d3"},
	// }

	// for _, user := range Users {
	// 	_, err = db.Exec(InsertQueary, user.Name, user.Car_id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("10 new users added successfully!")

	// rows, err := db.Query(`Select U.name, C.Brand, C.year, C.price from Users as U 
	// Join Cars as C
	// On C.id = U.car_id`)
	// if err != nil {
	// 	panic(err)
	// }
	// j := Join{}
	// for rows.Next() {
	// 	err = rows.Scan(&j.Name1, &j.Name2, &j.Year, &j.Price)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(j)
	// }

}
