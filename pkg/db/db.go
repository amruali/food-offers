package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host       = "localhost"
	port       = 5433
	user       = "postgres"
	password   = "a12345"
	dbname     = "DBDEV"
	searchPath = "FOODOFFS"
)


func ConnectDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, searchPath)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.Exec(`set search_path='FOODOFFS'`)
	return db
}
