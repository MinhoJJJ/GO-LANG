package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 서버 실행(에러가 발생했을 때 에러 메시지를 출력하고 프로그램을 종료)
// http.ListenAndServe(주소, nil) HTTP 서버를 시작하는 함수
// nil은 기본 라우터(DefaultServeMux)를 사용하겠다는 의미입니다.
// fmt.Sprintf(":%s", port)를 사용하는 주된 이유는 유연성(flexibility)과 동적인 포트 할당 때문
// localhost를 지정하는 것은 ListenAndServe의 첫 번째 매개변수에서 합니다.
// localhost로만 제한하고 싶다면 다음과 같이 수정하면 됩니다:
// log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil))

// PORT 환경 변수가 설정되어 있지 않은 경우(즉, 빈 문자열인 경우)

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Host         string
}

// NewServerConfig 함수 이름을 대문자로 시작
func NewServerConfig() *ServerConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &ServerConfig{
		Port:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Host:         "",
	}
}

// StartServer 함수 이름을 대문자로 시작
func StartServer(cfg *ServerConfig) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Printf("Server starting on %s:%s", cfg.Host, cfg.Port)
	log.Printf("Open http://localhost:%s in the browser", cfg.Port)

	return server.ListenAndServe()
}

//config\server.go:1:1: expected 'package', found 'EOF'
