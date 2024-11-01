package handlers

import (
	"html/template"
	"net/http"
)

// HTML 템플릿 파일을 읽고 렌더링하여 HTTP 응답에 씁니다.
// w는 Go의 HTTP 서버에서 응답을 작성하는 데 사용되는 http.ResponseWriter 인스턴스입니다.

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + filename) // 지정된 파일 경로에서 템플릿 파일을 불러와서 담습니다.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 실행중 오류가 나면 HTTP 500 서버내부오류 응답을 반환
		return
	}

	err = tmpl.Execute(w, data) // 담겨진 템플릿을 w에 렌더링하고 data를 템플릿에 삽입해 클라이언트에 전송합니다.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 실행중 오류가 나면 HTTP 500 서버내부오류 응답을 반환
	}
}
