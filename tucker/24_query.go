package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수
	name := values.Get("name") // 특정 키값 있는지 확인
	if name == "" {
		name = "World"
	}
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		// 에러처리
	}
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}
func main() {
	http.HandleFunc("/bar", barHandler)

	// 웹 서버 시작
	http.ListenAndServe(":3000", nil)
	// http://localhost:3000/bar = Hello World! id:0
	// http://localhost:3000/bar?name=joy&id=7 = Hello joy! id:7
	// http://localhost:3000/ = 404 page not found

}
