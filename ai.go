package main

import (
	"AI/config"
	"AI/router"
	"log"

	"github.com/gin-gonic/gin"
)

// :=는 Go 언어에서 변수를 선언하고 동시에 값을 할당하는 단축 문법입니다.
func main() {

	// Gin 엔진 초기화
	r := gin.Default()

	// MIME 타입 초기화 및 템플릿 로드
	if err := config.InitMimeTypes(r); err != nil {
		log.Fatal("Failed to initialize MIME types and templates:", err)
	}

	// DB 초기화 및 커넥션
	var err error              // 초기화 과정에서 발생할 수 있는 오류 저장할 변수
	db, err := config.InitDB() // config 패키지에 정의된 InitDB 함수를 호출하여 데이터베이스 연결을 초기화합니다.
	if err != nil {            // 초기화 과정에서 오류가 발생했는지 확인합니다. err이 nil이 아닌 경우, 즉 오류가 발생한 경우에 실행됩니다.
		log.Fatal("Failed to initialize database:", err) // log.Fatal 함수는 오류 메시지를 출력하고, 프로그램을 즉시 종료합니다.
	}
	defer db.Close() // defer 키워드는 해당 함수를 종료할 때 db.Close() 메서드를 호출하여 데이터베이스 연결을 종료하도록 예약합니다.

	// 미들웨어 설정
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// DB를 gin.Context에 저장하기 위한 미들웨어
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 라우트 초기화 - router.InitRoutes 함수 직접 호출
	router.InitRoutes(r)

	// 서버 설정 및 시작
	r.Run(":8080")

	//serverConfig := config.NewServerConfig()
	//log.Fatal(config.StartServer(serverConfig))

}
