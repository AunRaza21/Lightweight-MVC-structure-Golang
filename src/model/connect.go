package model

// First install this github package using
// go get github.com/mattn/go-sqlite3
import (
   "database/sql"
   "fmt"
   _"github.com/mattn/go-sqlite3"
   "log"
)

var conn *sql.DB

func Connect() *sql.DB {
  database, err := sql.Open("sqlite3", "nraboy.db")
  if err != nil {
    log.Fatal(err)
  }
  conn = database
  fmt.Println("Connected to database")
  return database
}
