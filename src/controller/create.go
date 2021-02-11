package controller

import (
  "net/http"
  "encoding/json"
  "views"
  "model"
  "fmt"
)

func create() http.HandlerFunc  {
  return func (w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodPost {
      data := views.PostRequest {}
      json.NewDecoder(r.Body).Decode(&data)
      fmt.Println(data)
      model.CreateTodo(data.Name, data.Todo)
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(data)
      //json.NewEncoder(w).Encode(data)
    }
  }
}
