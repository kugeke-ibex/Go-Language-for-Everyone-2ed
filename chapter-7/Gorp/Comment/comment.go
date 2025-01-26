package comment

import (
	"time"
	"github.com/go-gorp/gorp/v3"
)

type Comment struct {
	ID int64 `db:"id,primarykey,autoincrement"`
	Name string `db:"name,notnull,default:'名無し',size:200"`
	Text string `db:"text,notnull,size:400"`
	CreatedAt time.Time `db:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `db:"updated_at,notnull,default:current_timestamp"`
}

func (c *Comment) BeforeInsert(s gorp.SqlExecutor) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Comment) BeforeUpdate(s gorp.SqlExecutor) error {
	c.UpdatedAt = time.Now()
	return nil
}
