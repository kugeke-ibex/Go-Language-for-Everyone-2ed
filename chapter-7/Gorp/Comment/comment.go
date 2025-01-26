package main

import (
	"time"

	"github.com/go-gorp/gorp/v3"
)

// ORMではstructのフィールドにタグを付ける
// Beego orm:
// gorm gorm:
// xorm xorm:
// gorp db:

type Comment struct {
	ID int64 `db:"id,primarykey,autoincrement"`
	Name string `db:"name,notnull,size:200"`
	Text string `db:"text,notnull,size:400"`
	CreatedAt time.Time `db:"created_at,notnull"`
	UpdatedAt time.Time `db:"updated_at,notnull"`
}

// フック関数(データベースに挿入する前に実行される)
func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	// デフォルト値を設定
	if c.Name == "" {
		c.Name = "名無し"
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return nil
}

// フック関数(データベースに更新する前に実行される)
func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.UpdatedAt = time.Now()
	return nil
}
