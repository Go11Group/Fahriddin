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

}
