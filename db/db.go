package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	conection := "user=postgres dbname=web_store password=Lock1net host=localhost sslmode=disable"
	db, err := sql.Open("postgres",conection)
	if err != nil{
		panic(err.Error())
	}
	return db
}
