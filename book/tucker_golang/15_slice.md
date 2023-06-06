
## 슬라이스 (slice)
- `var 변수명 []타입` 선언
- `변수명 := []타입{}` 선언
- `var slice []int`
- 동적 배열: 자동으로 배열 크기를 증가 시키는 자료구조
- 동적은 사이즈를 안적음
- 길이가 요소 개수에 따라 자동으로 증가해 관리 편리
- 동적으로 길이 늘어나는 점만 제외하면 배열과 같음
- SliceHeader {Data: unitptr, Len(요소개수): int, Cap(실제 배열의 길이) int}
- 배열 가리키는 포인터, 요소개수(Len), 전체 배열 길이(Cap)으로 구성된 구조체
- 남은 빈 공간 = cap - len
- 남은 공간의 개수가 추가하는 값개수보다 크거나 같은 경우 배열 뒷부분에 값을 추가한 뒤 len값 증가
- 빈공간이 없을 때 (cap) 새로운 기존 배열의 2배 크기로 기존 배열의 요소와 새로운 배열 복사
- cap: 새로운 배열의 길이 값(기존cap*2), len 기존 길이에 추가한 개수, 포인터 새로운 배열
- 슬라이스 내부에 배열을 가리키는 포인터
- append(): 배열의 빈공간이 있을 경우 기존 배열에 추가 / 없을 경우 새로운 기존 배열의 2배 크기 생성 후 추가
- 배열의 이루를 나타내는 타입 
- 실제 배열은 따로 있고 그 배열을 가리키는 형태
- **Go에서 제공하는 배열을 가리키는 포인터 타입**
- 배열에 포인터 타입인데 길이 변경가능 
- `slice == *[num]type` -> num 변경 가능
- slice size: 24byte (복사할경우 배열의 길이에 상관없이 똑같은 사이즈 복사: 배열보다 사이즈가 작다)

// string / slice는 복사해도 괜찮다


```go
type SliceHeader struct {
	Data uintptr // 실제 배열을 가리키는 포인터
	Len int // 요소 개수
	Cap int // 실제 배열의 길이 / 최대길이 (capacity)
}


```

#### 동적 / 정적
- 정적: static (compile time) // const (실행도중에 변경X)
- 동적: dynamic (runtime) // var (실행도중에 변경가능)

- 동적 배열
```
// C++
vector<int>
// Java
Array List
// Python 
slice
// Javascript
[]

```


### 배열
- 처음 배열때 정한 길이에서 더 이상 늘어나지 않음
```go
package main
import "fmt"
func main() {
	var array [10]int
	fmt.Println(array)
	// [0 0 0 0 0 0 0 0 0 0]
	// 슬라이스일 경우 []대괄호 안에 길이 넣지 말아야함
	var slice []int
	if len(slice) == 0 {
		fmt.Println("slice empty ", slice)
		// slice empty  []
	}
	slice[1] = 10 // ERROR
	panic: runtime error: index out of range [1] with length 0
}
```
### 배열 vs 슬라이스
```go 
	var arrayType = [...]int{1, 2, 3} // 배열: 고정길이
	// [...]int{1, 2, 3} == [3]int{1,2,3}
	var sliceType = []int{1, 2, 3}    // 슬라이스
	// 동적은 사이즈를 안적음

	fmt.Println(arrayType)
	fmt.Printf("type of arrayType is %T\n", arrayType)
	// [1 2 3]
	// type of arrayType is [3]int
	fmt.Println(sliceType)
	fmt.Printf("type of sliceType is %T\n", sliceType)
	// [1 2 3]
	// type of sliceType is []int
```

### 슬라이스 초기화
- {}사용해서 초기화
```go
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice1)
	// [1 2 3]
	fmt.Println(slice2)
	// [1 0 0 0 0 2 0 0 0 0 3]
```

### make() 사용해서 초기화
``` go
	var sliceMake = make([]int, 3)
	// make([]int, 3) == []int{0, 0, 0}
	fmt.Println(sliceMake) // [0 0 0]
	fmt.Println(len(sliceMake)) // 3 (요소의 개수)
	fmt.Println(cap(sliceMake)) // 3 (실제 배열의 길이)
	fmt.Printf("type of sliceMake is %T\n", sliceMake)
	// type of sliceMake is []int

	var sliceMakeWithCap = make([]int, 3, 5)
	fmt.Println(sliceMakeWithCap) // [0 0 0]
	fmt.Println(len(sliceMakeWithCap)) // 3
	fmt.Println(cap(sliceMakeWithCap)) // 5
```

### 슬라이스 순회
```go
var slice  = []int{1,2,3}

for i := 0; i < len(slice); i++ {
	fmt.Println(slice[i])
}

for i, v := range slice {
	fmt.Println("인덱스", i, " / 값", v)
}
```



