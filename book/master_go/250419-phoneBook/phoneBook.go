package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}
func main() {
	// go build phoneBook.go
	arguments := os.Args
	if len(arguments) == 1 {
		// -> 인자가 하나도 없는 경우
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"joy", "sk", "010"})
	data = append(data, Entry{"happy", "kt", "011"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
		// ./phoneBook search kt
		// {happy kt 011}
	case "list":
		list()
		// ./phoneBook list
		// {joy sk 010}
		// {happy kt 011}
	default:
		fmt.Println("Not a valid option")
	}

}

// os.Args
// -> 프로그램 실행 시 전달되는 명령줄 인자들을 담고 있는 문자열 슬라이스
// os.Args[0]: 실행한 프로그램 자체의 경로
// os.Args[1]~: 실행시 입력한 추가 인자들
