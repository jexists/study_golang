package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

// func (f *File) Close() {
// Reader 인터페이스에 Close() 메소드가 없으면 에러 생김
// c := reader.(Closer)
// c.Close()
// }

func ReadFile(reader Reader) {
	// 서로 다른 인터페이스로 변환가능 -> 실행중에 에러가 발생할수있음
	// c := reader.(Closer)
	// c.Close()
	// Reader 인터페이스에 Close() 메소드가 없으면 에러 생김

	// 타입변환 성공 여부 반환
	// 타입 변환한 결과 , 변환 성공 여부
	c, ok := reader.(Closer)
	fmt.Println(c, ok)
	// &{} true: Reader 인터페이스에 Close() 메소드 있는 경우
	// <nil> false: Reader 인터페이스에 Close() 메소드 없는 경우
	if ok {
		c.Close()
	}

	// if문 초기화를 사용해서 한줄로 간단하게 사용
	if c, ok := reader.(Closer); ok {
		fmt.Println(c, ok) // &{} true
		c.Close()
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}
