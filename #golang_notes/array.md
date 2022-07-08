# 배열

# arrays

→ `[number]Type`: Type n개 원소 배열

→ 배열의 크기를 조정불가

```go
package main

import "fmt"

func main() {
	var a [10]int //값 할당X -> zero value
	fmt.Println(a) //[0 0 0 0 0 0 0 0 0 0] 
}
```

배열 바로 만드는 방법

```go
package main

import "fmt"

func main() {
	b := [3]int{1, 2}
	fmt.Println(b, len(b))
}
```

```go
package main

import "fmt"

func main() {
	d := [3]string{"hello", "world"}
	fmt.Println(d, d[2], len(d))
	fmt.Println(d[2] == "") //true
}
//[hello world ]  3
//true
```

### [...]

→ 길이 생략 (빈 길이 만들수 없음)

→ 길이 명시 안할수 있음 (생략)

```go
package main

import "fmt"

func main() {
	c := [...]int{1, 2}
	fmt.Println(c[0], c[1], len(c))
}
//1 2 2
```

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
// Hello World
// [Hello World]
// [2 3 5 7 11 13]
```

→ var 배열이름 [원소개수]타입 (초깃값 0)

→ 배열이름 := [원소갯수]타입{원소1, 원소2 ... }

```go
var 
```

### 배열 복사

```go
package main

import "fmt"

func main() {
	//배열 복사
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone) //[1 2 3 4 5]
}
```

### 배열 역순

```go
package main

import "fmt"

func main() {
	//배열 역순 (배열 2개 만들어서 복사: 12번 대입 -> 2n)
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}

	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
	fmt.Println("temp", temp) //[5 4 3 2 1]
	fmt.Println("arr", arr)   //[5 4 3 2 1]
}
```

```go
package main

import "fmt"

func main() {
	//배열 역순 (서로 변경 -> n/2)
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("arr", arr) //[5 4 3 2 1]
}
```

### 정렬(sort): RADIX

→ 종류: bubble, merge, insert, heap, radix etc

### RADIX 정렬 (N개)

→ 모든 경우에 사용X

→ 원소의 값의 범위 한정, 범위 작아야함



```go
package main

import "fmt"

func main() {

	// array
	a := [3]int{1, 2}
	// a[4] = 4 // ERROR (크기제한)
	fmt.Println(a, len(a))

	// slice: length 없음
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Println(s, len(s))
}
```