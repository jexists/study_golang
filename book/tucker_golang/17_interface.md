# 인터페이스
- 구현을 포함하지 않는 메서드 집합
- 구체화된 타입이 아닌 인터페이스만 가지고 메서드 호출
- 상호작용면 (interface)
- 메서드 구현을 포함한 구체화된 객체가 아닌 추상화된 객체로 상호작용할 수 있음
- 인터페이스 변수 선언 가능 (변수의 값으로 사용가능)
- 타입
- 추상화(abstraction): 내부 동작(구현)을 감춰서 서비스 제공자와 사용자에게 자유를 주는 방식
- 의존성 Down, 결합도 Up (decoupling)
- 추상계층(Abstract Layer)

## 인터페이스 선언
- type + 인터페이스명 + inerface키워드 + 중괄호 { 메서드 집합 }
- 메서드명 필수
- 매개변수와 반환이 다르더라도 같은 이름 메서드 불가
- 메서드 구현을 포함X (signiture: 타입, 이름, 출력)

```
type InerfaceName interface {
  Method()
  MethodSet()
  // Method(int) string // ERROR (중복 메서드)
  // _(x int) // ERROR (메서드명 필수)
}
```