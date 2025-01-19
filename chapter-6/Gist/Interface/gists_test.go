package main_test

import (
	"testing"
	"strings"
	"io"
	. "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-6/Gist/Interface"
)

type mockDoer struct {}

func (m *mockDoer) DoGistRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
		[
			{"html_url": "https://gist.github.com/mattn/1234567890"},
			{"html_url": "https://gist.github.com/mattn/0987654321"}
		]
	`), nil
}

func TestListGists(t *testing.T) {
	client := Client{GistClient: &mockDoer{}}
	urls, err := client.ListGists("mattn")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
	if err != nil {
		t.Fatal(err)
	}
}
