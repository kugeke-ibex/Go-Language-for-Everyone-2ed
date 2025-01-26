package main

import (
	"net/http"
	"fmt"

	. "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-7/echo/CommentList/comment"
	. "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-7/echo/CommentList/validator"
	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()
	e.Validator = NewValidator()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// コメント一覧を取得
	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment

		dbMap := NewGorpDbMap()
		defer dbMap.Db.Close()
		_, err := dbMap.Select(&comments, "SELECT * FROM comments ORDER BY created_at desc LIMIT 10")
		if err != nil {
			c.Logger().Error("Select:", err)
			return c.String(http.StatusBadRequest, "Select: "+err.Error())
		}

		return c.JSON(http.StatusOK, comments)
	})

	// コメントの登録
	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			c.Logger().Error("Bind:", err)
			return c.String(http.StatusBadRequest, "Bind: "+err.Error())
		}

		if err := c.Validate(comment); err != nil {

			c.Logger().Error("Validate:", err)
			return c.String(http.StatusBadRequest, "Validate: "+err.Error())
		}

		dbMap := NewGorpDbMap()
		defer dbMap.Db.Close()

		if err := dbMap.Insert(&comment); err != nil {
			c.Logger().Error("Insert:", err)
			return c.String(http.StatusBadRequest, "Insert: "+err.Error())
		}

		c.Logger().Infof("ADDED: %v", comment.ID)
		fmt.Println("OK")
		return c.String(http.StatusCreated, "")
	})

	// 静的ファイルの配信
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
