package postgres

import (
	"database/sql"
	"mymode/modul"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (p *ProductRepo) Create(Product modul.Product) error {
	tr, err := p.DB.Begin()
	defer tr.Commit()
	if err != nil {
		return err
	}
	_, err = p.DB.Exec("INSERT INTO Products(name,description,price,stock_quantity) VALUES($1,$2,$3,$4)",
		Product.Name, Product.Description, Product.Price, Product.Stock_quantity)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepo) GetAllProducts() ([]modul.Product, error) {
	tr, err := p.DB.Begin()
	defer tr.Commit()

	if err != nil {
		return []modul.Product{}, err
	}
	products := []modul.Product{}
	product := modul.Product{}

	rows, err := p.DB.Query("SELECT * FROM Products")
	if err != nil {
		return []modul.Product{}, err
	}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil {
			return []modul.Product{}, err
		}
		products = append(products,product)
	}
	return products,nil
}

func (p *ProductRepo)GetById(id string) (modul.Product,error){
	tr,err := p.DB.Begin()
	if err != nil{
		return modul.Product{},nil
	}
	defer tr.Commit()

	product := modul.Product{}

	row := p.DB.QueryRow("Select * from Products Where id = $1",id)
	
	err = row.Scan(&product.Id,&product.Name,&product.Description,&product.Price,&product.Stock_quantity)
	if err != nil{
		return modul.Product{},nil
	}
	return product,nil
}

func (p *ProductRepo) Update(Product modul.Product,id int) error{
	tr,err := p.DB.Begin()
	defer tr.Commit()

	if err != nil{
		return nil
	}
	_,err = p.DB.Exec(`Update Products Set
	name = $1,
	description = $2,
	price = $3,
	stock_quantity = $4  Where id = $5`,
	Product.Name,Product.Description,Product.Price,Product.Stock_quantity,id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepo) Delate(id int) error{
	tr,err := p.DB.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()
	
	_,err = p.DB.Exec("DELETE FROM Products WHERE id = $1",id)
	if err != nil{
		return err
	}
	return nil
}