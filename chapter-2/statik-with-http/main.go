package main

// マジックコメント
//go:generate statik

import (
	"net/http"

	_ "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-2/use-statik/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, _ := fs.New()
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
