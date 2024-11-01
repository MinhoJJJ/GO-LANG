package router

import (
	"AI/handlers"
	"net/http"
)

// GO의 라우트 역할
//  http.HandleFunc 메서드를 이용해서 특정 경로에 어떤 함수가 실행될지 설정
//  EX) 22번줄의 경우 localhost:8080/login 요청이 들어왔을 경우 handlers에 있는 LoginHandler가 실행됨
//  자바 스프링의 경우 어노테이션이 이 역할을 합니다.

// http.FileServer(http.Dir("static"))
// http.FileServer는 특정 디렉토리에 있는 파일을 제공하는 핸들러를 생성합니다.

// http.Handle는 특정 경로(/static/)로 요청이 들어올 때 어떤 핸들러를 사용할지 지정합니다.
// http.StripPrefix("/static/", fileServer)는 요청 경로에서 /static/을 제거하고 나머지 경로를 fileServer에 전달합니다.
// 예를 들어, 클라이언트가 /static/js/app.js에 접근하면, 실제 파일 서버는 static/js/app.js 경로를 찾습니다.

func InitRoutes() {
	// 정적 파일 라우트
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	initAuthRoutes()

}

func initAuthRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/login.do", handlers.LoginHandler)

}