### append()
- 슬라이스만의 기능
- 요소 추가
- append(추가하고자 하는 슬라이스, 추가하는 요소) = 새로운 슬라이스 결과 반환
- 기존 슬라이스는 변경X (기존을 새로운 슬라이스로 대입 필요)
- 첫번째 인수로 넣은 값을 ㅕㄴ경하는게 아니라 추가된 새로운 슬라이스 반환
- 기존 슬라이스에 추가하 싶을 경우 대입해서 변경
- 남은 빈 공간 = cap - len
- 남은 공간의 개수가 추가하는 값개수보다 크거나 같은 경우 배열 뒷부분에 값을 추가한 뒤 len값 증가
- 빈공간이 없을 때 (cap) 새로운 기존 배열의 2배 크기로 기존 배열의 요소와 새로운 배열 복사
- cap: 새로운 배열의 길이 값(기존cap*2), len 기존 길이에 추가한 개수, 포인터 새로운 배열
- 슬라이스 내부에 배열을 가리키는 포인터
- append(): 배열의 빈공간이 있을 경우 기존 배열에 추가 / 없을 경우 새로운 기존 배열의 2배 크기 생성 후 추가
- 슬라이스에 요소를 추가한 새로운 슬라이스 반환
- 기존 슬라이스가 바뀔수도 있고 아닐수 있음
- 복사 -> 추가 -> 새로운 struct

### 슬라이싱 Slicing
- 배열의 일부를 집어내는 기능: 슬라이스 반환
- 배열의 일부를 가리키는 기능
- array[startIndex:endIndex]
- 배열의 일부 = 배열[시작인덱스:끝인덱스]
- 배열의 시작인덱스 부터 끝인데스-1 일부 슬라이스 반환
- 새로운 배열이 만들어지는게 아니라 포인터로 가리키는 슬라이스 생성
- 인덱스 2개 사용시 cap: 전체 길이에서 시작 인덱스를 뺀 값 
- 슬라이싱의 결과가 슬라이스다.
- array[startIndex:endindex]
- array[startIndex:endindex:maxindex]
- len == 끝인덱스 - 시작인덱스
- cap == 최대인덱스 - 시작인덱스
- startIndex 생략 == 0
- endindex 생략 == len(slice)
- maxindex 생략 == 최대인덱스 (cap사이즈 조절)

- 처음부터 슬라이싱: 시작인덱스 생략가능
```go
	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := slice4[0:3]
	slice6 := slice4[:3]
	// slice4:  [1 2 3 4 5] 5 5
	// slice5:  [1 2 3] 3 3
	// slice6:  [1 2 3] 3 3
```
- 끝까지 슬라이싱: 끝인덱스 생략가능
```go
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := slice4[2:len(slice7)]
	slice9 := slice4[2:]
	// slice7:  [1 2 3 4 5] 5 5
	// slice8:  [3 4 5] 3 3
	// slice9:  [3 4 5] 3 3
```

- 전체 슬라이싱: 시작인덱스, 끝인덱스 생략가능 (배열 전체를 가리키는 슬라이스 생성)
- 배열을 슬라이스로 바꾸고 싶을때 사용
```go
	slice10 := []int{1, 2, 3, 4, 5}
	slice12 := slice10[0:len(slice10)]
	slice11 := slice10[:]
	// slice10:  [1 2 3 4 5] 5 5
	// slice11:  [1 2 3 4 5] 5 5
	// slice12:  [1 2 3 4 5] 5 5
```
- cap크기 조절
- slice[시작인덱스:끝인덱스:최대인덱스]
- len == 끝인덱스 - 시작인덱스
- cap == 최대인덱스 - 시작인덱스
```go
	slice13 := []int{1, 2, 3, 4, 5}
	slice14 := slice13[1:3]
	slice15 := slice13[1:3:4]
	// slice13:  [1 2 3 4 5] 5 5
	// slice14:  [2 3] 2 4
	// slice15:  [2 3] 2 3
	slice16 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice17 := slice16[1:3]
	slice18 := slice16[1:3:4]
	// slice13:  [1 2 3 4 5 6 7 8] 8 8
	// slice14:  [2 3] 2 7 (최대인덱스: 8 - 시작인덱스: 1)
	// slice15:  [2 3] 2 3 (최대인덱스: 4 - 시작인덱스: 1)
```
- 끝에서 한개전 (끝 인덱스를 정확하게 적어야함:음수X)
```go
	slice19 := []int{1, 2, 3, 4, 5}
	slice20 := slice19[2:len(slice19)-1]
	// slice19:  [1 2 3 4 5] 5 5
	// slice20:  [3 4] 2 3 
```
- 파이썬이랑 고랭의 슬라이스는 다르다.
```python
// 파이썬 슬라이스는 배열 복사
a = [1, 2, 3]
b = a[0:2]
b[0] = 100
print(a) // [1, 2, 3]
print(b) // [100, 2]
// 끝에서 한개 전까지 (음수 지원)
c = a[1:-1] 
// 파이썬에선 배열 복사할때 사용 (고랭은 기존 배열)
slices = slices[:]
```

## append()
`변수 := append([]int{}, 복사할배열...)` == `변수 := append([]int{}, 복사할배열[0], 복사할배열[1], 복사할배열[...])`

```
	slice := []int{1, 2, 3, 4, 5}
	slice3 := append([]int{}, slice...)
	fmt.Println(slice3)
	// [1 2 3 4 5]

	slice4 := append([]int{}, slice[0], slice[1], slice[2], slice[3], slice[4])
	fmt.Println(slice4)
	// [1 2 3 4 5]
```

### append () 동작원리 
- 기존 슬라이스가 바뀔수도 있고 아닐수 있음
- 복사 -> 추가 -> 새로운 struct
- 빈공간 충분 -> 빈공간 요소 추가
- 빈공간 충분 X -> 새로운 배열 할당 -> 복사 -> 요소 추가
- 남은 빈공간 = cap - len
- cap - len >= 요소갯수 (true: 추가 / false: 복사)

# slice 정렬
int -> Ints()
float64 -> Float64s()
string -> strings()