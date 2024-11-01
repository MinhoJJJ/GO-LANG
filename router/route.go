package router

import (
    "AI/handlers"
    "net/http"
)

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
