# range

→ for에서 range는 슬라이스 또는 맵의 요소 순회

→ 슬라이스에서 range사용하면 각 순회마다 1.인덱스. 2. 해당 인덱스 값 복사본

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
//2**0 = 1
//2**1 = 2
//2**2 = 4
//2**3 = 8
//2**4 = 16
//2**5 = 32
//2**6 = 64
//2**7 = 128
```

### range continued

→ _ 인덱스 또는 값 생략 가능

→ 인덱스만 원할경우 두번째 변수 생략

```go
for i, _ := range pow
for _, value := range pow

for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

//1
//2
//4
//8
//16
//32
//64
//128
//256
//512
```