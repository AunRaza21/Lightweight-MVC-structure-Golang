package main

import (
  "controller"
  "net/http"
  "model"
)

func main() {

  mux := controller.Register()
  db := model.Connect()
  defer db.Close()  // It will run this line at the end of all executions.
  http.ListenAndServe("localhost:8080", mux)
}
