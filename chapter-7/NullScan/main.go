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
	var name *string // NULLを許容するカラムに対してはポインタ型を宣言する
	var age int64
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	if name != nil {
		fmt.Println(*name)
	}
	fmt.Println(id, *name, age)
}
