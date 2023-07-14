package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("valueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("valueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop3() {
	f := make([]func(), 3)
	fmt.Println("valueLoop3")
	i := 0
	for ; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for j := 0; j < 3; j++ {
		f[j]()
	}
}
func main() {
	CaptureLoop()
	CaptureLoop2()
	CaptureLoop3()
	// valueLoop
	// 3
	// 3
	// 3
	// valueLoop2
	// 0
	// 1
	// 2
	// valueLoop3
	// 3
	// 3
	// 3
}
