package main

import (
	"github.com/amrali/FOODOFFERS/pkg/db"
	"github.com/amrali/FOODOFFERS/pkg/router"
)

func main() {
	dataBase := db.ConnectDB()
	db.MigrateDB(dataBase)
	//db.AssignForeignKeys(dataBase)
	defer dataBase.Close()
	router.HttpRequests()
}
