package main

import (
	"io"
	"log"
	"os"

	_ "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-2/use-statik/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	f, err := statikFS.Open("/index.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	io.Copy(os.Stdout, f)
}
