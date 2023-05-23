package main

import (
	"fmt"
	"net/http"
)

func main() {
	// / へのリクエストを受け各種処理を行う
	// http.ResponseWriter: レスポンスを書き戻すためのインターフェース
	// http.Request: handlerが処理するHTTPリクエスト
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "%s\n", r)
	})

	// HTTPサーバの立ち上げて80番ポートで待ち受け
	http.ListenAndServe(":80", nil)
}
