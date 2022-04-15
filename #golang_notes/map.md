# Maps

→ Map / Dictionary / hash table

→ key - value 형태로 저장 (예: 전화번호부)

→ 키를 값에 매핑

→ zero vallue: nil 

→ nil 맵: 키x, 키 추가 불가

→ make 함수: 주어진 타입 초기화, 사용준비된 맵 반환

→ 맵 리터럴: 구조체 리터럴과 같지만, 키가 필요

→ 최상위 타입이 타입 이름일 경우, 리터럴의 요소에서 생략 가능

### 요소를 추가하거나 업데이트

```
m[key] = elem
```

### 요소 검색

```
elem = m[key]
```

### 요소 제거

```
delete(m, key)
```

### 두 개의 값을 할당하여 키가 존재하는지 테스트

```
elem, ok = m[key]
```

→ key가 m 안에 있다면 ok가 true / 아니라면 false

→ key가 맵 안에 없다면, elem은 map의 요소 타입의 zero value

### 아직 선언되지 않았다면, 정의식을 사용가능

```
elem, ok := m[key]
```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) //The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) //The value: 48

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) //The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) //The value: 0 Present? false
}
```