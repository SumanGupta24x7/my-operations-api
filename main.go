package main

import (
	"fmt"
	"my-operations/database"
	"my-operations/router"
	"net/http"
)

func main() {
	db, err := database.InitDbConnection()

	if err != nil {
		fmt.Println(err)
	}

	engine := router.Init(db)

	err = http.ListenAndServe(":8080", engine)
	if err != nil {
		fmt.Println("Unable to start sever, error: ", err)
	}
}
