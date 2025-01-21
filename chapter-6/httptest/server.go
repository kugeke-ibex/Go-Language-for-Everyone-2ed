package main

import (
	"fmt"
	"log"
	"net/http"
)

// RouteはこのAPIサーバーのルーティングを定義する
func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// nameを受け取ってパラメータとして埋め込んで返す
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	})
	return m
}

func main() {
	m := Route()
	log.Println("server is running on port 8080")
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
}
