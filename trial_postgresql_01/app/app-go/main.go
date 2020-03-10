package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type USER struct {
	id   string
	name string
	age  string
}

func main() {
	db, err := sql.Open(
		"postgres",
		"host=127.0.0.1 port=5555 user=example1 password=example1 dbname=example1 sslmode=disable",
	)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// INSERT
	/*
		var id string
		name := "strawberry.chocomint"
		age := 96
		err = db.QueryRow("INSERT INTO USERS(name, age) VALUES($1,$2)", name, age).Scan(&id)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id)
	*/

	// SELECT
	rows, err := db.Query("SELECT * FROM USERS")

	if err != nil {
		fmt.Println(err)
	}

	var us []USER
	for rows.Next() {
		var u USER
		rows.Scan(&u.id, &u.name, &u.age)
		us = append(us, u)
	}
	fmt.Printf("%v", us)
}
