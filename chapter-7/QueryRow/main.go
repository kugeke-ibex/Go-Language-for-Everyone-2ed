package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main () {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow(`SELECT id, name, age FROM chapter_07.users WHERE id = $1`, 1)

	var id int64
	var name string
	var age int64
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, name, age)
}
