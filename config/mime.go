package config

import (
	"mime"
	"net/http"
	"path/filepath"
)

// MIME 타입을 설정하는 미들웨어
// MIME (Multipurpose Internet Mail Extensions) 타입은 웹에서 파일의 형식을 식별하기 위한 표준
// 브라우저는 서버로부터 받은 파일이 어떤 종류인지 알아야 합니다

// MIME 타입 등록
// MIME 타입은 파일의 "설명서" 같은 것입니다
// 우편물에 "취급주의" 스티커를 붙이는 것처럼, 파일에 "이건 CSS파일이에요"라는 표시를 해주는 것입니다
// 이 표시가 없으면 브라우저는 보안상의 이유로 파일을 원하는 방식으로 처리하지 않습니다

func InitMimeTypes() {
	// 기본 MIME 타입 등록
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".json", "application/json")
	// ... 필요한 MIME 타입 추가
}

func AddMimeTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path) //경로의 마지막 요소에서 확장자를 반환합니다.

		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".html":
			w.Header().Set("Content-Type", "text/html")
		case ".json":
			w.Header().Set("Content-Type", "application/json")
			// ... 필요한 MIME 타입 추가
		}

		next.ServeHTTP(w, r) // HTTP 요청을 다음 핸들러로 전달하는 역할을 합니다. 현재 핸들러에서 요청을 처리한 후, 다음 핸들러로 넘겨주는 기능을 제공합니다.
	})
}
