## switch 문
- 비교값과 같은 case에 해당하는 문장 실행
- 여러 조건 검사할때 유용
- 변수값에 다라 다른 명령 실행할때 유용

```go
	switch 비교값 { // 검사하는 값
		case 값1: // 비교값이 값1이 같을 때 실행
		문장
		case 값2: // 비교값이 값2이 같을 때 실행
		문장
		default: // 만족하는 case가 없을 때 실행 (생략 가능)
		문장
	}
```

### 하나의 case 한번에 여러값 비교
```go
  var day string
	switch day {
    case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("평일")
	case "saturday", "sunday":
		fmt.Println("주말")
	}
```

### 조건문 비교
- if문처럼 true가 되는 조건문 검사
- default 값 true 사용
- `switch true {...}` == `switch {...}`
```go
	switch { // 비교값 생략하면 비교값은 true 
		case 조건문1: // 조건문1이 true이면 실행
		문장
		case 조건문2: // 조건문2이 true이면 실행
		문장
		default: // 모든 case 조건문이 false이면 실행 (생략 가능)
		문장
	}
```
```go
	temp := 18
	switch true {
	case temp < 10, temp < 30:
		fmt.Println("추운날씨")
	case temp > 15, temp < 30:
		fmt.Println("약간 추위")
	case temp >= 15, temp < 25:
		// 실행안함 (두번째 조건에서 실행후 종료)
		fmt.Println("좋은 날씨")
	default:
		fmt.Println("따뜻")
	}
```

## 초기문
```go
	switch 초기문; 비교값 { // 초기문 먼저 실행되고 비교값 case들과 비교
    case 값1: // 비교값이 값1이 같을 때 실행
		문장
		case 값2: // 비교값이 값2이 같을 때 실행
		문장
		default: // 만족하는 case가 없을 때 실행 (생략 가능)
		문장
	}
```

## break 
- break 사용 안해도 case하나 실행 후 자동으로 switch문 종료
- break 옵션 (go언에서는 break 상관없이 switch문 종료)
- 다른언어에서는 switch문의 각 case 종료 시에 break문 사용해야 다음 case 실행 안함

## fallthrough
- case실행 후 다음 case까지 실행하고 싶을때 사용