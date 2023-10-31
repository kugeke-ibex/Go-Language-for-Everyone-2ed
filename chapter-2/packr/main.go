package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	f, err := packr.NewBox("./public").Open("index.html")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)

	box := packr.NewBox("./templates")

	s, err := box.FindString("admin/index.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
