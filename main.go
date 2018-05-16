package main

import (
	"net/http"
)

func main() {
	err := initializeDB()
	defer db.Close()
	if err != nil {
		return
	}
	initializeRoutes()
	http.ListenAndServe(":5051", router)
}
