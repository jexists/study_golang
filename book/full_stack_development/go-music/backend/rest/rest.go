package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	/*
		// Gin 엔진
		r := gin.Default()
	*/

	// 핸들러 생서
	h, _ := NewHandler()

	return RunAPIWithHandler(address, h)
	/*
		// r.GET("/relativepath/to/url", func(c *gin.Context) {})

		r.GET("/products", h.GetProducts)
		r.GET("/promos", h.GetPromos)

		// 사용자 로그인
		// r.POST("/users/signin", h.SignIn)
		// r.POST("/users", h.AddUser)
		// r.POST("/user/:id/signout", h.SignOut)
		// r.POST("/user/:id/orders", h.GetOrders)
		// r.POST("/users/charge", h.Charge)

		// 그룹 라우팅
		// - url 일부를 공유하는 http 라우팅은 같은 코드 블록으로 묶음 가능
		usersGroup := r.Group("/user")
		{
			usersGroup.POST("/signin", h.SignIn)
			usersGroup.POST("/charge", h.Charge)
			usersGroup.POST("", h.AddUser)
		}
		userGroup := r.Group("/user")
		{
			userGroup.POST("/:id/signout", h.SignOut)
			userGroup.POST("/:id/orders", h.GetOrders)
		}

		// 서버 시작
		return r.Run(address)
	*/
}

// 핸들러를 인자로 전달받는 함수
func RunAPIWithHandler(address string, h HandlerInterface) error {
	// Gin 엔진
	r := gin.Default()

	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	usersGroup := r.Group("/user")
	{
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("", h.AddUser)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.POST("/:id/orders", h.GetOrders)
	}
	return r.Run(address)
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("*********")
		c.Next()
		fmt.Println("*********")
	}
}

/*
// 미들웨어 추가 방식 1 : 기본 미들웨어 유지 & 새로운 미들웨어 추가
func RunAPIWithHandler(address string, h HandlerInterface) error {
	// gi 기본 엔진
	r := gin.Default()
	r.Use(MyCustomLogger())
	// 나머지 코드
	return r.Run(address)
}
// 미들웨어 추가 방식 2 : 기본 미들웨어 사용X & 새로운 미들웨어 추가
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.New()
	r.Use(MyCustomLogger())
	// 나머지 코드
	return r.Run(address)
}
// 2개 이상의 미들웨어 필요
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.New()
	r.Use(MyCustomLogger1(), MyCustomLogger2(), MyCustomLogger3())
	// 나머지 코드
	return r.Run(address)
}
*/

// 미들웨어
// http요청을 처리하는 핸들러 실행 전에 실행되는 코드

// Gin 웹서버 기본 미들웨어
// 1. logger 미들웨어: 어플리케이션 실행되는 동안 모든 활동을 로그로 남김
// 2. recovery 미들웨어: 어플리케이션에서 패닉이 발생하면 500번 http 에러 코드로 응답
// github.com/gin-gonic/contrib

// 커스텀 미들웨어
func MyCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청을 처리하기 전에 실행할 코드
		// 예제 변수 설정
		c.Set("v", "123")
		//  c.Get("v") // 호출하면 변수값 확인 가능

		// 요청 처리 로직 실행
		c.Next()

		// 이 코드는 핸들러 실행이 끝나면 실행된다.

		// 응답 코드 확인
		status := c.Writer.Status()
		// status 활용하는 코드 추가
		fmt.Println(status)
	}
}
