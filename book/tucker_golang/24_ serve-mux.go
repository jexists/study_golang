package main

import (
	"fmt"
	"net/http"
)

func main() {
	// serveMux 인스턴스 생성
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bar")
	})

	// mux 인스턴스 사용
	http.ListenAndServe(":3000", mux)
}
