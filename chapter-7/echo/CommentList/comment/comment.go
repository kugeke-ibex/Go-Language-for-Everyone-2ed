package comment

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp/v3"
	_ "github.com/lib/pq"
)

type Comment struct {
	ID        int64     `json:"id" db:"id,primarykey,autoincrement"`
	Name      string    `json:"name" db:"name,notnull,size:200" validate:"max=200"`
	Text      string    `json:"text" db:"text,notnull,size:400" validate:"required,max=400"`
	CreatedAt time.Time `json:"created_at" db:"created_at,notnull"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at,notnull"`
}

func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	if c.Name == "" {
		c.Name = "名無し"
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return nil
}

func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.UpdatedAt = time.Now()
	return nil
}

func NewGorpDbMap() *gorp.DbMap {
	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	dbMap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.PostgresDialect{},
	}

	dbMap.AddTableWithName(Comment{}, "comments")
	return dbMap
}

func NewComment(name, text string) *Comment {
	return &Comment{Name: name, Text: text}
}

