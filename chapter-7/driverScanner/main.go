package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/lib/pq"
)

type myType string

func (mt *myType) Scan(src interface{}) error {
	switch x := src.(type) {
	case int64:
		*mt = myType(fmt.Sprint(src))
	case float64:
		*mt = myType(fmt.Sprintf("%.3f", src))
	case bool:
		*mt = myType(fmt.Sprint(src))
	case []byte:
		if len(x) < 10 {
			*mt = myType(fmt.Sprintf("[% 02X]", x))
		} else {
			x = x[:10]
			*mt = myType(fmt.Sprintf("[% 02X...]", x))
		}
	case string:
		*mt = myType(x)
	case time.Time:
		*mt = myType(x.Format("2004/01/02 15:04:05"))
	case nil:
		*mt = "nil"

	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return nil
}
func main () {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, age FROM chapter_07.users;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int64
	var name string
	var age myType  // 文字列型をベースにしたmyType型にScanメソッドを実装し、各型から文字列に変換
	for rows.Next() {
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("value: %d %s %s\n", id, name, age)
		fmt.Printf("type: %T %T %T\n", id, name, age)
		age.Scan(1.234567890) // メソッドを呼び出すと、ageの値が変わるが型は文字列のまま
		fmt.Printf("value: %d %s %s\n", id, name, age)
		fmt.Printf("type: %T %T %T\n", id, name, age)
	}
}

