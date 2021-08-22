package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	//mongodb is the name of a docker container from a mongo image
	conection := "user=postgres dbname=web_store password=example host=postgresdb sslmode=disable"
	db, err := sql.Open("postgres",conection)
	if err != nil{
		panic(err.Error())
	}
	return db
}
