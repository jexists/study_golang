
### 이항연산자

→ A+B

### 단항연산자

→ -A

---

### 산술연산자

→ + - x /

### 비트 연산자

→ & | ^

→ binary_op  = "||" | "&&" |rel_op |add_op |mul_op 

### 논리연산자

→ < > == ≠ && || !

→ rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .

### 그외

add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .
unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

```go
package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b) //a&b = 0
	fmt.Printf("a|b = %v\n", a|b) //a|b = 6
	fmt.Println("result", a^b)    //result 6
}
```

```go
package main

import (
	"fmt"
)

func main() {
	a := 21
	c := a % 10                              //1
	a = a / 10                               //2
	d := a % 10                              //2.1 -> 2
	fmt.Printf("첫번째 수 %v, 두번째 수 %v\n", c, d) //첫번째 수 1, 두번째 수 2
}
```

[The Go Programming Language Specification](https://golang.org/ref/spec#Operators)