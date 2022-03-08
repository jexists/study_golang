## 내장타입 (built-in type)

### 1. 정수: integer

### 2. 실수: float

### 3. 문자열: string

### 4. 불리언: boolean

### 5. 복소수: complex number

## 리터럴

### 1. 정수 리터럴 (integer)

→ 숫자 (입반적으로 10진수)

→ ob (2진수) / 0o(8진수) / 0x(16진수)

→ 정수 리터럴 사이에 밑줄 넣는 것 허용

```go
//정수 리터럴 사이에 밑줄 넣는 것 허용 예시

package main

import "fmt"

func main() {
	var num = 1_234
	fmt.Println(num) // 1234
}
// 해당 밑줄은 숫자 값 자체에 영향X 
// → 숫자 맨앞이나 맨뒤에 넣을 수 없고 다른 숫자와 연결되게 사용X
// → 10진수 숫자를 천단위 나타낼때 추천 (예: 1_234)
// → 2진수,8진수,16진수를 1,2,4바이트 경계로 나누어 표시할때 추천
// 비추천) var num = 1_2_3_4
```

### 2. 부동소수점 리터럴 (floating point)

→ 소수부를 구분하는 소수점

→ 문자 e와 양수 혹은 음수로 지정된 지수

### 3. 룬 리터럴 (rune)

→ 문자 (작은 따옴표 묶어 사용)

→ 역슬래시 특수 룬 리터럴: 줄바꿈 (‘\n’), 탭 (’\t’), 작은따옴표 (’\’’), 큰따옴표 (’\”’), 역슬래시(’\\’)

### 4. 문자 리터럴

→ 문자 리터럴 표시하는 두가지 방법: 해석된 문자열 리터럴, 로우 문자열 리터럴 

→ 해석된 문자열 리터럴(interpreted string): 큰따옴표(”) 사용

→ 로우 문자열 리터럴(raw string): 역따옴표(`) 사용

```go
package main

import "fmt"

func main() {
	var interpretedString = "test \n \"hello\""
	var rawString = `test 
	"hello"`

	fmt.Println(interpretedString)
	fmt.Println(rawString)
}
```

## 기본 자료형

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 별칭

rune // int32의 별칭
     // 유니코드에서 code point를 의미합니다.

float32 float64

complex64 complex128
```