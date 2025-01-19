package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

// Doer は GistsのAPIにリクエストするインターフェース
type Doer interface {
	DoGistRequest(user string) (io.Reader, error)
}

/// Client は GistのList APIを扱うためのクライアント実装
type Client struct {
	GistClient Doer
}

type GistClient struct {}

func (g *GistClient) DoGistRequest(user string) (io.Reader, error) {
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

func (c *Client) ListGists(user string) ([]string, error) {
	r, err := c.GistClient.DoGistRequest(user)
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
	client := Client{GistClient: &GistClient{}}
	urls, err := client.ListGists("mattn")
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}
