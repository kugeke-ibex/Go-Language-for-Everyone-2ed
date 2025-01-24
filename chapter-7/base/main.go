package main

import (
	"database/sql"
	"log"
	"fmt"
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

	rows, err := db.Query("SELECT * FROM users")
	fmt.Println(rows)
}
