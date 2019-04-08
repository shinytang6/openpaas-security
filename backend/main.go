package main

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8001")
}