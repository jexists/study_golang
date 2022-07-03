# 조건문

# If

→ `if 조건문bool값 {bool값이 true 실행 }`

→ `if 조건문bool값 {bool값이 true 실행 } else {bool값이 false 실행}`

→ `if 조건문bool값 { 첫번째 조건이 true일때 실행 } else if 조건문 { 첫번째 값이 false이지만 두번째 값이 true일때 실행} else { 모든 조건이 false 일때 실행 }` 

→ {}필수

```go
package main

import (
	"fmt"
)

func main() {
	var ok bool = true
	if ok {
		fmt.Println("hello")
	}
}
```

→ if문 변수 선언과 동시 사용가능 (다른곳에서는 선언한 변수 사용X)

```go
package main

import (
	"fmt"
)

func main() {
	if str := "hello world"; str == "HelloWorld" {
		fmt.Println(str)
	} else {
		fmt.Println("else", str)
	}
}
```

→ else 생략가능
```go
package main

import (
	"fmt"
)
func canIDrink(age int) bool {
	if age <18 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canIDrink(16))
}
```

→ if문 안에서 변수 선언 가능
```go
package main

import (
	"fmt"
)
func canIDrink(age int) bool {
	if koreanAge := age + 1; koreanAge <18 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canIDrink(16))
}
```

## Switch

→ 연속적인 if-else구문 사용하는 짧은 방안

→ 값이 조건문과 같은 첫번째 case 실행 (성공한 케이스가 있을 경우 종료)

→ break 자동 제공