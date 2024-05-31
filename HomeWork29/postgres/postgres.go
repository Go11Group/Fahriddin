package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM() (*gorm.DB,error ){
	db, err := gorm.Open(postgres.Open("postgres://postgres:0412@localhost:5432/nt?sslmode=disable"))
	if err != nil{
		return &gorm.DB{},nil
	}

	return db,nil
}
