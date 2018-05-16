package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := initializeDB()
	defer db.Close()
	if err != nil {
		fmt.Println("DB NE CONNECT")
		return
	}
	initializeRoutes()
	http.ListenAndServe(":5051", router)
}
