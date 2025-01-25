package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, age FROM chapter_07.users ORDER BY name`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]interface{}, len(columns))
	refs := make([]interface{}, len(columns))
	for i := 0; i < len(columns); i++ {
		refs[i] = &values[i]
	}

	fmt.Println(columns)
	fmt.Println(values)
	fmt.Println(refs)

	for rows.Next() {
		rows.Scan(refs...)
		fmt.Println(values...)
	}

	fmt.Println(columns)
	fmt.Println(values)
	fmt.Println(refs)

}
