package main

import (
	"database/sql"
	"log"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

type User struct {
	ID int	`db:"id"`
	Name string `db:"name"`
	Age int `db:"age"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM chapter_07.users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Printf("%+v\n", users)
	
	// Pingでの接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
