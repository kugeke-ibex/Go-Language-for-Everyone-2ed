package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp/v3"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbMap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.PostgresDialect{},
	}

	// gorpによるテーブル作成
	table := dbMap.AddTableWithName(Comment{}, "comments")
	table.SetKeys(true, "ID")
	table.AddIndex("idx_comments", "Btree", []string{"text"}).
		SetUnique(true)

	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	// データの挿入
	comment := Comment{
		Name: "",
		Text: "こんにちわ",
	}
	err = dbMap.Insert(&comment)
	if err != nil {
		log.Fatal(err)
	}

	comment2 := Comment{
		Name: "Alice",
		Text: "こんにちわ",
	}
	// データの挿入2
	err = dbMap.Insert(&comment2)
	if err != nil {
		log.Fatal(err)
	}

	// 1件抽出
	var selectComment Comment
	err = dbMap.SelectOne(&selectComment, "SELECT * FROM comments WHERE id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selectComment: %+v\n", selectComment)

	// データの更新
	selectComment.Text = "こんばんは"
	_, err = dbMap.Update(&selectComment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updateComment: %+v\n\n", selectComment)

	// 複数の抽出
	var comments []Comment
	_, err = dbMap.Select(&comments, "SELECT * FROM comments")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("comments: %+v\n\n", comments)
}
