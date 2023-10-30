package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// httpリクエストは論理パスなのでpathパッケージを使う。
		if ok, err := path.Match("/data/*.txt", r.URL.Path); err != nil || !ok {
			fmt.Println(ok)
			http.NotFound(w, r)
			return
		}

		// 以降はパスを物理パスとして扱うのでpath/filepathパッケージを使う。
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))

		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		defer f.Close()
		io.Copy(w, f)
	})

	http.ListenAndServe(":8080", nil)
}

// アップロードされたファイルをフォルダ内に格納するコード
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, err := r.FormFile("file")
		s, _ := header.Open()
		p := filepath.Join("files", header.Filename)
		buf, _ := ioutil.ReadAll(s)
		ioutil.WriteFile(p, buf, 0644)
		http.Redirect(w, r, "/"+path, 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

// filepath.Baseを利用したファイル名の取得
func uploadUseBase(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stream, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		p := filepath.Join("files", filepath.Base(header.Filename))
		println(p)
		f, err := os.Create(p)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, stream)
		http.Redirect(w, r, path.Join("/files", p), 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
