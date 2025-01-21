package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	// ルーティングの設定はそのままにテスト用のサーバーを立ち上げることができる
	// httptest.Serverが便利なのはhttp.Handlerインターフェースを引数に渡せる
	// → 今回だと func Routeが *http.ServeMuxを返すので、これはhttp.Handlerインターフェースを満たしている
	ts := httptest.NewServer(Route())
	defer ts.Close()

	// 通常通りHTTPリクエストを送ることができる
	// テスト用サーバーのURLはts.URLで取得できる
	res, err := http.Get(ts.URL + "/greet?name=gopher")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("read from HTTP Response Body failed: %s", err)
	}

	expected := "Hello, gopher!"
	if string(greeting) != expected {
		t.Fatalf("response of /greet?name=gopher: %s want %s", string(greeting), expected)
	}
}

/**
	http.HandlerインターフェースはServeHTTPメソッドを持つインターフェース
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
