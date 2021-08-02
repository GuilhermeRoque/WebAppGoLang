package models

import (
	"goWebApp/db"
)
type Product struct {
	Id int
	Name string
	Description string
	Price float32
	Amount int
}

func GetAllProducts() []Product {
	dbConn := db.ConnectDB()
	lines, err := dbConn.Query("select * from products")
	p := Product{}
	var products []Product
	for lines.Next(){
		var id,amount int
		var name,description string
		var price float32
		err = lines.Scan(&id,&name,&description,&price,&amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Amount = amount
		p.Price = price
		p.Amount = amount
		p.Description = description
		products = append(products, p)
	}

	defer dbConn.Close()
	return products
}

func CreateNewProduct(product Product){
	dbConn := db.ConnectDB()
	prepare, err := dbConn.Prepare("insert into Products (name,description,price,amount) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(product.Name,product.Description,product.Price,product.Amount)

	defer dbConn.Close()

}

func UpdateProduct(product Product){
	dbConn := db.ConnectDB()
	prepare, err := dbConn.Prepare("update Products set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(product.Name,product.Description,product.Price,product.Amount)
	defer dbConn.Close()

}

func DeleteProduct(id string) {
	dbConn := db.ConnectDB()
	prepare, err := dbConn.Prepare("delete from Products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(id)

	defer dbConn.Close()

}

func GetProduct(id string) Product {
	dbConn := db.ConnectDB()
	lines, err := dbConn.Query("select * from products where id=$1",id)	//

	p :=Product{}

	for lines.Next(){
		var id,amount int
		var name,description string
		var price float32
		err = lines.Scan(&id,&name,&description,&price,&amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Amount = amount
		p.Price = price
		p.Amount = amount
		p.Description = description
	}

	defer dbConn.Close()

	return p
}
