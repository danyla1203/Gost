package db

import (
	"database/sql"
	"log"
)

type Database struct {
	connection *sql.DB
}

func connect(dbms, user, password, dbName string) Database {
	connection, err := sql.Open(dbms, user+":"+password+"@/"+dbName)
	if err != nil {
		log.Fatal("Error while connect to db")
	}
	db := Database{connection: connection}
	return db
}
