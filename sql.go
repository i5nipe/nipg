package main

import (
	"database/sql"
	"fmt"
)

type Users struct {
	Id       int
	User     string
	Password string
	Role     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "nipg"
	dbPass := "root"
	dbName := "nipg"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println("Falha ao se conectar ao Banco de dados")
		panic(err.Error())
	}
	return db
}
