package main

import (
	"controller"
	"fmt"
	"model"
	"net/http"
)

func main() {

	mux := controller.Register()
	db := model.Connect()
	defer db.Close() // It will run this line at the end of all executions.
	fmt.Println("Server is listening on Port: 8090")
	http.ListenAndServe("localhost:8090", mux)
}
