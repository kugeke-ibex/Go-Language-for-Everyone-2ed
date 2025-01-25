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


	result_1, err := db.Exec(`INSERT INTO chapter_07.users (name, age) VALUES ('Bob', 18)`)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Println(result_1)
	fmt.Println(result_1.LastInsertId()) // 最終挿入ID PostgreSQLはサポートしていない
	fmt.Println(result_1.RowsAffected()) // 更新件数

	// プレースホルダーを使用してSQLを実行する
	result_2, err := db.Exec(`INSERT INTO chapter_07.users (name, age) VALUES ($1, $2)`, "Alice", 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result_2)
	fmt.Println(result_2.LastInsertId()) // 最終挿入ID PostgreSQLはサポートしていない
	fmt.Println(result_2.RowsAffected()) // 更新件数
	
	// 名前付きプレースホルダーを使用してSQLを実行する (PostgreSQLはサポートしていない)
	result_3, err := db.Exec(`INSERT INTO chapter_07.users (name, age) VALUES ($name, $age)`, map[string]interface{}{"name": "Charlie", "age": 22})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result_3)
}
