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

	row_1 := db.QueryRow(`SELECT id, name, age FROM chapter_07.users WHERE id = $1`, 1)

	var id_1 int64
	var name_1 string
	var age_1 float64  // 整数型の値をfloat64で取得可能
	err = row_1.Scan(&id_1, &name_1, &age_1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id_1, name_1, age_1)

	// row_2 := db.QueryRow(`SELECT id, name, age FROM chapter_07.users WHERE id = $1`, 2)
	// var id_2 int64
	// var name_2 int64  // 文字列型のint64型で宣言するエラーが発生
	// var age_2 float64
	// err = row_2.Scan(&id_2, &name_2, &age_2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(id_2, name_2, age_2)

	row_3 := db.QueryRow(`SELECT id, name, age FROM chapter_07.users WHERE id = $1`, 3)
	var id_3 int64
	var name_3 string
	var age_3 interface{} // どのような値でScanするかわからない場合はinterface{}で宣言する
	err = row_3.Scan(&id_3, &name_3, &age_3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id_3, name_3, age_3)
	n := age_3.(int64) // 型アサーションを使用してint64型に変換
	fmt.Println(n)
}
