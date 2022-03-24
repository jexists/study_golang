- 

# 정적배열 vs 동적배열

## 정적배열 (static)

→ fixed size array (배열의 길이가 변하지 않는)

→ compile time 

ex) [10]int

## 동적배열 (dynamic)

→ run time

→ 실행도중에 바뀔 수 있다. 

→ 길이가 늘어난다. 

→ 실제 배열을 포인트하고 있다.

### 동적배열

- c++: STL vector
- python: slice
- java: Array List
- C#: List
- go: slice

# 슬라이스

→ Go에서 제공하는 배열을 가리키는 포인터 타입

→ 배열의 요소들을 동적인 크기로 유연하게 사용 (배열: 고정된 크기)

→ `[]T`: T타입을 원소로 가지는 슬라이스

→ `a[처음포함 : 끝제외]`

→ 슬라이스 요소 변경시 기본 배열 해당 요소 수정

```go
// c++
vector<int>
//java
Array List 
//python
slice
//javascript
[]
```

## 슬라이스 선언

```go
var slice []int
slice := []int{}
```

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) //[3, 5, 7]
}
```

→ 데이터 저장X, 기본 배열의 한 영역 나타냄

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
// [John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
```

### 슬라이스 리터럴

```go
//배열 리터럴
[3]bool{true, true, false}

//동일한 배열 생성 -> 참조하는 슬라이스 만들어짐
[]bool{true, true, false}

//slice 바로 만드는 방법
//배열과 똑같이 만든 후 앞에 길이 생략
slice := []int{1, 2, 3, 4}
fmt.Println(slice)
```

```go
package main

import "fmt"

func main() {
	q := [3]bool{true, false, true}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

//[true false true]
//[true false true]
//[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
```

### 슬라이스 기본 값

→ `a[0: 슬라이스 길이]`: 생략 가능

```go
var a [10]int

//동일한 의미
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}
```

## 슬라이스 길이

→ slice length

→ 슬라이스가 포함하는 요소의 개수

→ `len(s)`

→ 슬라이스 길이를 연장하기 위해서는 슬라이스에 충분한 용량이 있을 때 가능

## 슬라이스 용량

→ slice capacity

→ 슬라이스의 첫번째 요소부터 계산하는 기본 배열의 요소의 개수

→ `cap(s)`

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]
```

## nil 슬라이스

→ 슬라이스 zero value: nil

→ 슬라이스 길이 = 0, 용량 = 0,  기본배열 X

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
//[] 0 0
//nil!
```

```go
package main

import "fmt"

func main() {
	var slice []int 
	//slice 기본값이  nill
	//nil 슬라이스의 길이와 용량은 0이며, 기본 배열을 가지고 있지 않습니다.
	
	fmt.Println(slice)        //[] 빈배열
	fmt.Println(slice == nil) // true

	// fmt.Println(slice[0])     //ERROR
}
```

## make

→ 동적 크기의 배열을 생성하는 방법

→ 0으로 이루어진 배열 할당 / 해당 배열 참조하는 슬라이스 반환

→ 용량 생략가능 (같은 값이라고 생각)

```go
a := make([]int, 5)  // len(a)=5

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) //a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) //b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) //c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) //d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
//X _ X
//O _ X
//_ _ O
```

## 슬라이스 요소 추가하기

→ `append`

→ `func append(s []T, vs ...T) []T`

→ append(`s` 는 슬라이스의 타입 `T`,  `T` 값들은 슬라이스에 추가할 값)

→ 결과 값: 원래 슬라이스의 모든 요소와 추가로 제공된 값들을 포함하는 슬라이스

→ 원래 배열이 용량 부족할 경우 새로운 내부 배열 할당 (반환된 슬라이스는 새로 할당된 배열)

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s) //len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) //len=1 cap=1 [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s) //len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s) //len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```go
package main

import "fmt"

func main() {
	f := make([]int, 0, 2)
	//길이:0, 용량 2의 slice 생성 
	fmt.Println(f, len(f), cap(f)) // [] 0 2
	f = append(f, 10) //append 값 추가
	fmt.Println(f, len(f), cap(f)) // [10] 1 2
	f = append(f, 20)
	fmt.Println(f, len(f), cap(f)) // [10 20] 2 2 
	f = append(f, 30)
	//용량 부족 => 새로운 내부 배열 할당
	fmt.Println(f, len(f), cap(f)) // [10 20 30] 3 4
	f = append(f, 40) 
	fmt.Println(f, len(f), cap(f)) // [10 20 30 40] 4 4
	f = append(f, 50)
	fmt.Println(f, len(f), cap(f)) // [10 20 30 40 50] 5 8
	//cap => 새로운 배열 할당해서 값이 2배오름 (타입마다 다름)
}
```

---

# slice

→ slice == struct

→ 3개의 properties로 구성

### Properties of Slice

→ pointer * (시작 주소)

→ len int (길이 / 갯수)

→ cap int (할당된 최대 갯수)

## 배열 만드는 방법

## 1. 빈배열 var a []int

```go
var a []int
	fmt.Printf("len(a) = %d\n", len(a)) //0
	fmt.Printf("cap(a) = %d\n", cap(a)) //0
```

## 2. 초기값 세팅 a := []int{1,2,3,4}

```go
b := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(b) = %d\n", len(b)) //5
	fmt.Printf("cap(b) = %d\n", cap(b)) //5
```

## 3. make 사용 a:= make([]int,3)

```go
a:= make([]int,3)
```

## 4. make 사용 with cap a:= make([]int, 0, 8)

```go
c := make([]int, 0, 8)
	fmt.Printf("len(c) = %d\n", len(c)) //0
	fmt.Printf("cap(c) = %d\n", cap(c)) //8
