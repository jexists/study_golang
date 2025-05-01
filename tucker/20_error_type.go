package main

import "fmt"

type PasswordError struct {
	Len       int
	RequreLen int
}

func (err PasswordError) Error() string {
	// 구조체 메서드로 Error()함수 선언 ->  error 인터페이스 사용가능
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len: %d RequireLen %d \n", errInfo, errInfo.Len, errInfo.RequreLen)
			// 암호 길이가 짧습니다. Len: 4 RequireLen 8
		} else {
			fmt.Println("회원가입 완료")
		}
	}
}
