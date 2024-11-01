package main

import (
	"AI/config"
	"AI/handlers" // 모듈화
	"log"
	"net/http"
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

	// 정적 파일 제공을 위한 핸들러 설정
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", config.AddMimeTypeHandler(http.StripPrefix("/static/", fileServer)))

	// http 함수를 이용한 URL 경로에 대한 핸들러 함수를 등록
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/login.do", handlers.LoginHandler)

	// 서버 설정 및 시작
	serverConfig := config.NewServerConfig()
	log.Fatal(config.StartServer(serverConfig))

}
