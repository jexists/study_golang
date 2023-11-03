package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 웹 핸들러 등록
	// http.HandleFunc("경로",func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// 웹 서버 시작
	http.ListenAndServe(":3000", nil)
	// http://127.0.0.1:3000/
	// http://localhost:3000/
}
