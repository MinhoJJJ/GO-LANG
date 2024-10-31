package main

import (
	"AI/config"
	"AI/handlers" // 모듈화
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

// 데이터베이스 연결을 위한 전역 변수
var db *sql.DB

// :=는 Go 언어에서 변수를 선언하고 동시에 값을 할당하는 단축 문법입니다.
func main() {

	// 데이터베이스 초기화
	var err error

	db, err = config.InitDB()

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// MIME 타입 등록
	// MIME 타입은 파일의 "설명서" 같은 것입니다
	// 우편물에 "취급주의" 스티커를 붙이는 것처럼, 파일에 "이건 CSS파일이에요"라는 표시를 해주는 것입니다
	// 이 표시가 없으면 브라우저는 보안상의 이유로 파일을 원하는 방식으로 처리하지 않습니다
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")

	// 정적 파일 제공을 위한 핸들러 설정
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", addMimeTypeHandler(fileServer)))

	// http 함수를 이용한 URL 경로에 대한 핸들러 함수를 등록합니다.
	// 이렇게 등록된 핸들러들은 DefaultServeMux에서 관리됨
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	// HandleFunc 등록 부분을 다음과 같이 수정:
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	})
	http.HandleFunc("/login.do", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	})

	// 포트 설정
	port := os.Getenv("PORT") // 포트에 대한 환경변수 값을 불러옴

	// PORT 환경 변수가 설정되어 있지 않은 경우(즉, 빈 문자열인 경우)
	if port == "" {
		port = "8080" // port에 8080 선언
	}

	// 로그설정
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	// 서버 실행(에러가 발생했을 때 에러 메시지를 출력하고 프로그램을 종료)
	// http.ListenAndServe(주소, nil) HTTP 서버를 시작하는 함수
	// nil은 기본 라우터(DefaultServeMux)를 사용하겠다는 의미입니다.
	// fmt.Sprintf(":%s", port)를 사용하는 주된 이유는 유연성(flexibility)과 동적인 포트 할당 때문
	// localhost를 지정하는 것은 ListenAndServe의 첫 번째 매개변수에서 합니다.
	// localhost로만 제한하고 싶다면 다음과 같이 수정하면 됩니다:
	// log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil)) // 동적으로 포트 문자열 생성
}

// MIME 타입을 설정하는 미들웨어
// MIME (Multipurpose Internet Mail Extensions) 타입은 웹에서 파일의 형식을 식별하기 위한 표준
// 브라우저는 서버로부터 받은 파일이 어떤 종류인지 알아야 합니다
func addMimeTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 파일 확장자 확인
		ext := filepath.Ext(r.URL.Path)

		// 확장자에 따른 Content-Type 설정
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		}

		next.ServeHTTP(w, r)
	})
}

// w http.ResponseWriter: HTTP 응답을 클라이언트에게 보내기 위한 인터페이스
// r *http.Request: 클라이언트로부터 받은 HTTP 요청 정보를 담고 있는 포인터
func indexHandler(w http.ResponseWriter, r *http.Request) {

	// 요청된 URL의 경로를 가져옵니다
	if r.URL.Path != "/" {
		// 404 Not Found 에러를 반환
		http.NotFound(w, r)
		return
	}

	//작동 순서:
	//fmt.Fprint()로 응답을 보냅니다
	//이 작업의 결과로 두 가지 값이 반환됩니다

	//첫 번째 값(작성된 바이트 수)은 _로 무시
	//두 번째 값은 err에 저장 (에러가 없으면 nil, 있으면 에러 정보)

	//err != nil 체크:
	//err이 nil이면 = 정상 작동
	//err이 nil이 아니면 = 에러 발생 → 500 에러 반환

	// err은 작성 작업 중 발생할 수 있는 에러를 나타냅니다
	// 이 nil은 "에러가 없음"을 의미합니다
	// if err != nil은 "에러가 발생했는지" 확인하는 구문입니다
	_, err := fmt.Fprint(w, "Hellodfgdfgdfg, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About Page</h1><p>This is a simple Go-based HTTP server.</p>")
}

func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