```

## Append

```go
// make 사용
	c := make([]int, 0, 8)
	fmt.Printf("len(c) = %d\n", len(c)) //0
	fmt.Printf("cap(c) = %d\n", cap(c)) //8

// append 추가
	c = append(c, 1)
	// slice = append(slice, 추가할 항목)
	// 반환값 => slice (cap여유가 있을 경우는 입력한 slice /  cap자리가 없을 경우 새로운 slice)
	fmt.Printf("len(c) = %d\n", len(c)) //1
	fmt.Printf("cap(c) = %d\n", cap(c)) //8
```

### Append - 공간이 없을 경우 새로운 메모리 복사

```go
	d := []int{1, 2}
	d[0] = 1
	d[1] = 2
	e := append(d, 3)
	fmt.Println(d) //[1 2]
	fmt.Println(e) //[1 2 3]
	e[0] = 4
	e[1] = 5
	fmt.Println(d) //[1 2]
	fmt.Println(e) //[4 5 3]

	fmt.Printf("%p %p\n", d, e)
	//0xc0000b4040 0xc0000b8000 -> 포인터 다름

	for i := 0; i < len(d); i++ {
		fmt.Printf("%d, ", e[i]) //1, 2,
	}
	fmt.Println()

	for i := 0; i < len(e); i++ {
		fmt.Printf("%d, ", e[i]) //1, 2, 3,
	}
	fmt.Println()

	fmt.Println(cap(d), " ", cap(e)) //2   4
```

### Append - 공간이 남을 경우 같은 메모리 주소

```go
	f := make([]int, 2, 4)
	f[0] = 1
	f[1] = 2
	g := append(f, 3)
	fmt.Println(f) //[1 2]
	fmt.Println(g) //[1 2 3]
	g[0] = 4
	g[1] = 5
	fmt.Println(f) //[4 5]
	fmt.Println(g) //[4 5 3]

	fmt.Printf("%p %p\n", f, g)
	//0xc000016180 0xc000016180 -> 포인터 같음

	n := []int{1, 2}
	n[0] = 1
	n[1] = 2
	m := make([]int, len(n))

	// 포인터가 같기때문에 복사 후 Append
	for i := 0; i < len(n); i++ {
		m[i] = n[i]
	}
	m = append(m, 3)

	fmt.Printf("%p %p\n", n, m)
	//0xc0000121b0 0xc00001a050
	m[0] = 4
	m[1] = 5
	fmt.Println(n) //[1 2]
	fmt.Println(m) //[4 5 3]

```

## 정리

- 동적배열 -> 길이 늘어남
- 추가 → append (여유없을때 두배 / 같은 메모리반환 또는 다른 메모리 반환할 경우 있음)

# Slice 정리

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8] //[slice 시작하는 index +1:slice 끝나는 index]
	// b := []int{5, 6, 7, 8}

	c := a[4:] //[slice 시작하는 index +1:] 끝나는 것을 생략했을 경우 끝까지 slice

	d := a[:4] //[:slice 끝나는 index] 시작하는 거 생략했을 경우 처음부터 slice

	fmt.Println(b) //[5 6 7 8]
	fmt.Println(c) //[5 6 7 8 9 10]
	fmt.Println(d) //[1 2 3 4]

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := e[4:8]
	f[0] = 1
	f[1] = 2
	fmt.Println(e) //[1 2 3 4 1 2 7 8 9]
	// 일부분을 가져오는 것 (같은 메모리)

	for i := 0; i < 5; i++ {
		// var back int
		var front int
		// a, back = RemoveBack(a)
		a, front = RemoveFront(a)
		// fmt.Printf("%d, ", back) // 10, 9, 8, 7, 6,
		fmt.Printf("%d, ", front) // 1, 2, 3, 4, 5,
	}
	fmt.Println()
	fmt.Println(a) // [1 2 3 4 5] //[6 7 8 9 10]
}

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
```

```go
package main

import "fmt"

func main() {
	var s []int
	var t []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	fmt.Println(s, len(s), cap(s)) //[100 200 300] 3 3

	fmt.Println("//////////////////")

	// s = append(추가하고자 하는 slice, 넣을 요소들)
	t = append(s, 400)
	s = append(s, 400, 500, 600, 700)

	fmt.Println(t, len(t), cap(t)) //[100 200 300 400] 4 6
	fmt.Println(s, len(s), cap(s)) //[100 200 300 400 500 600 700] 7 8

	fmt.Println("//////////////////")
	var u []int
	u = append(t, 500)

	fmt.Println(s, len(s), cap(s)) //[100 200 300 400 500 600 700] 7 8
	fmt.Println(t, len(t), cap(t)) //[100 200 300 400] 4 6
	fmt.Println(u, len(u), cap(u)) //[100 200 300 400 500 600 700] 5 6

	u[0] = 9999

	fmt.Println("//////////////////")

	fmt.Println(s, len(s), cap(s)) //[100 200 300 400 500 600 700] 7 8
	fmt.Println(t, len(t), cap(t)) //[9999 200 300 400] 4 6
	fmt.Println(u, len(u), cap(u)) //[9999 200 300 400 500] 5 6
	// 포인터: 메모리주소 번지 값으로
}
```

주의!

```go
var array = [...]int{1, 2, 3,}
v
```

### 슬라이스

```go
type SliceHeader struct {
	Data uintptr // 실제 배열을 가리키는 포인터
	Len int // 요소 개수
	Cap int // 최대길이 (실제 배열의 길이)
}
//실제 배열은 따로 있다.
```