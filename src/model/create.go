package model

import (
	"fmt"
	"log"
)

func CreateTodo(name, todo string) error {
	//query, err := conn.Prepare("CREATE TABLE IF NOT EXISTS TODO (id INTEGER PRIMARY KEY, name TEXT, todo TEXT)")
	//query.Exec()
	//str1 := "INSERT INTO TODO VALUES($1, $2, $3)"
	sql := "INSERT INTO TODO(id, name, todo) VALUES (3, 'Moscow', 'Hello')"
	res, err := conn.Exec(sql)
	//fmt.Println(res)
	if err != nil {
		panic(err)
	}
	lastId, err := res.LastInsertId()
	fmt.Println(lastId)

	if err != nil {
		log.Fatal(err)
	}

	//defer res.Close()
	return nil
}
