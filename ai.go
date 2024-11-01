package main

import (
	"AI/config"
	"AI/router"
	"log"
)

// :=는 Go 언어에서 변수를 선언하고 동시에 값을 할당하는 단축 문법입니다.
func main() {

	// MIME 타입 초기화
	config.InitMimeTypes()

	// DB 초기화 및 커넥션
	var err error
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// 라우트 초기화
	router.InitRoutes()

	// 서버 설정 및 시작
	serverConfig := config.NewServerConfig()
	log.Fatal(config.StartServer(serverConfig))

}
