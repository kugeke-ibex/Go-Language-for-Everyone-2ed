package main

import (
	"fmt"
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"log"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

// DoGistRequestはグローバルスコープに定義された手続きのオブジェクト
var DoGistRequest = func(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	return &buf, nil
}

func ListGists(user string) ([]string, error) {
	r, err := DoGistRequest(user)
	if err != nil {
		return nil, err
	}
	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}
	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}

func main() {
	urls, err := ListGists("mattn")
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}
