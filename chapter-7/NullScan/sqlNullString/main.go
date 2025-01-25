package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"encoding/json"

	_ "github.com/lib/pq"
)

// JSON出力にはNullStringに対応していないので、別の方を作り、専用のMarshalJSONとUnmarchalJSONを実装する必要がある。
type MyNullString struct {
	s sql.NullString
}

func (s *MyNullString) Scan(src interface{}) error {
	return s.s.Scan(src)
}

func (s MyNullString) String() string {
	if !s.s.Valid {
		return ""
	}
	return s.s.String
}

func (s MyNullString) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.s.String)
}

func (s *MyNullString) UnmarshalJSON(data []byte) error {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		s.s.String = ""
		s.s.Valid = false
		return err
	}
	s.s.String = ss
	s.s.Valid = true
	return nil
}

type User struct {
	ID int64 `json:"id"`
	Name MyNullString `json:"name"` // sql.NullStringに対応した型
	Age int64 `json:"age"`
}

func main () {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 変数がnilかどうかのチェック、*stringとして宣言するのはわずわらしいので、sql.NullString型を使用する
	row := db.QueryRow(`SELECT name FROM chapter_07.users WHERE id = $1`, 1)

	var name sql.NullString
	err = row.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sql.NullString: %s\n\n", name.String) // NULLの場合は空文字列が返る

	rows, err := db.Query(`SELECT id, name, age FROM chapter_07.users`)
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

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("jsonUsers: %s\n\n", string(jsonUsers))

	decodedUsers := []User{}
	err = json.Unmarshal(jsonUsers, &decodedUsers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decodedUsers: %+v\n\n", decodedUsers)
}
