package model
import "fmt"

func CreateTodo(name, todo string) error {
  //query, err := conn.Prepare("CREATE TABLE IF NOT EXISTS TODO (id INTEGER PRIMARY KEY, name TEXT, todo TEXT)")
  //query.Exec()
  query, err := conn.Query("INSERT INTO TODO VALUES($1, $2, $3)", 2,  name, todo)
  fmt.Println(query)
  defer query.Close()
  if err != nil {
    return err
  }
  return nil
}
